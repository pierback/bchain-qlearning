package usermanagement

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"runtime"
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

	_, filename, _, _ := runtime.Caller(0)
	usrs, _ := readUsers(path.Join(path.Dir(filename), "/usrs.txt"))
	fmt.Println("REad users \n", usrs)

	for _, usr := range usrs {
		ep, qt := db.GetQl(usr)
		fmt.Println("ep, qt: ", ep, qt)
		if ep != "" {
			ql := setQl(ep, qt)
			bcm.set(User{ethaddress: usr}, ql)
		}
	}
}

func setQl(ep, qt string) l.QLearning {
	ql := l.QLearning{}
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
	saveUser()

	return ql
}

func saveUser() {
	usrs := make([]string, 0, len(bcm.users))
	for k := range bcm.users {
		usrs = append(usrs, k.ethaddress)
	}
	_, filename, _, _ := runtime.Caller(0)
	writeUsers(usrs, path.Join(path.Dir(filename), "/usrs.txt"))
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

// readUsers reads a whole file into memory
// and returns a slice of its users.
func readUsers(path string) ([]string, error) {
	fmt.Println("readUsers: ", path)
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
	fmt.Println("path: ", path)
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
