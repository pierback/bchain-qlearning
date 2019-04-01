package main

import (
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
	// db.StartJsonDB()

	// dp.DeploySC()

	//defer to close when you're done with it, not because you think it's idiomatic!
	um.StartUserSimulation(8)

	if *en.DplFlag == "cffcn" || *en.DplFlag == "bvrglst" {
		dp.DeploySC()
	} else {

		// ws.WsInit()

		// bc.InsertNRetrieve()
		// bc.ReadWrite()

		// go um.StartWorker()
		if *en.BcFlag == "watch" {
			// go bc.Watch()
		} else if *en.BcFlag == "rw" {
			bc.TestBl()
		}

	}
}
