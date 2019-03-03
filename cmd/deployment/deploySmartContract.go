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
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	bl "github.com/pierback/bchain-qlearning/internal/contracts/BeverageList"
	cc "github.com/pierback/bchain-qlearning/internal/contracts/CoffeeCoin"
)

const key = `{"address":"e8816898d851d5b61b7f950627d04d794c07ca37","crypto":{"cipher":"aes-128-ctr","ciphertext":"1ff4add6955cba7ddaf29f66d7d21c5e1d714ef6191fbc651ae60f2ea3c95e8f","cipherparams":{"iv":"3ff869fbdbe1a523cdb327780365976e"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"7372dbae5fb318f8684902e099c311d4188721d677974d729711762c7ef6030c"},"mac":"485fa5dc701067782baa1589716a53110c7f917eb259e35ebca7265bbb7150b1"},"id":"89edb004-5b00-4607-a3af-a0d9ab9b1c34","version":3}`

func DeploySC(args string) {
	client, err := ethclient.Dial("ws://0.0.0.0:8546")

	if err != nil {
		log.Fatalf("Unable to connect to network:%v\n", err)
	}

	auth, err := bind.NewTransactor(strings.NewReader(key), "password")

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

	err12 := ioutil.WriteFile("/var/tmp/bvrglst", address.Bytes(), 0644)

	check(err12)
}

func ccDeploy(auth *bind.TransactOpts, client *ethclient.Client) {
	chairAddress := common.HexToAddress("0x18ef96d887954472de5e9f47d60ba8dea371dbfe")
	coffeePrice := new(big.Int).SetUint64(3)
	matePrice := new(big.Int).SetUint64(5)
	waterPrice := new(big.Int).SetUint64(2)

	address, _, _, err1 := cc.DeployCoffeecoin(auth, client, chairAddress, coffeePrice, matePrice, waterPrice)
	if err1 != nil {
		fmt.Println("err1: ", err1)
		log.Fatal(err1)
	}

	fmt.Printf("Contract Coffe Coin pending deploy: 0x%x\n", address)

	err12 := ioutil.WriteFile("/var/tmp/cffcn", address.Bytes(), 0644)

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
