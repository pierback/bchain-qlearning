package learning

import (
	"fmt"
	"time"
)

type ethaddress string

type User struct {
	ea ethaddress
}

type SimulatedUser struct {
	vsa SimulatedStateActions
}

type Training struct {
	State
	stateActns SimulatedStateActions
}

type successratio struct {
	steps  int
	neg    int
	greedy int
}

//InitLearner for SimulatedUser
func (su *SimulatedUser) InitLearner() {
	q := QLearning{}
	q.Initialize()
	su.Start(&q)
}

//InitLearner for real user
func (u *User) InitLearner() QLearning {
	q := QLearning{}
	q.Initialize()
	return q
}

//Start kicks of qlearning proccess for user
func (u *User) Start(q *QLearning) {
	feedback := Nothing
	q.learn(feedback)
}

//Start kicks of qlearning proccess for simulated user
func (su *SimulatedUser) Start(q *QLearning) {
	var wts [][]float64
	vs := VirtualState{}
	wts, su.vsa = GenerateTrainingSet()
	var wa []int

	fmt.Println("start")
	for reps := 0; reps < 100; reps++ {

		q.state = vs.New(drinkcount{CoffeeCount: 0, WaterCount: 0, MateCount: 0}, int(time.Monday), float64(7))

		fmt.Println(" ")
		fmt.Println(" ")

		for d := 1; d < q.workdays+1; d++ {
			for sl := 7; sl < 19; sl += 3 {
				//filter all from trainingsset equals day and slot
				fsc := FilterSlice(wts[d-1], getCurrentTimeSlot(sl))

				q.state = q.setNewState(vs, d, sl)
				q.makePrediction()

				for st := 0; st < fsc+1; st++ {
					newState := q.setNewState(vs, d, sl)
					q.state = newState
					fb := su.UserMock(q.state)
					q.learn(fb)
				}
				q.evalPrediction(Nothing)
			}
		}
		wa = append(wa, q.sr.neg)
		fmt.Println("Negs: ", wa)
		// fmt.Println("q.sr.neg: ", q.sr.neg)
		fmt.Println(" ")
		q.sr.neg = 0
	}

	// writeJsonFile(mapToString(q.qt))

	fmt.Println("Q-Table \n", q.qt)
	fmt.Println("successratio", wa)
}

//UserMock mocks user behavior
func (su *SimulatedUser) UserMock(s State) Action {
	return su.vsa[s.Get()]
}

//GenerateTrainingSet returns action state pairs for mocking user
func GenerateTrainingSet() ([][]float64, SimulatedStateActions) {
	mondayTimes := []float64{8.33, 10, 15}
	tuesdayTimes := []float64{8.49, 10.30, 12.30}
	wednesdayTimes := []float64{8.15, 10, 14.37}
	thursdayTimes := []float64{8.30, 9.30, 13, 15.20}
	fridayTimes := []float64{8.37, 10.15, 13.23, 15.57}

	wts := [][]float64{mondayTimes, tuesdayTimes, wednesdayTimes, thursdayTimes, fridayTimes}
	trs := map[State]Action{}

	ss := VirtualState{}

	for day, wt := range wts {
		for slot, clock := range wt {
			s := ss.New(drinkcount{CoffeeCount: slot, WaterCount: 0, MateCount: 0}, day+1, clock)
			if time.Weekday(day+1) == time.Friday {
				fmt.Println("state state state", s)
			}
			trs[s] = Coffee
		}
	}

	return wts, trs
}
