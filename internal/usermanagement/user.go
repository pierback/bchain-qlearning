package usermanagement

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"

	l "github.com/pierback/bchain-qlearning/internal/learning"
)

type User struct {
	ethaddress string
}

type SimulatedUser struct {
	vsa l.SimulatedStateActions
}

type Training struct {
	l.State
	stateActns l.SimulatedStateActions
}

var wg sync.WaitGroup

func StartUserSimulation(userCnt int) {
	wg.Add(userCnt)
	for i := 0; i < userCnt; i++ {
		su := SimulatedUser{}
		go su.InitLearner()
	}
	wg.Wait()
}

//InitLearner for SimulatedUser
func (su *SimulatedUser) InitLearner() {
	defer wg.Done()
	q := l.QLearning{}
	q.Initialize()
	su.Start(&q)
}

//InitLearner for real user
func (u *User) InitLearner() *l.QLearning {
	q := &l.QLearning{}
	q.Initialize()

	tmpState := l.UserState{}
	q.State = tmpState.New(l.Drinkcount{}, -1, -1)

	return q
}

//Start kicks of qlearning proccess for user
func (u *User) Start(q *l.QLearning) {
	feedback := l.Nothing
	q.Learn(feedback)
}

//Start kicks of qlearning proccess for simulated user
func (su *SimulatedUser) Start(q *l.QLearning) {
	var wts [][]float64
	vs := l.VirtualState{}

	var wa, st []int
	rand.Seed(time.Now().UnixNano())
	seed := rand.Int63n(11164)

	fmt.Println("\n\nstart", seed)
	for reps := 0; reps < 250; reps++ {

		// q.State = vs.New(l.Drinkcount{CoffeeCount: 0, WaterCount: 0, MateCount: 0}, int(time.Monday), 7)

		// log.Println(" ")
		// log.Println(" ")

		wts, su.vsa = GenerateTrainingSet(seed)
		if reps == 0 || reps == 1 {
			fmt.Printf("\nwts: %v %d %d\n", wts, seed, reps)
		}

		for d := 1; d < q.Workdays+1; d++ {
			for sl := 7; sl < 19; sl += 3 {
				//filter all from trainingsset equals day and slot
				fsc := l.FilterSlice(wts[d-1], l.GetCurrentTimeSlot(sl))

				q.State = q.SetNewVirtState(vs, d, sl)
				q.MakePrediction()

				for st := 0; st < fsc+1; st++ {
					newState := q.SetNewVirtState(vs, d, sl)
					q.State = newState
					fb := su.UserMock(q.State)
					q.Learn(fb)
				}
				q.EvalPrediction(l.Nothing)
			}
		}
		wa = append(wa, q.Sr.Neg)
		st = append(st, q.Sr.Steps)
		// log.Println("Negs: ", wa)
		// log.Println("q.Sr.neg: ", q.Sr.neg)
		// log.Println(" ")
		q.Sr.Neg = 0
		q.Sr.Steps = 0
	}

	// writeJsonFile(MapToString(q.Qt))

	fmt.Printf("\n\nQ-Table %v \n", q.Qt)
	wass, _ := json.Marshal(wa)
	stepss, _ := json.Marshal(st)

	fmt.Printf("\n\nsuccessratio %v\n\n", strings.Trim(string(wass), "[]"))
	// fmt.Printf("\n\nstepss %v\n\n", strings.Trim(string(stepss), "[]"))
	_ = stepss
}

//UserMock mocks user behavior
func (su *SimulatedUser) UserMock(s l.State) l.Action {
	return su.vsa[s.Get()]
}

//GenerateTrainingSet returns action state pairs for mocking user
func GenerateTrainingSet(seed int64) ([][]float64, l.SimulatedStateActions) {
	rand.Seed(seed)

	mondayTimes := CreateDrinkinTimes()
	tuesdayTimes := CreateDrinkinTimes()
	wednesdayTimes := CreateDrinkinTimes()
	thursdayTimes := CreateDrinkinTimes()
	fridayTimes := CreateDrinkinTimes()

	wts := [][]float64{mondayTimes, tuesdayTimes, wednesdayTimes, thursdayTimes, fridayTimes}
	trs := map[l.State]l.Action{}

	ss := l.VirtualState{}

	for day, wt := range wts {
		for slot, clock := range wt {
			s := ss.New(l.Drinkcount{CoffeeCount: slot, WaterCount: 0, MateCount: 0}, day+1, clock)
			if time.Weekday(day+1) == time.Friday {
				// log.Println("state state state", s)
			}
			trs[s] = l.Coffee
		}
	}

	return wts, trs
}
