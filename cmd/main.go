package main

import (
	"fmt"

	bc "github.com/pierback/bchain-qlearning/internal/blockchain"
	// l "github.com/pierback/bchain-qlearning/internal/learning"
	// ws "github.com/pierback/bchain-qlearning/internal/websocket"
)

func main() {
	// ws.WsInit()

	/* su := l.SimulatedUser{}
	su.InitLearner() */

	// bc.InsertNRetrieve()
	bc.StartDeploy()
	// l.StartWorker()

	fmt.Println("Main")
}
