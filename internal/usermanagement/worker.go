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
	fmt.Printf("\nWorker started \n\n")
	initBm()

	ticker := nextTick()
	go func() {
		for {
			select {
			case <-ticker.C:
				ticker.Stop()
				run()
				ticker = nextTick()
				dbFile, err := os.OpenFile("logs", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
				if err != nil {
					log.Fatal(err)
				}
				defer dbFile.Close()
				wrt := io.MultiWriter(os.Stdout, dbFile)
				log.SetOutput(wrt)
			}
		}
		log.Println("stopped")
	}()

}

func nextTick() time.Ticker {
	var lowerBoundary int
	curH := time.Now().Hour()
	curts := l.GetCurrentTimeSlot(curH)
	day := time.Now().Day()

	//get lower boundary of next relevant timeslot
	if curts+1 > 3 {
		_, lowerBoundary = l.TsBoundaries(4) //is always 6

		//next tick next monday at 7
		if time.Weekday(day) == time.Friday {
			day = time.Now().AddDate(0, 0, 3).Day()
		} else if time.Now().Hour() < 7 { //ticker started after 00:00
			day = time.Now().Day()
		} else { //ticker started after curts+1
			day = time.Now().AddDate(0, 0, 1).Day()
		}
	} else {
		_, lowerBoundary = l.TsBoundaries(uint(curts))
	}
	hour := lowerBoundary + 1

	nextTick := time.Date(time.Now().Year(), time.Now().Month(),
		day, hour, int(0), int(0), int(0), time.Local)

	diff := nextTick.Sub(time.Now())
	d := diff.Round(diff)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute

	diffStr := fmt.Sprintf("%02dh %02dmin", h, m)
	fmt.Printf("firstTick at: %02d:%02d:%02d in %s\n", nextTick.Hour(), nextTick.Minute(), nextTick.Second(), diffStr)

	return *time.NewTicker(diff)
	// return *time.NewTicker(10 * time.Second)
}

func run() {
	log.Printf("Timeslot ended worker runs: \n\n")
	bcm.mu.RLock()
	defer bcm.mu.RUnlock()

	//last action of user in current timeslot
	actn := l.Nothing
	var qvalsN, qvalsC float64

	for usr, ql := range bcm.users {
		ql.Learn(actn)
		qvalsN += ql.GetQ(l.Nothing, ql.GetState())
		qvalsC += ql.GetQ(l.Coffee, ql.GetState())

		go saveToDb(usr.ethaddress, ql)

		log.Printf("User %s Steps: %d/ Neg: %d \n\n", usr, ql.Sr.Steps, ql.Sr.Neg)

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
