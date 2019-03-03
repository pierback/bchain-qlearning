package blockchain

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	bl "github.com/pierback/bchain-qlearning/internal/contracts/BeverageList"
)

const key = `{"address":"02e9f84165314bb8c255d8d3303b563b7375eb61","crypto":{"cipher":"aes-128-ctr","ciphertext":"f4952ba9725d5ae83f6e2c47714e4a1ed533a60a4ea97d635fd674e84b419f8a","cipherparams":{"iv":"98de0f7f0469a83ee3f543b232a9ec67"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"7629ecaf07277427a80e8e3138d9fc4058ca63b004f7ba3a2b253f22db55bced"},"mac":"b10cc5c56aa64f2246a5786ea3247ab6a133961b3b62985bc999dcc7cb51ade7"},"id":"3dfb452c-af90-41aa-aa8f-6da654b818e9","version":3}`
const Key = key

func Watch() {
	client, err := ethclient.Dial("ws://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := getLatestContractAddress()

	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	var ch = make(chan types.Log)
	ctx := context.Background()

	sub, err := client.SubscribeFilterLogs(ctx, query, ch)

	if err != nil {
		log.Fatal(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(bl.BeveragelistABI)))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Watch")

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case eventLog := <-ch:
			fmt.Println("eventlog")
			event := struct {
				Id      common.Address
				Time    [32]byte
				Drink   [32]byte
				Weekday [32]byte
			}{}

			err1 := contractAbi.Unpack(&event, "NewDrink", eventLog.Data)
			if err1 != nil {
				log.Fatal(err1)
			}

			fmt.Println(" ")
			fmt.Println("Id", event.Id.Hex())
			fmt.Println("time", string(event.Time[:]))       // foo
			fmt.Println("drink", string(event.Drink[:]))     // bar
			fmt.Println("weekday", string(event.Weekday[:])) // bar
			fmt.Println(" ")

			if err != nil {
				log.Println("Failed to unpack")
				continue
			}
		}
	}
}

func TestBl() {
	client, err := ethclient.Dial("ws://127.0.0.1:8545")
	if err != nil {
		fmt.Println("Unable to connect to network:", err)
	}

	contractAddress := getLatestContractAddress()

	fmt.Printf("wait\n\n")

	instance, err1 := bl.NewBeveragelist(contractAddress, client)
	if err1 != nil {
		fmt.Println("err1: ", err1)
		log.Fatal(err)
	}

	setDrinkData(instance)
}

func setDrinkData(instance *bl.Beveragelist) {
	auth, _ := bind.NewTransactor(strings.NewReader(key), "0000")
	drink := [32]byte{}
	weekday := [32]byte{}
	ti := [32]byte{}
	copy(drink[:], []byte("mate"))
	copy(weekday[:], []byte(time.Now().Weekday().String()))
	copy(ti[:], []byte(time.Now().String()))

	_, err := instance.SetDrinkData(auth, ti, drink, weekday)

	if err != nil {
		fmt.Println("SetDrinkData err: ", err)
	}
}

func getFromTime(instance *bl.Beveragelist, t [32]byte) {
	drink, weekday, _ := instance.GetDrinkData(nil, t)
	fmt.Println("getFromTime drink", string(drink[:]))     // foo
	fmt.Println("getFromTime weekday", string(weekday[:])) // foo
}

func getLastDrink(instance *bl.Beveragelist) {
	addr, drink, weekday, _ := instance.LastDrink(nil)

	fmt.Println("getDrinkData id", addr.Hex())              // foo
	fmt.Println("getDrinkData drink", string(drink[:]))     // foo
	fmt.Println("getDrinkData weekday", string(weekday[:])) // foo
}

func getLatestContractAddress() common.Address {
	dat, err := ioutil.ReadFile("/tmp/dat1")
	fmt.Printf("Contract Address 0x%s\n", common.Bytes2Hex(dat))
	if err != nil {
		fmt.Println("Error ReadFile:", err)
	}
	return common.BytesToAddress(dat)
}

func ReadVlog(vLog types.Log, client *ethclient.Client) [32]byte {
	fmt.Println("ReadVlog: ")

	contractAbi, err := abi.JSON(strings.NewReader(string(bl.BeveragelistABI)))
	if err != nil {
		log.Fatal(err)
	}

	/* fmt.Println(vLog.BlockHash.Hex()) // 0x3404b8c050aa0aacd0223e91b5c32fee6400f357764771d0684fa7b3f448f1a8
	fmt.Println(vLog.BlockNumber)     // 2394201
	fmt.Println(vLog.TxHash.Hex())    // 0x280201eda63c9ff6f305fcee51d5eb86167fab40ca3108ec784e8652a0e2b1a6
	*/
	event := struct {
		Id      common.Address
		Time    [32]byte
		Drink   [32]byte
		Weekday [32]byte
	}{}

	err1 := contractAbi.Unpack(&event, "NewDrink", vLog.Data)
	if err1 != nil {
		log.Fatal(err1)
	}

	str := fmt.Sprintf("%s", (event.Id))

	fmt.Println(" ")
	fmt.Println(" s := string(byteArray[:n])", string(event.Id[:20]))
	fmt.Println(" fmt.Sprintf", str)
	fmt.Println("id", string(event.Id[:]))           // foo
	fmt.Println("time", string(event.Time[:]))       // foo
	fmt.Println("drink", string(event.Drink[:]))     // bar
	fmt.Println("weekday", string(event.Weekday[:])) // bar
	fmt.Println(" ")

	block, err := client.BlockByNumber(context.Background(), big.NewInt(int64(vLog.BlockNumber)))
	if err != nil {
		log.Fatal(err)
	}

	for _, tx := range block.Transactions() {
		/* fmt.Println(tx.Hash().Hex())            // 0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2
		fmt.Println(tx.Value().String())        // 10000000000000000
		fmt.Println(tx.Gas())                   // 105000
		fmt.Println(tx.GasPrice().Uint64())     // 102000000000
		fmt.Println(tx.Nonce())                 // 110644
		fmt.Println(tx.Data())                  // [] */
		fmt.Println("Recipient", tx.To().Hex()) // 0x55fE59D8Ad77035154dDd0AD0388D09Dd4047A8e

		/* receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(receipt.Status) // 1 */
	}

	return event.Time
}
