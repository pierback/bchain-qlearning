package blockchain

import (
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
)

var states []string

type bchainFunc func(*bind.TransactOpts, string, string) (*types.Transaction, error)
type bchainData struct {
	bcfunc bchainFunc
	auth   *bind.TransactOpts
	state  string
	vals   string
}

func (q *Queue) Worker() {
	if len(q.Values) > 0 {
		bcd := q.Dequeue()
		fmt.Println("bcd: ", bcd)
		// _, err := newinstance.AddState(bcd.auth, bcd.state, bcd.vals)
		_, err := bcd.bcfunc(bcd.auth, bcd.state, bcd.vals)
		if err != nil {
			fmt.Println("Failed to set new Qvals", err)
		}
		time.Sleep(5 * time.Second)
	} else {
		time.Sleep(5 * time.Second)
	}

	fmt.Println("Worker Worker Worker")
	fmt.Println(" ")
	fmt.Println(" ")
	q.Worker()
}

func (q *Queue) AddToQueue(ff bchainFunc, auth *bind.TransactOpts, state string, vals string) {
	fmt.Println("AddToQueue: ", state, vals)
	q.Enqueue(bchainData{ff, auth, state, vals})
}

const key = `{"address":"02e9f84165314bb8c255d8d3303b563b7375eb61","crypto":{"cipher":"aes-128-ctr","ciphertext":"f4952ba9725d5ae83f6e2c47714e4a1ed533a60a4ea97d635fd674e84b419f8a","cipherparams":{"iv":"98de0f7f0469a83ee3f543b232a9ec67"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"7629ecaf07277427a80e8e3138d9fc4058ca63b004f7ba3a2b253f22db55bced"},"mac":"b10cc5c56aa64f2246a5786ea3247ab6a133961b3b62985bc999dcc7cb51ade7"},"id":"3dfb452c-af90-41aa-aa8f-6da654b818e9","version":3}`

//SetQValue on given state on blockchain
func SetQValue(st string, vals string, wg *sync.WaitGroup) {
	defer wg.Done()
	instance := GetInstance()
	if output, _ := instance.StateExists(nil, st); !output {
		AddState(st, vals, wg)
	} else {
		auth, _ := bind.NewTransactor(strings.NewReader(key), "0000")
		// _, err := instance.SetQValue(auth, st, vals)

		jobQueue.AddToQueue(instance.SetQValue, auth, st, vals)
	}
}

//AddState to blockchain
func AddState(st string, vals string, wg *sync.WaitGroup) {
	defer wg.Done()
	newinstance := GetInstance()
	if output, _ := newinstance.StateExists(nil, st); !output {
		auth, _ := bind.NewTransactor(strings.NewReader(key), "0000")
		fmt.Println("adding NEW state", output, st)
		// _, err := newinstance.AddState(auth, st, vals)

		jobQueue.AddToQueue(newinstance.AddState, auth, st, vals)

	}
}

//GetStateCnt from blockchain
func GetStateCnt() uint8 {
	newinstance := GetInstance()
	length, err := newinstance.GetStateCnt(nil)
	if err != nil {
		log.Fatalf("Failed to get StateCnt %v", err)
	}
	return length
}
