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
	var args string
	flag.StringVar(&args, "dp", "bar", "a string var")
	flag.Parse()

	if args == "cffcn" || args == "bvrglst" {
		dp.DeploySC(args)
	} else {

		// ws.WsInit()

		// bc.InsertNRetrieve()
		// bc.ReadWrite()
		// l.StartWorker()

		/* su := l.SimulatedUser{}
		su.InitLearner() */

		// bc.Watch()
		// bc.TestBl()

		_ = l.SimulatedUser{}

		_ = bc.Key
		fmt.Println("Main")
	}
}
