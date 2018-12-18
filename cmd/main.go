package main

import (
	"fmt"
	// bc "bchain-qlearning/internal/blockchain"
	. "bchain-qlearning/internal/learning"
	// ws "bchain-qlearning/internal/websocket"
)

func main() {
	// ws.WsInit()
	su := SimulatedUser{}
	su.InitLearner()
	// bc.InsertNRetrieve()

	fmt.Println("Main")
}
