package usermanagement

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	l "github.com/pierback/bchain-qlearning/internal/learning"
	db "github.com/pierback/bchain-qlearning/pkg/database"
)

var dbFile *os.File

//StartWorker starts worker job
func StartWorker() {

	defer dbFile.Close()
	fmt.Printf("\nWorker started \n")
	initBm()

	run()

	<-nextTick()
	run()

	for range time.Tick(3 * time.Hour) {
		// for range time.Tick(30 * time.Second) {
		run()
		dbFile, err := os.OpenFile("logs", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer dbFile.Close()
		wrt := io.MultiWriter(os.Stdout, dbFile)
		log.SetOutput(wrt)
	}
}

func nextTick() <-chan time.Time {
	curH := time.Now().Hour()
	curts := l.GetCurrentTimeSlot(curH)
	_, lb := l.TsBoundaries(uint(curts))

	nextTick := time.Date(time.Now().Year(), time.Now().Month(),
		time.Now().Day(), int(lb+1), int(0), int(0), int(0), time.Local)
	diff := nextTick.Sub(time.Now())
	d := diff.Round(diff)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute

	diffStr := fmt.Sprintf("%02dh %02dmin", h, m)
	fmt.Printf("nextTick at: %02d:%02d:%02d in %s\n", nextTick.Hour(), nextTick.Minute(), nextTick.Second(), diffStr)

	fmt.Printf("time.Tick(diff) %s\n", diff)

	return time.Tick(diff)
}

func run() {
	log.Printf("Timeslot ended worker runs: \n\n")
	bcm.mu.RLock()
	defer bcm.mu.RUnlock()

	//last action of user in current timeslot
	actn := l.Nothing
	var qvalsN, qvalsC float64

	for usr, ql := range bcm.users {
		log.Println("next urs", usr)
		ql.Learn(actn)
		qvalsN += ql.GetQ(l.Nothing, ql.GetState())
		qvalsC += ql.GetQ(l.Coffee, ql.GetState())

		go saveToDb(usr.ethaddress, ql)

		log.Printf("User %s Steps: %d/ Neg: %d \n", usr, ql.Sr.Steps, ql.Sr.Neg)

		if time.Now().Weekday() == time.Friday && time.Now().Hour() == 20 {
			ql.Sr.Wa = append(ql.Sr.Wa, ql.Sr.Neg)
			ql.Sr.Neg = 0
		}
	}
	bcm.SetGenQvals(qvalsN, qvalsC)
}

func saveToDb(usr string, ql l.QLearning) {
	qt := l.MapToString(ql.Qt)
	ep := fmt.Sprintf("%f", ql.Epsilon)
	db.SaveQl(usr, qt, ep, ql.Sr.Neg, ql.Sr.Wa)
}

//Learn triggers learning func of qlearning
func Learn(ethAdrs string, at string) {
	log.Printf("Learn ethAdrs: %s \n", ethAdrs)
	actn := l.GetAction(at)
	if actn == l.Coffee || actn == l.Nothing {
		bcm.mu.RLock()
		q := bcm.getQlearning(ethAdrs)
		q.Learn(actn)
		bcm.set(User{ethAdrs}, *q)
	}
}
