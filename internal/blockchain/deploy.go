package blockchain

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	bl "github.com/pierback/bchain-qlearning/internal/contracts/BeverageList"
	cfd "github.com/pierback/bchain-qlearning/internal/contracts/Store/CoffeeDash"
)

var jobQueue Queue

type singleton struct {
	dpi cfd.Coffeedash
}

var instance *cfd.Coffeedash
var once sync.Once

func GetInstance() *cfd.Coffeedash {
	once.Do(func() {
		fmt.Println("GetInstance")
		blockchain, err := ethclient.Dial("http://localhost:8545")
		fmt.Println("blockchain: ", blockchain)

		if err != nil {
			log.Fatalf("Unable to connect to network:%v\n", err)
		}
		/* 		auth, err := bind.NewTransactor(strings.NewReader(key), "0000")

		   		if err != nil {
		   			log.Fatalf("Failed to create authorized transactor: %v", err)
		   		} */
		// projektmodul blockchain contracts addresses
		//cont := common.HexToAddress("0x76696ffcb3fad0fea8228e99766222339a4abd24")
		// cont := common.HexToAddress("0xfaf1858ec3b324718f9401f633ca9880f9e64e03")
		cont := common.HexToAddress("0x2cfef625b8421dde00474d6805ef4758363fbdf7")

		//docker bchain adresses
		// cont := common.HexToAddress("0xb7f2adbf4a6449af41f29a40b92b1b9334990e43")

		//
		dpi, _ := cfd.NewCoffeedash(cont, blockchain)
		instance = dpi
		jobQueue = Queue{&sync.Mutex{}, make([]bchainData, 0)}
		auth, err := bind.NewTransactor(strings.NewReader(key), "0000")
		jobQueue.AddToQueue(hello, auth, "df", "df")

		go jobQueue.Worker()
	})
	return instance
}

func hello(*bind.TransactOpts, string, string) (*types.Transaction, error) {
	fmt.Println("hello hello ")
	return nil, nil
}

func ReadWrite() {
	// blockchain, err := ethclient.Dial("http://127.0.0.1:8501")
	blockchain, err := ethclient.Dial("http://localhost:8545")

	if err != nil {
		log.Fatalf("Unable to connect to network:%v\n", err)
	}
	auth, err := bind.NewTransactor(strings.NewReader(key), "0000")

	// projektmodul blockchain contracts addresses
	//cont := common.HexToAddress("0x76696ffcb3fad0fea8228e99766222339a4abd24")
	// cont := common.HexToAddress("0xfaf1858ec3b324718f9401f633ca9880f9e64e03")
	cont := common.HexToAddress("0x2cfef625b8421dde00474d6805ef4758363fbdf7")

	instance, err1 := cfd.NewCoffeedash(cont, blockchain)
	if err1 != nil {
		fmt.Println("err1: ", err1)
		log.Fatal(err)
	}

	// _, err = instance.SetQValue(auth, "coffee", "-0.d, -d")
	newstate := "{Monday 3 {4 0 0}}"
	if output, _ := instance.StateExists(nil, newstate); !output {
		_, err = instance.AddState(auth, newstate, "-0.9999999956953279, -0.99999822853")
		if err != nil {
			fmt.Printf("set error")
			log.Fatal(err)
		}
	} else {
		fmt.Println("StateExists: ", output)
	}

	length, _ := instance.GetStateCnt(nil)
	fmt.Println("length: ", length)

	for index := 0; index < int(length); index++ {
		state, val, err1 := instance.GetQtableState(nil, uint8(index))
		if err1 != nil {
			fmt.Printf("set error")
			log.Fatal(err)
		}
		fmt.Println("state", state, "values", val)
	}
}

func TestBl() {
	client, err := ethclient.Dial("ws://127.0.0.1:8545")
	if err != nil {
		fmt.Println("Unable to connect to network:%v\n", err)
	}
	contractAddress := common.HexToAddress("0x763f3b1223e4fc2d13993875a00cbb91cfb17aa7")
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	fmt.Println("wait\n")

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	instance, err1 := bl.NewBeveragelist(contractAddress, client)
	if err1 != nil {
		fmt.Println("err1: ", err1)
		log.Fatal(err)
	}

	auth, err := bind.NewTransactor(strings.NewReader(key), "0000")
	drink := [32]byte{}
	weekday := [32]byte{}
	ti := [32]byte{}
	copy(drink[:], []byte("cola"))
	copy(weekday[:], []byte("mannfred"))
	// copy(weekday[:], []byte(time.Now().Weekday().String()))
	copy(ti[:], []byte(time.Now().String()))
	_, errr := instance.SetDrinkData(auth, ti, drink, weekday)

	if errr != nil {
		fmt.Println("err1: ", errr)
	}

	getDrinkData(instance)

	for {

		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			// fmt.Println("vLog: ", vLog)

			tim := ReadVlog(vLog, client)
			getFromTime(instance, tim)
			// pointer to event log
		}
	}
}

func getFromTime(instance *bl.Beveragelist, t [32]byte) {
	drink, weekday, _ := instance.GetDrinkData(nil, t)
	fmt.Println("getFromTime drink", string(drink[:]))     // foo
	fmt.Println("getFromTime weekday", string(weekday[:])) // foo
}

