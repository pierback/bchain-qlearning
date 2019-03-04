package main

import (
	"flag"
	"fmt"

	bc "github.com/pierback/bchain-qlearning/internal/blockchain"
	l "github.com/pierback/bchain-qlearning/internal/learning"

	// ws "github.com/pierback/bchain-qlearning/internal/websocket"
	dp "github.com/pierback/bchain-qlearning/cmd/deployment"
)

func main() {
	var dplFlag string
	flag.StringVar(&dplFlag, "dp", "bar", "a string var")
	flag.Parse()

	var bcFlag string
	flag.StringVar(&bcFlag, "bc", "bar", "a string var")
	flag.Parse()

	if dplFlag == "cffcn" || dplFlag == "bvrglst" {
		dp.DeploySC(dplFlag)
	} else {

		// ws.WsInit()

		// bc.InsertNRetrieve()
		// bc.ReadWrite()
		// l.StartWorker()

		/* su := l.SimulatedUser{}
		su.InitLearner() */

		if bcFlag == "watch" {
			bc.Watch()
		} else if bcFlag == "rw" {
			bc.TestBl()
		}

		_ = l.SimulatedUser{}

		_ = bc.Nothing
		fmt.Println("Main")
	}
}
