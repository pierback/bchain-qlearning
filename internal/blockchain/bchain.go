package blockchain

import (
	"fmt"
	"log" //"github.com/ethereum/go-ethereum/common/hexutil"

	// for demo

	userManagement "github.com/pierback/bchain-qlearning/internal/contracts/UserManagement" // for demo

	"github.com/ethereum/go-ethereum/common" //"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
)

func InsertNRetrieve() {

	client, err := ethclient.Dial("ws://192.168.178.34:8545")

	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress("0xab37fd5b395e8acaa7c0c7009f6a1f193eb03dba")
	instance, err := userManagement.NewUserManagement(address, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("contract is loaded")

	/* privateKey, err := crypto.HexToECDSA("9fa16474b3bc9351a93987003034443ebfe4cbf4a7c3cd19d04d2c5ac76f6d0d")
	if err != nil {
		log.Fatal(err)
	}

	account := common.HexToAddress("0x02e9f84165314bb8c255d8d3303b563b7375eb61")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance)

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
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
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice */

	email := [32]byte{}
	value := [32]byte{}
	copy(email[:], []byte("fabp_92@hotmail.de"))
	copy(value[:], []byte("bar"))

	ready, err := instance.Ready(nil)

	userCount, err := instance.GetUserCoffeeCnt(nil, email, 1, 1)

	fmt.Println(string(ready))
	fmt.Println("userCount", userCount)
}
