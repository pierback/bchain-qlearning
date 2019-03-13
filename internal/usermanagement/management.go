package usermanagement

import (
	"fmt"
	"sync"

	l "github.com/pierback/bchain-qlearning/internal/learning"
)

//UserMap map of all user with qlearning struct as key
type UserMap map[User]l.QLearning

//UserManagement wrapper for usermap with Mutex
type UserManagement struct {
	users UserMap
	mu    sync.RWMutex
}

//GENERALUSER address
const GENERALUSER = "0x67656e6572616c75736572"

var (
	bcm  *UserManagement
	once sync.Once
)

//Worker for prediction evaluation
type Worker interface {
	run()
}

//BchainMock singleton constructor
func initBm() *UserManagement {
	once.Do(func() {
		bcm = &UserManagement{users: make(map[User]l.QLearning)}
		bcm.initUser(GENERALUSER)
	})

	return bcm
}

//Set new user in usermap
func (bcm *UserManagement) set(key User, ql l.QLearning) {
	bcm.mu.Lock()
	defer bcm.mu.Unlock()
	bcm.users[key] = ql
}

//Get user from usermap
func (bcm *UserManagement) get(key User) (l.QLearning, error) {
	bcm.mu.RLock()
	defer bcm.mu.RUnlock()
	ql, ok := bcm.users[key]
	if !ok {
		return l.QLearning{}, fmt.Errorf("The '%v' is not presented", key)
	}
	return ql, nil
}

func (bcm *UserManagement) getQlearning(ea string) l.QLearning {
	for usr, ql := range bcm.users {
		if usr.ethaddress == ea {
			return ql
		}
	}

	bcm.mu.RUnlock()
	fmt.Println("NEW USER:", ea)
	ql := bcm.initUser(ea)
	ql.Prediction = bcm.getGeneralPrediction(ql)

	return ql
}

func (bcm *UserManagement) initUser(ethaddrs string) l.QLearning {
	usr := User{ethaddrs}
	ql := usr.InitLearner()

	bcm.set(usr, ql)
	return ql
}

func (bcm *UserManagement) printUsers() {
	fmt.Println(bcm.users)
}

func (bcm *UserManagement) getGeneralPrediction(newUserQl l.QLearning) l.Action {
	ql := bcm.getQlearning(GENERALUSER)
	currentState := ql.GetState()
	return ql.GetAction(currentState)
}

func (bcm *UserManagement) SetGenQvals(qvsn float64, qvsc float64) {
	fmt.Printf("qvsn %f, qvsc %f \n", qvsn, qvsc)
	usrCnt := len(bcm.users)
	ql := bcm.getQlearning(GENERALUSER)
	qvalsNothing := qvsn / float64(usrCnt)
	qvalsCoffee := qvsc / float64(usrCnt)
	fmt.Printf("qvalsNothing %f, qvalsCoffee %f \n\n", qvalsNothing, qvalsCoffee)
	ql.SetQ(l.Nothing, qvalsNothing)
	ql.SetQ(l.Coffee, qvalsCoffee)
}
