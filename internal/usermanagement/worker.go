package usermanagement

import (
	"fmt"
	"time"

	l "github.com/pierback/bchain-qlearning/internal/learning"
	db "github.com/pierback/bchain-qlearning/pkg/database"
)

//StartWorker starts worker job
func StartWorker() {
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

	fmt.Println("usrs ", len(bcm.users))
	for usr, ql := range bcm.users {
		fmt.Println("next urs", usr)
		ql.Learn(actn)
		qvalsN += ql.GetQ(l.Nothing, ql.GetState())
		qvalsC += ql.GetQ(l.Coffee, ql.GetState())
		saveToDb(usr.ethaddress, ql)
	}
	bcm.SetGenQvals(qvalsN, qvalsC)
}

func saveToDb(usr string, ql l.QLearning) {
	qt := l.MapToString(ql.Qt)
	ep := fmt.Sprintf("%f", ql.Epsilon)
	db.SaveQl(usr, qt, ep)
}

//Learn triggers learning func of qlearning
func Learn(ethAdrs string, at string) {
	actn := l.GetAction(at)
	if actn == l.Coffee || actn == l.Nothing {
		bcm.mu.RLock()
		q := bcm.getQlearning(ethAdrs)
		q.Learn(actn)
	}
}
