package learning

import (
	"errors"
	"fmt"
	"sync"
)

type UserMap map[User]QLearning

type bchainMock struct {
	users UserMap
	mu    sync.RWMutex
}

func (bcm *bchainMock) Set(key User, data QLearning) {
	bcm.mu.Lock()
	defer bcm.mu.Unlock()
	bcm.users[key] = data
}

func (bcm *bchainMock) Get(key User) (QLearning, error) {
	bcm.mu.RLock()
	defer bcm.mu.RUnlock()
	ql, ok := bcm.users[key]
	if !ok {
		return QLearning{}, fmt.Errorf("The '%v' is not presented", key)
	}
	return ql, nil
}

func (bcm *bchainMock) getQlearning(ea ethaddress) (QLearning, error) {
	bcm.mu.RLock()
	defer bcm.mu.RUnlock()

	for usr, ql := range bcm.users {
		if usr.ea == ea {
			return ql, nil
		}
	}
	return QLearning{}, errors.New("no user found")
}

func (bcm *bchainMock) printUsers() {
	fmt.Println(bcm.users)
}
