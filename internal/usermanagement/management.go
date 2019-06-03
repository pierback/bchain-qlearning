package usermanagement

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"

	l "github.com/pierback/bchain-qlearning/internal/learning"
	db "github.com/pierback/bchain-qlearning/pkg/database"
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
		initPrevQl()
	})

	return bcm
}

func initPrevQl() {
	bcm.mu.Lock()
	defer bcm.mu.Unlock()

	usrs := []string{
		"0x5409ed021d9299bf6814279a6a1411a7e866a631",
		"0x6ecbe1db9ef729cbe972c83fb886247691fb6beb",
		"0xe36ea790bc9d7ab70c55260c66d52b1eca985f84",
		"0xe834ec434daba538cd1b9fe1582052b880bd7e63",
		"0x78dc5d2d739606d31509c31d654056a45185ecb6",
		"0x06cef8e666768cc40cc78cf93d9611019ddcb628",
		"0x4404ac8bd8f9618d27ad2f1485aa1b2cfd82482d",
		"0x7457d5e02197480db681d3fdf256c7aca21bdc12",
		GENERALUSER,
	}

	for _, usr := range usrs {
		result := db.GetQl(usr)
		fmt.Printf("\n User: %s \n\n", usr)
		if result["qt"] != nil {
			qt := result["qt"].(string)
			ep := result["ep"].(string)
			// wa := result["Wk_negs"].([]string)
			// fmt.Println("result: ", wa)
			/* negs := result["negs"].(int)
			steps := result["steps"].(int) */

			// sr := l.Successratio{Steps: steps, Neg: negs, Wa: wa}

			ql := setQl(ep, qt)
			// ql.Sr = sr
			log.Printf("\n User: %s Epsilon: %f \n Qtable: %v \n\n", usr, ql.Epsilon, ql.Qt)
			bcm.set(User{ethaddress: usr}, *ql)
		}
	}
}

/*
func UnmarshalJSON(b []byte) error {
	// Try string first
	var s string
	if err := json.Unmarshal(b, &s); err == nil {
		value, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return err
		}
		*i = int(value)
		return nil
	}

	// Fallback to number
	return json.Unmarshal(b, (*int)(i))
} */

func setQl(ep, qt string) *l.QLearning {
	ql := &l.QLearning{}
	ql.Initialize()
	if ep != "" {
		ql.Epsilon, _ = strconv.ParseFloat(ep, 64)
	}

	if qt != "" {
		ql.Qt.StringToMap(qt)
	}

	return ql
}

//Set new user in usermap
func (bcm *UserManagement) set(key User, ql l.QLearning) {
	bcm.users[key] = ql
}

//Get user from usermap
func (bcm *UserManagement) get(key User) (l.QLearning, error) {
	/* 	bcm.mu.RLock()
	   	defer bcm.mu.RUnlock() */
	ql, ok := bcm.users[key]
	if !ok {
		return l.QLearning{}, fmt.Errorf("The '%v' is not presented", key)
	}
	return ql, nil
}

func (bcm *UserManagement) getQlearning(ea string) *l.QLearning {
	for usr, ql := range bcm.users {
		if usr.ethaddress == ea {
			return &ql
		}
	}

	bcm.mu.RUnlock()
	log.Println("NEW USER:", ea)
	ql := bcm.initUser(ea)
	ql.Qt = bcm.getGeneralQt()

	return ql
}

func (bcm *UserManagement) initUser(ethaddrs string) *l.QLearning {
	usr := User{ethaddrs}
	ql := usr.InitLearner()

	bcm.set(usr, *ql)
	// saveUser()

	return ql
}

func saveUser() {
	usrs := make([]string, 0, len(bcm.users))
	for k := range bcm.users {
		usrs = append(usrs, k.ethaddress)
	}
	// _, filename, _, _ := runtime.Caller(0)
	// writeUsers(usrs, path.Join(path.Dir(filename), "/usrs.txt"))
}

func (bcm *UserManagement) printUsers() {
	log.Println("Users", bcm.users)
}

func (bcm *UserManagement) getGeneralPrediction(newUserQl l.QLearning) l.Action {
	ql := bcm.getQlearning(GENERALUSER)
	currentState := ql.GetState()
	return ql.GetAction(currentState)
}

func (bcm *UserManagement) getGeneralQt() l.QTable {
	ql := bcm.getQlearning(GENERALUSER)
	return ql.Qt
}

func (bcm *UserManagement) SetGenQvals(qvsn float64, qvsc float64) {
	// log.Printf("qvsn %f, qvsc %f \n", qvsn, qvsc)
	usrCnt := len(bcm.users)
	ql := bcm.getQlearning(GENERALUSER)
	qvalsNothing := qvsn / float64(usrCnt)
	qvalsCoffee := qvsc / float64(usrCnt)
	// log.Printf("qvalsNothing %f, qvalsCoffee %f \n\n", qvalsNothing, qvalsCoffee)
	ql.SetQ(l.Nothing, qvalsNothing)
	ql.SetQ(l.Coffee, qvalsCoffee)
}

// readUsers reads a whole file into memory
// and returns a slice of its users.
func readUsers(path string) ([]string, error) {
	log.Println("readUsers: ", path)
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var users []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		users = append(users, scanner.Text())
	}
	return users, scanner.Err()
}

// writeUsers writes the users to the given file.
func writeUsers(users []string, path string) error {
	log.Println("path: ", path)
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range users {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}
