package usermanagement

import (
	"fmt"
	"log"
	"os"
	"time"

	l "github.com/pierback/bchain-qlearning/internal/learning"
	db "github.com/pierback/bchain-qlearning/pkg/database"
)

var dbFile *os.File

//StartWorker starts worker job
func StartWorker() {
	dbFile, err := os.OpenFile("logs", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer dbFile.Close()
	fmt.Printf("\nWorker started \n")
	initBm()

	<-nextTick()
	run()

	for range time.Tick(3 * time.Hour) {
		run()
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
	bcm.mu.RLock()
	defer bcm.mu.RUnlock()

	//last action of user in current timeslot
	actn := l.Nothing
	var qvalsN, qvalsC float64

	log.Println("usrs ", len(bcm.users))
	for usr, ql := range bcm.users {
		log.Println("next urs", usr)
		ql.Learn(actn)
		qvalsN += ql.GetQ(l.Nothing, ql.GetState())
		qvalsC += ql.GetQ(l.Coffee, ql.GetState())
		// saveToDb(usr.ethaddress, ql)
		log.Printf("User %s Steps: %d/%d=%f \n", usr, ql.Sr.Neg, ql.Sr.Steps)

		if time.Now().Weekday() == time.Friday && time.Now().Hour() == 20 {
			ql.Sr.Wa = append(ql.Sr.Wa, ql.Sr.Neg)
			ql.Sr.Neg = 0
		}
	}
	bcm.SetGenQvals(qvalsN, qvalsC)
	log.SetOutput(dbFile)
}

func saveToDb(usr string, ql l.QLearning) {
	qt := l.MapToString(ql.Qt)
	fmt.Println("qt:", qt)
	ep := fmt.Sprintf("%f", ql.Epsilon)
	fmt.Println("ep:", ep)
	db.SaveQl(usr, qt, ep)
}

//Learn triggers learning func of qlearning
func Learn(ethAdrs string, at string) {
	log.Printf("Learn ethAdrs: %s \n", ethAdrs)
	actn := l.GetAction(at)
	if actn == l.Coffee || actn == l.Nothing {
		bcm.mu.RLock()
		q := bcm.getQlearning(ethAdrs)
		q.Learn(actn)
		// usr := User{ethAdrs}
		bcm.set(User{ethAdrs}, *q)
	}
}
