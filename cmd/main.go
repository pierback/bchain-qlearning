package main

import (
	"fmt"

	bc "github.com/pierback/bchain-qlearning/internal/blockchain"
	l "github.com/pierback/bchain-qlearning/internal/learning"
	// ws "github.com/pierback/bchain-qlearning/internal/websocket"
)

func main() {
	// ws.WsInit()

	// bc.InsertNRetrieve()
	// bc.ReadWrite()
	// l.StartWorker()

	su := l.SimulatedUser{}
	su.InitLearner()

	_ = l.SimulatedUser{}

	_ = bc.Queue{}
	fmt.Println("Main")
}
