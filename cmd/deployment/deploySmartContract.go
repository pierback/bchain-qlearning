package deployment

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	bl "github.com/pierback/bchain-qlearning/internal/contracts/BeverageList"
	cc "github.com/pierback/bchain-qlearning/internal/contracts/CoffeeCoin"
)

const key = `{"address":"02e9f84165314bb8c255d8d3303b563b7375eb61","crypto":{"cipher":"aes-128-ctr","ciphertext":"f4952ba9725d5ae83f6e2c47714e4a1ed533a60a4ea97d635fd674e84b419f8a","cipherparams":{"iv":"98de0f7f0469a83ee3f543b232a9ec67"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"7629ecaf07277427a80e8e3138d9fc4058ca63b004f7ba3a2b253f22db55bced"},"mac":"b10cc5c56aa64f2246a5786ea3247ab6a133961b3b62985bc999dcc7cb51ade7"},"id":"3dfb452c-af90-41aa-aa8f-6da654b818e9","version":3}`

func DeploySC(args string) {
	client, err := ethclient.Dial("ws://127.0.0.1:8545")

	if err != nil {
		log.Fatalf("Unable to connect to network:%v\n", err)
	}

	auth, err := bind.NewTransactor(strings.NewReader(key), "0000")

	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	if args == "cffcn" {
		ccDeploy(auth, client)
	} else {
		bvglDeploy(auth, client)
	}
}

func bvglDeploy(auth *bind.TransactOpts, client *ethclient.Client) {
	address, _, _, err1 := bl.DeployBeveragelist(auth, client)
	if err1 != nil {
		fmt.Println("err1: ", err1)
		log.Fatal(err1)
	}

	fmt.Printf("Contract Beveragelist pending deploy: 0x%x\n", address)

	err12 := ioutil.WriteFile("/tmp/dat1", address.Bytes(), 0644)

	check(err12)
}

func ccDeploy(auth *bind.TransactOpts, client *ethclient.Client) {
	address, _, _, err1 := cc.DeployCoffeecoin(auth, client, new(big.Int).SetUint64(0))
	if err1 != nil {
		fmt.Println("err1: ", err1)
		log.Fatal(err1)
	}

	fmt.Printf("Contract Coffe Coin pending deploy: 0x%x\n", address)

	err12 := ioutil.WriteFile("/tmp/dat2", address.Bytes(), 0644)

	check(err12)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

//old approach with privatekey did not work
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

	/* _st := "coffee"
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

		_ = instance */
}
