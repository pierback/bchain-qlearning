package usermanagement

import (
	"fmt"
	"time"

	l "github.com/pierback/bchain-qlearning/internal/learning"
	db "github.com/pierback/bchain-qlearning/pkg/database"
)

//StartWorker starts worker job
func StartWorker() {
	// fmt.Println("worker started")
	initBm()

	for range time.Tick(4 * time.Minute) {
		run()
	}
}

func run() {
	bcm.mu.RLock()
	defer bcm.mu.RUnlock()

	actn := l.Nothing
	var qvalsN, qvalsC float64

	for usr, ql := range bcm.users {
		fmt.Println("next tick", usr)
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