func getDrinkData(instance *bl.Beveragelist) {
	drink, weekday, _ := instance.LastDrink(nil)
	fmt.Println("getDrinkData drink", string(drink[:]))     // foo
	fmt.Println("getDrinkData weekday", string(weekday[:])) // foo
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
		Time    [32]byte
		Drink   [32]byte
		Weekday [32]byte
	}{}

	err1 := contractAbi.Unpack(&event, "NewDrink", vLog.Data)
	if err1 != nil {
		log.Fatal(err1)
	}

	fmt.Println("time", string(event.Time[:]))       // foo
	fmt.Println("drink", string(event.Drink[:]))     // bar
	fmt.Println("weekday", string(event.Weekday[:])) // bar

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

		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(receipt.Status) // 1
	}

	return event.Time
}

func DeployBeveragelist() {
	// connect to an ethereum node  hosted by infura
	// blockchain, err := ethclient.Dial("http://127.0.0.1:8501")
	client, err := ethclient.Dial("ws://127.0.0.1:8545")

	if err != nil {
		log.Fatalf("Unable to connect to network:%v\n", err)
	}

	// Get credentials for the account to charge for contract deployments
	auth, err := bind.NewTransactor(strings.NewReader(key), "0000")

	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	address, _, _, err1 := bl.DeployBeveragelist(auth, client)
	if err1 != nil {
		fmt.Println("err1: ", err1)
		log.Fatal(err)
	}
	fmt.Printf("Contract pending deploy: 0x%x\n", address)
	/* for index := 0; index < int(length); index++ {
		state, values, err1 := instance.GetQtableState(nil, uint8(index))
		if err1 != nil {
			fmt.Printf("set error")
			log.Fatal(err)
		}
		fmt.Println("state", state, values)
	} */
}

func Playwithit() {
	blockchain, _ := ethclient.Dial("http://localhost:8545")
	cont := common.HexToAddress("0xa1d50d1e2e1be9293a82cb597086a46ff57b7193")
	instance, _ := cfd.NewCoffeedash(cont, blockchain)

	// _, err = instance.SetQValue(auth, "coffee", "-0.d, -d")
	/* newstate := `adsf` //`"{\"weekday\":1,\"timeslot\":0,\"drinkcount\":{\"coffeeCount\":0,\"waterCount\":0,\"mateCount\":0}}"`
	output, err := instance.StateExists(nil, newstate)
	if err != nil {
		fmt.Println("StateExists error", err, output)
		// log.Fatal(err)
	}
	fmt.Println("StateExists: ", output)
	// newstate := "{Monday 7 {4 0 0}}"
	AddState(newstate, "0ddd, 0ddd")*/
	length, _ := instance.GetStateCnt(nil)
	fmt.Println("length: ", length)

	for index := 0; index < int(length); index++ {
		state, val, err1 := instance.GetQtableState(nil, uint8(index))
		if err1 != nil {
			fmt.Printf("set error")
			log.Fatal(err1)
		}
		fmt.Println("state", state, "values", val)
	}
}

func UpdateContract() {
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}

	account := common.HexToAddress("0x02e9f84165314bb8c255d8d3303b563b7375eb61")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(balance)

	privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	if err != nil {
		log.Fatal(err)
	}

	/*publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

		fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
		 	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
			if err != nil {
				log.Fatal(err)
			} */

	address := common.HexToAddress("0xc50079926ba5b401D52048Aaf3698705aC63f106")
	newinstance, err := cfd.NewCoffeedash(address, client)

	/* 	length, _ := newinstance.GetStateCount(nil)

	   	for index := 0; index < int(length.Uint64()); index++ {
	   		state, values, err1 := newinstance.GetQtableState(nil, uint8(index))
	   		if err1 != nil {
	   			fmt.Printf("set error")
	   			log.Fatal(err)
	   		}
	   		val1, _ := strconv.ParseFloat(values[0], 64)
	   		val2, _ := strconv.ParseFloat(values[1], 64)

	   		fmt.Println("state", state, val1, val2)
	   	} */

	auth := bind.NewKeyedTransactor(privateKey) // in wei

	val := new(big.Int)
	val.SetString("1000000000000000000", 10)

	s := fmt.Sprintf("%f", -0.111111111111111)
	arrrrr := "0.34111" + s

	tx, err := newinstance.SetQValue(&bind.TransactOpts{
		From:   auth.From,
		Signer: auth.Signer,
		Value:  val,
	}, "coffee", arrrrr)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("tx sent: %s", tx.Hash().Hex())
}

func StartDeploy() {
	fmt.Println("StartDeploy")
	// connect to an ethereum node  hosted by infura
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("gasPrice: ", gasPrice)

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)         // in wei
	auth.GasLimit = uint64(3141500000) // in units
	auth.GasPrice = gasPrice

	_st := "coffee"
	_lr := big.NewInt(700000)
	_gm := big.NewInt(700000)
	_ep := big.NewInt(700000)
	_epd := big.NewInt(700000)
	_sr := big.NewInt(0)
	_pd := big.NewInt(0)

	address, tx, instance, err1 := cfd.DeployCoffeedash(
		auth,
		client, _st,
		_lr,
		_gm,
		_ep,
		_epd,
		_sr,
		_pd,
	)
	if err1 != nil {
		fmt.Println("err1: ", err1)
		log.Fatal(err)
	}

	fmt.Printf("Contract pending deploy: 0x%x\n", address)
	fmt.Println(address.Hex()) // 0x147B8eb97fD247D06C4006D269c90C1908Fb5D54
	fmt.Println(tx.Hash().Hex())

	_ = instance
}
