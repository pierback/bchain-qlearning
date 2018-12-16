package main

import (
	"fmt"
	"math/rand"
)

type Action int8

const (
	Nothing Action = iota
	Coffee
	Mate
	Water
)

type Training struct {
	vs         State
	stateActns trainingsdata
}

type QLearning struct {
	qt QTable

	state State
	actns int

	workdays int
	dc       drinkcount

	learningRate float64
	epsilon      float64
	gamma        float64

	reward float64

	train bool
	tr    Training

	sr successratio
}

type trainingsdata map[State]Action

type successratio struct {
	steps  int
	posr   int
	greedy int
}

//initLearner
func initLearner() {
	q := QLearning{}
	q.Initialize(true)
	q.Start()
}

// Initialize set init vals of qlearning object
func (q *QLearning) Initialize(training bool) {
	q.train = training

	q.learningRate = 0.2
	q.epsilon = 0.1
	q.gamma = 1

	q.actns = 2
	q.workdays = 5

	q.qt = make(QTable)
}

//Start kicks of qlearning proccess
func (q *QLearning) Start() {
	if q.train {
		var tdc int
		wts := GenerateTrainingSet()

		for reps := 0; reps < 3; reps++ {
			for d := 0; d < q.workdays; d++ {
				//drinkcount reset
				tdc = 0
				for sl := 7; sl < 19; sl += 2 {
					//filter all from trainingsset equals day and slot
					fsc := FilterSlice(wts[d], GetCurrentTimeSlot(sl))
					for st := 0; st < fsc+1; st++ {
						tdc += st
						q.tr.vs = NewState(drinkcount{CoffeeCount: tdc, WaterCount: 0, MateCount: 0}, d, float64(sl))
						fmt.Println(q.tr.vs)
						q.learn()
					}
				}
			}
		}
	}

	fmt.Println("Q-Table \n", q.qt)
	fmt.Println("successratio", q.sr)
}

//learn one iteration of qlearning proccess
func (q *QLearning) learn() {

	s := q.AddState(q.GetState())

	greedyAction := q.EpsilonGreedy(s)
	actionTook := q.UserMock()
	reward, newstate := q.TakeAction(greedyAction, actionTook)

	q.AddState(newstate)

	QVal := q.GetQ(greedyAction, q.GetState())
	MaxAction := q.GetAction(newstate)
	_QVal := q.GetQ(MaxAction, newstate)

	qval := QVal + q.learningRate*(reward+q.gamma*_QVal-QVal)
	q.SetQ(greedyAction, qval)

	q.state = newstate

	fmt.Println(" ")
}

//GetAction returns action with highest qval on given state
func (q *QLearning) GetAction(s State) Action {
	action := Action(0)
	max := q.qt[s][0]

	for i := 1; i < q.actns; i++ {
		if max < q.qt[s][Action(i)] {
			max = q.qt[s][Action(i)]
			action = Action(i)
		}
	}
	return action
}

//UserMock mocks user behavior
func (q *QLearning) UserMock() Action {
	return q.tr.stateActns[q.tr.vs]
}

//TakeAction exec given action and gets reward based on user action
func (q *QLearning) TakeAction(a Action, actionTook Action) (float64, State) {

	//socket connection: get which action user took
	reward := GetReward(a, actionTook)
	newstate := q.UpdateState(actionTook)

	q.sr.steps++
	if reward >= 0 {
		q.sr.posr++
		fmt.Println("	Right Action Predicted", a)
	}

	return reward, newstate
}

//EpsilonGreedy greedy-policy
func (q *QLearning) EpsilonGreedy(s State) Action {
	ran := rand.Float64() < 1-q.epsilon
	if ran {
		q.sr.greedy++
		a := q.GetAction(s)
		fmt.Println("		greedyAction: ", a)
		return a
	}

	ra := Action(rand.Intn(q.actns))
	fmt.Println("		randmom: ", ra)
	return ra
}

//GetReward returns reward based on
//calculated action and user feedback
func GetReward(a Action, feedback Action) float64 {
	if a == feedback {
		if a == Coffee {
			return 1
		}
		return 0
	}
	return -1
}

// GetQ returns qval of given state action pair
func (q *QLearning) GetQ(a Action, s State) float64 {
	return q.qt[s][int(a)]
}

// SetQ sets qval of given state action pair
func (q *QLearning) SetQ(a Action, qv float64) {
	s := q.GetState()
	q.qt[s][a] = qv
}

//GenerateTrainingSet returns action state pairs for mocking user
func GenerateTrainingSet() [][]float64 {
	mondayTimes := []float64{8.33, 10, 15}
	tuesdayTimes := []float64{8.49, 10.30, 12.30}
	wednesdayTimes := []float64{8.15, 10, 14.37}
	thursdayTimes := []float64{8.30, 9.30, 13, 15.20}
	fridayTimes := []float64{8.37, 10.15, 13.23, 15.57}

	return [][]float64{mondayTimes, tuesdayTimes, wednesdayTimes, thursdayTimes, fridayTimes}
	/*trs := map[State]Action{}

		for day, wt := range wts {
			for slot, time := range wt {
				s := NewState(drinkcount{CoffeeCount: slot, WaterCount: 0, MateCount: 0}, day, time)
				trs[s] = Coffee
			}
		}

	return trs*/
}
