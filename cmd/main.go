package main

import (
	"fmt"

	bc "github.com/pierback/bchain-qlearning/internal/blockchain"
	l "github.com/pierback/bchain-qlearning/internal/learning"

	// ws "github.com/pierback/bchain-qlearning/internal/websocket"
	dp "github.com/pierback/bchain-qlearning/cmd/deployment"
	en "github.com/pierback/bchain-qlearning/cmd/environment"
)

func main() {
	en.SetEnvVars()

	if *en.DplFlag == "cffcn" || *en.DplFlag == "bvrglst" {
		dp.DeploySC()
	} else {

		// ws.WsInit()

		// bc.InsertNRetrieve()
		// bc.ReadWrite()
		// l.StartWorker()

		/* su := l.SimulatedUser{}
		su.InitLearner() */

		if *en.BcFlag == "watch" {
			bc.Watch()
		} else if *en.BcFlag == "rw" {
			bc.TestBl()
		}
	}

	_ = l.SimulatedUser{}
	fmt.Println("Main")
}
