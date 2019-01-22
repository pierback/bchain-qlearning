package blockchain

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	cfd "github.com/pierback/bchain-qlearning/internal/contracts/Store/CoffeeDash"
)

/*
func deploy(privateKey *ecdsa.PublicKey, client *ethclient.Client) *cfd.Coffeedash {

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	_st := "coffee"
	_lr := 700000
	_gm := 700000
	_ep := 700000
	_epd := 700000
	_sr := big.NewInt(0)
	_pd := big.NewInt(0)

	address, tx, instance, err := cfd.DeployCoffeedash(&auth, &client, _st, _lr, _gm, _ep, _epd, _sr, _pd)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(address.Hex()) // 0x147B8eb97fD247D06C4006D269c90C1908Fb5D54
	fmt.Println(tx.Hash().Hex())

	return instance
} */

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
