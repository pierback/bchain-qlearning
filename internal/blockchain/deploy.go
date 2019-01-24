package blockchain

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	cfd "github.com/pierback/bchain-qlearning/internal/contracts/Store/CoffeeDash"
)

const key = `{"address":"02e9f84165314bb8c255d8d3303b563b7375eb61","crypto":{"cipher":"aes-128-ctr","ciphertext":"f4952ba9725d5ae83f6e2c47714e4a1ed533a60a4ea97d635fd674e84b419f8a","cipherparams":{"iv":"98de0f7f0469a83ee3f543b232a9ec67"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"7629ecaf07277427a80e8e3138d9fc4058ca63b004f7ba3a2b253f22db55bced"},"mac":"b10cc5c56aa64f2246a5786ea3247ab6a133961b3b62985bc999dcc7cb51ade7"},"id":"3dfb452c-af90-41aa-aa8f-6da654b818e9","version":3}`

func DeployCommon() {
	// connect to an ethereum node  hosted by infura
	// blockchain, err := ethclient.Dial("http://127.0.0.1:8501")
	blockchain, err := ethclient.Dial("http://localhost:8501")

	if err != nil {
		log.Fatalf("Unable to connect to network:%v\n", err)
	}

	// Get credentials for the account to charge for contract deployments
	auth, err := bind.NewTransactor(strings.NewReader(key), "0000")

	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}
	_st := "coffee"
	_lr := big.NewInt(700000)
	_gm := big.NewInt(700000)
	_ep := big.NewInt(700000)
	_epd := big.NewInt(700000)
	_sr := big.NewInt(0)
	_pd := big.NewInt(0)

	address, _, instance, err1 := cfd.DeployCoffeedash(
		auth,
		blockchain, _st,
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

	length, _ := instance.GetStateCnt(nil)
	fmt.Println("length: ", length)

	_, err = instance.AddState(auth, "mate1", "-0.9999999999651321, -0.9999999999651321")

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
	// blockchain, err := ethclient.Dial("http://127.0.0.1:8501")
	blockchain, err := ethclient.Dial("http://0.0.0.0:8501")

	if err != nil {
		log.Fatalf("Unable to connect to network:%v\n", err)
	}
	auth, err := bind.NewTransactor(strings.NewReader(key), "0000")

	// projektmodul blockchain contracts addresses
	//cont := common.HexToAddress("0x76696ffcb3fad0fea8228e99766222339a4abd24")
	// cont := common.HexToAddress("0x5627e83b84e88965a5bc4d052e754448f3522098")

	//docker bchain adresses
	cont := common.HexToAddress("0xb7f2adbf4a6449af41f29a40b92b1b9334990e43")

	//
	instance, err1 := cfd.NewCoffeedash(cont, blockchain)
	if err1 != nil {
		fmt.Println("err1: ", err1)
		log.Fatal(err)
	}

	// _, err = instance.SetQValue(auth, "coffee", "-0.d, -d")
	newstate := "{Monday 0 {4 0 0}}"
	if output, _ := instance.StateExists(nil, newstate); !output {
		_, err = instance.AddState(auth, newstate, "-0.9999999956953279, -0.99999822853")
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
