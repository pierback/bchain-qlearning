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

	if *en.DplFlag == "cffcn" || *en.DplFlag == "bvrglst" {
		dp.DeploySC()
	} else if *en.SimFlag != 0 {
		um.StartUserSimulation(*en.SimFlag)
	} else {
		// bc.InsertNRetrieve()
		// bc.ReadWrite()

		if *en.BcFlag == "watch" {
			go um.StartWorker()
			bc.Watch()
		} else if *en.BcFlag == "rw" {
			bc.TestBl()
		}
	}
}
