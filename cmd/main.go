package main

import (
	"fmt"
	"sync"

	bc "github.com/pierback/bchain-qlearning/internal/blockchain"
	um "github.com/pierback/bchain-qlearning/internal/usermanagement"

	db "github.com/pierback/bchain-qlearning/pkg/database"

	// ws "github.com/pierback/bchain-qlearning/internal/websocket"
	dp "github.com/pierback/bchain-qlearning/cmd/deployment"
	en "github.com/pierback/bchain-qlearning/cmd/environment"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	en.SetEnvVars()

	db.StartDB()
	// db.StartJsonDB()

	// dp.DeploySC()

	//defer to close when you're done with it, not because you think it's idiomatic!

	if *en.DplFlag == "cffcn" || *en.DplFlag == "bvrglst" {
		dp.DeploySC()
	} else {

		// ws.WsInit()

		// bc.InsertNRetrieve()
		// bc.ReadWrite()

		/* su := l.SimulatedUser{}
		su.InitLearner() */

		go um.StartWorker()
		if *en.BcFlag == "watch" {
			go bc.Watch()
		} else if *en.BcFlag == "rw" {
			bc.TestBl()
		}

	}
	wg.Wait()
	fmt.Println("Main")
}
