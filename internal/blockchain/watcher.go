package blockchain

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	bl "github.com/pierback/bchain-qlearning/internal/contracts/BeverageList"
	wk "github.com/pierback/bchain-qlearning/internal/usermanagement"
	ut "github.com/pierback/bchain-qlearning/pkg/utils"
)

// const key = `{"address":"02e9f84165314bb8c255d8d3303b563b7375eb61","crypto":{"cipher":"aes-128-ctr","ciphertext":"f4952ba9725d5ae83f6e2c47714e4a1ed533a60a4ea97d635fd674e84b419f8a","cipherparams":{"iv":"98de0f7f0469a83ee3f543b232a9ec67"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"7629ecaf07277427a80e8e3138d9fc4058ca63b004f7ba3a2b253f22db55bced"},"mac":"b10cc5c56aa64f2246a5786ea3247ab6a133961b3b62985bc999dcc7cb51ade7"},"id":"3dfb452c-af90-41aa-aa8f-6da654b818e9","version":3}`
const key = `{"address":"e8816898d851d5b61b7f950627d04d794c07ca37","crypto":{"cipher":"aes-128-ctr","ciphertext":"1ff4add6955cba7ddaf29f66d7d21c5e1d714ef6191fbc651ae60f2ea3c95e8f","cipherparams":{"iv":"3ff869fbdbe1a523cdb327780365976e"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"7372dbae5fb318f8684902e099c311d4188721d677974d729711762c7ef6030c"},"mac":"485fa5dc701067782baa1589716a53110c7f917eb259e35ebca7265bbb7150b1"},"id":"89edb004-5b00-4607-a3af-a0d9ab9b1c34","version":3}`

func Watch() {
	client := ut.GetClientConnection()
	contractAddress := ut.GetLatestContractAddress()

	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	var ch = make(chan types.Log)
	ctx := context.Background()

	sub, err := client.SubscribeFilterLogs(ctx, query, ch)
	ut.PrintError(err)

	fmt.Println("\nStart watching for bchain events...")

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case eventLog := <-ch:
			readEvent(eventLog, client)

		}
	}
}

func readEvent(eventLog types.Log, client *ethclient.Client) {
	event := struct {
		Address common.Address
		Time    [32]byte
		Drink   [32]byte
		Weekday [32]byte
	}{}

	contractAbi, err := abi.JSON(strings.NewReader(string(bl.BeveragelistABI)))
	if err != nil {
		log.Println("Failed to unpack")
	}

	err1 := contractAbi.Unpack(&event, "NewDrink", eventLog.Data)
	ut.PrintError(err1)

	fmt.Println(" ")
	fmt.Println("New Event:")
	fmt.Println("   Address", event.Address.Hex())
	fmt.Println("   time", string(event.Time[:32]))
	fmt.Println("   drink", string(event.Drink[:32]))
	fmt.Println("   Weekday", string(event.Weekday[:32]))
	fmt.Println(" ")

	dr := string(bytes.Trim(event.Drink[:], "\x00"))
	wk.Learn(event.Address.Hex(), dr)
	// getBlockTransactions(eventLog, client)
}

func getBlockTransactions(eventLog types.Log, client *ethclient.Client) {
	block, err := client.BlockByNumber(context.Background(), big.NewInt(int64(eventLog.BlockNumber)))
	ut.PrintError(err)

	for _, tx := range block.Transactions() {
		fmt.Println("Recipient", tx.To().Hex())
	}
}
