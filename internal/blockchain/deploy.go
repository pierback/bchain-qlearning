package blockchain

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/binary"
	"fmt"
	"log"
	"math/big"
	"strconv"
	"unsafe"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	cfd "github.com/pierback/bchain-qlearning/internal/contracts/Store/CoffeeDash"
)

func DeployContract() {
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

	/* _st := "coffee"
	_lr := blabla
	_gm := blabla  //*byte32(float64ToByte(0.8))
	_ep := blabla  //*byte32(float64ToByte(0.99))
	_epd := blabla //*byte32(float64ToByte(0.9995))
	_sr := big.NewInt(0)
	_pd := big.NewInt(0)

	address, tx, instance, err := cfd.DeployCoffeedash(auth, client, _st, _lr, _gm, _ep, _epd, _sr, _pd)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(address.Hex()) // 0x147B8eb97fD247D06C4006D269c90C1908Fb5D54
	fmt.Println(tx.Hash().Hex())

	_ = instance */

	address := common.HexToAddress("0x28058Be156Fa5462706768cFb54931Ff6aFa4Aac")
	newinstance, err := cfd.NewCoffeedash(address, client)

	key := [32]byte{}
	value := [32]byte{}
	copy(key[:], []byte("foo"))
	copy(value[:], []byte("bar bar bar bar"))

	length, _ := newinstance.GetStateCount(nil)

	for index := 0; index < int(length.Uint64()); index++ {
		state, values, err1 := newinstance.GetQtableState(nil, big.NewInt(int64(index)))
		if err1 != nil {
			fmt.Printf("set error")
			log.Fatal(err)
		}
		val1, _ := strconv.ParseFloat(values[0], 64)
		val2, _ := strconv.ParseFloat(values[1], 64)

		fmt.Println("state", state, val1, val2)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)        // in wei
	auth.GasLimit = uint64(80000)     // in units 3000000
	auth.GasPrice = big.NewInt(50957) //gasPrice
	fmt.Println("gasPrice: ", gasPrice)
	_ = gasPrice
	s := fmt.Sprintf("%f", -0.111111111111111)
	arrrrr := [2]string{"0.34111", s}

	tx, err := newinstance.SetQValue(auth, "coffee", arrrrr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("tx sent: %s", tx.Hash().Hex())
}

func float64ToByte(f float64) []byte {
	var buf bytes.Buffer
	err := binary.Write(&buf, binary.BigEndian, f)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	return buf.Bytes()
}

func byte32(s []byte) (a *[32]byte) {
	if len(a) <= len(s) {
		a = (*[len(a)]byte)(unsafe.Pointer(&s[0]))
	}
	return a
}
