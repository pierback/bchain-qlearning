package learning

import (
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"
)

var (
	bcm  *bchainMock
	once sync.Once
)

//Worker for prediction evaluation
type Worker interface {
	run()
}

//BchainMock singleton constructor
func initBm() *bchainMock {
	once.Do(func() {
		bcm = &bchainMock{users: make(map[User]QLearning)}
		initUsers()
	})

	return bcm
}

func initUsers() {
	for index := 0; index < 300; index++ {
		ethaddrs := ethaddress("asdfasdf" + strconv.Itoa(index))
		fmt.Println("ethaddrs: ", ethaddrs)
		usr := User{ethaddrs}
		ql := usr.InitLearner()
		bcm.Set(usr, ql)
		tmpState := UserState{}
		ql.state = tmpState.New(drinkcount{}, -1, -1)
		fmt.Println("q.state: ", ql.state)
	}
	bcm.printUsers()
}

//StartWorker starts worker job
func StartWorker() {
	fmt.Println("worker started")
	initBm()

	for range time.Tick(10 * time.Second) {
		run()
	}
}

func run() {
	actn := Nothing
	for usr, ql := range bcm.users {
		fmt.Println("next tick", ql, usr)
		ql.learn(actn)
	}
}

func websocketInput(ethAdrs ethaddress, at Action) {
	// input is eth-address && action took
	// compare action took with usr.predicted action
	// get new state make prediction for next state/timeslot

	q, err := bcm.getQlearning(ethAdrs)
	if err != nil {
		log.Println("No user found with this address")
		return
	}
	q.learn(at)
}
