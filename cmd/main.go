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
	var dplFlag = flag.String("dp", "", "deploy sm")
	var bcFlag = flag.String("bc", "watch", "test bchain")

	flag.Parse()

	if *dplFlag == "cffcn" || *dplFlag == "bvrglst" {
		fmt.Println("*dplFlag: ", *dplFlag)
		dp.DeploySC(*dplFlag)
	} else {

		// ws.WsInit()

		// bc.InsertNRetrieve()
		// bc.ReadWrite()
		// l.StartWorker()

		/* su := l.SimulatedUser{}
		su.InitLearner() */

		if *bcFlag == "watch" {
			bc.Watch()
		} else if *bcFlag == "rw" {
			bc.TestBl()
		}

		_ = l.SimulatedUser{}
		_ = bc.Filename
		fmt.Println("Main")
	}
}
