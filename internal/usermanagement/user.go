package usermanagement

import (
	"log"
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

//InitLearner for SimulatedUser
func (su *SimulatedUser) InitLearner() {
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
	wts, su.vsa = GenerateTrainingSet()
	var wa []int

	log.Println("start")
	for reps := 0; reps < 100; reps++ {

		q.State = vs.New(l.Drinkcount{CoffeeCount: 0, WaterCount: 0, MateCount: 0}, int(time.Monday), float64(7))

		log.Println(" ")
		log.Println(" ")

		for d := 1; d < q.Workdays+1; d++ {
			for sl := 7; sl < 19; sl += 3 {
				//filter all from trainingsset equals day and slot
				fsc := l.FilterSlice(wts[d-1], l.GetCurrentTimeSlot(sl))

				q.State = q.SetNewState(vs, d, sl)
				q.MakePrediction()

				for st := 0; st < fsc+1; st++ {
					newState := q.SetNewState(vs, d, sl)
					q.State = newState
					fb := su.UserMock(q.State)
					q.Learn(fb)
				}
				q.EvalPrediction(l.Nothing)
			}
		}
		wa = append(wa, q.Sr.Neg)
		log.Println("Negs: ", wa)
		// log.Println("q.Sr.neg: ", q.Sr.neg)
		log.Println(" ")
		q.Sr.Neg = 0
	}

	// writeJsonFile(MapToString(q.Qt))

	log.Println("Q-Table \n", q.Qt)
	log.Println("successratio", wa)
}

//UserMock mocks user behavior
func (su *SimulatedUser) UserMock(s l.State) l.Action {
	return su.vsa[s.Get()]
}

//GenerateTrainingSet returns action state pairs for mocking user
func GenerateTrainingSet() ([][]float64, l.SimulatedStateActions) {
	mondayTimes := []float64{8.33, 10, 15}
	tuesdayTimes := []float64{8.49, 10.30, 12.30}
	wednesdayTimes := []float64{8.15, 10, 14.37}
	thursdayTimes := []float64{8.30, 9.30, 13, 15.20}
	fridayTimes := []float64{8.37, 10.15, 13.23, 15.57}

	wts := [][]float64{mondayTimes, tuesdayTimes, wednesdayTimes, thursdayTimes, fridayTimes}
	trs := map[l.State]l.Action{}

	ss := l.VirtualState{}

	for day, wt := range wts {
		for slot, clock := range wt {
			s := ss.New(l.Drinkcount{CoffeeCount: slot, WaterCount: 0, MateCount: 0}, day+1, clock)
			if time.Weekday(day+1) == time.Friday {
				log.Println("state state state", s)
			}
			trs[s] = l.Coffee
		}
	}

	return wts, trs
}
