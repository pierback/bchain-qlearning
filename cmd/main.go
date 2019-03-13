package main

import (
	"fmt"

	bc "github.com/pierback/bchain-qlearning/internal/blockchain"
	um "github.com/pierback/bchain-qlearning/internal/usermanagement"

	db "github.com/pierback/bchain-qlearning/pkg/database"

	// ws "github.com/pierback/bchain-qlearning/internal/websocket"
	dp "github.com/pierback/bchain-qlearning/cmd/deployment"
	en "github.com/pierback/bchain-qlearning/cmd/environment"
)

func main() {
	en.SetEnvVars()

	db.StartDB()
	db.StartJsonDB()

	dp.DeploySC()

	if *en.DplFlag == "cffcn" || *en.DplFlag == "bvrglst" {
		dp.DeploySC()
	} else {

		// ws.WsInit()

		// bc.InsertNRetrieve()
		// bc.ReadWrite()

		go um.StartWorker()

		/* su := l.SimulatedUser{}
		su.InitLearner() */

		if *en.BcFlag == "watch" {
			bc.Watch()
		} else if *en.BcFlag == "rw" {
			bc.TestBl()
		}
	}
	fmt.Println("Main")
}
