package blockchain

import (
	"log"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	bl "github.com/pierback/bchain-qlearning/internal/contracts/BeverageList"
	ut "github.com/pierback/bchain-qlearning/pkg/utils"
)

func TestBl() {
	client := ut.GetClientConnection()
	contractAddress := ut.GetLatestContractAddress()
	instance, err1 := bl.NewBeveragelist(contractAddress, client)
	ut.PrintError(err1)

	setDrinkData(instance)
}

func setDrinkData(instance *bl.Beveragelist) {
	auth, _ := bind.NewTransactor(strings.NewReader(key), "password")
	drink := [32]byte{}
	weekday := [32]byte{}
	ti := [32]byte{}
	copy(drink[:], []byte("mate"))
	copy(weekday[:], []byte(time.Now().Weekday().String()))
	copy(ti[:], []byte(time.Now().String()))

	_, err := instance.SetDrinkData(auth, ti, drink, weekday)

	ut.PrintError(err)
}

func getFromTime(instance *bl.Beveragelist, t [32]byte) {
	drink, weekday, _ := instance.GetDrinkData(nil, t)
	log.Println("getFromTime drink", string(drink[:]))
	log.Println("getFromTime weekday", string(weekday[:]))
}

/* func getLastDrink(instance *bl.Beveragelist) {
	addr, drink, weekday, _ := instance.LastDrink(nil)

	fmt.Println("getDrinkData id", addr.Hex())              // foo
	fmt.Println("getDrinkData drink", string(drink[:]))     // foo
	fmt.Println("getDrinkData weekday", string(weekday[:])) // foo
} */
