package blockchain

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/binary"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	// "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	users "github.com/pierback/bchain-qlearning/internal/contracts/Store/Users.go"
)

func DeployContract() {
	client, err := ethclient.Dial("ws://192.168.178.34:8545")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("9fa16474b3bc9351a93987003034443ebfe4cbf4a7c3cd19d04d2c5ac76f6d0d")
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

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	_st := "coffee"
	_lr := float64ToByte(0.7)
	_gm := float64ToByte(0.8)
	_ep := float64ToByte(0.99)
	_epd := float64ToByte(0.9995)
	_sr := 0
	_pd := 0

	address, tx, _, err := users.DeployUsers(auth, client, _st, _lr, _gm, _ep, _epd, _sr, _pd)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(address.Hex()) // 0x147B8eb97fD247D06C4006D269c90C1908Fb5D54
	fmt.Println(tx.Hash().Hex())

	// _ = instance

	address := common.HexToAddress("0x077E839ACC2c1a141aC75A3Bb7b47E5e76758b5e")
	newinstance, err := store.NewUsers(address, client)
	fmt.Println("newinstance: ", newinstance)

	key := [32]byte{}
	value := [32]byte{}
	copy(key[:], []byte("foo"))
	copy(value[:], []byte("bar bar bar bar"))

	/* 	_, err1 := newinstance.SetItem(auth, key, value)
	   	if err1 != nil {
	   		fmt.Printf("set error")
	   		log.Fatal(err)
	   	}
	*/
	// fmt.Printf("tx sent: %s", tx.Hash().Hex()) // tx sent: 0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870

	result, err2 := newinstance.GetItems(nil, key)
	if err2 != nil {
		fmt.Printf("get error")
		log.Fatal(err2)
	}

	fmt.Println(string(result[:]))
}

func float64ToByte(f float64) []byte {
	var buf bytes.Buffer
	err := binary.Write(&buf, binary.BigEndian, f)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	return buf.Bytes()
}
