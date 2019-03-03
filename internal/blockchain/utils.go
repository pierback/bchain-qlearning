package blockchain

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	bl "github.com/pierback/bchain-qlearning/internal/contracts/BeverageList"
	//"github.com/ethereum/go-ethereum/common/hexutil"
	// for demo
	// for demo
	//"github.com/ethereum/go-ethereum/common/hexutil"
)

const Nothing = ""

func getLatestContractAddress() common.Address {
	dat, err := ioutil.ReadFile("/var/tmp/bvrglst")
	fmt.Printf("Contract Address 0x%s\n", common.Bytes2Hex(dat))
	if err != nil {
		fmt.Println("Error ReadFile:", err)
	}
	return common.BytesToAddress(dat)
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
