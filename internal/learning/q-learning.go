package learning

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
	stateActns SimulatedStateActions
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
	sr     successratio
}

// QTable type
type QTable map[State][]float64

type QLearner interface {
	InitLearner()
}

type User struct {
	QLearning
}

type SimulatedUser struct {
	QLearning
	vsa SimulatedStateActions
}
type SimulatedStateActions map[State]Action

type successratio struct {
	steps  int
	neg    int
	greedy int
}

//initLearner
func (su *SimulatedUser) InitLearner() {
	q := QLearning{}
	q.Initialize()
	su.Start(&q)
}

func (u *User) InitLearner() {
	q := QLearning{}
	q.Initialize()
	// q.Start()
}

// Initialize set init vals of qlearning object
func (q *QLearning) Initialize() {
	q.learningRate = 0.7
	q.epsilon = 0.3

	q.actns = 2
	q.workdays = 5

	q.qt = make(QTable)
}

//Start kicks of qlearning proccess
func (su *SimulatedUser) Start(q *QLearning) {
	var tdc int
	var wts [][]float64
	vs := VirtualState{}
	wts, su.vsa = GenerateTrainingSet()

	var wa []int

	for reps := 0; reps < 1; reps++ {
		for d := 0; d < q.workdays; d++ {
			//drinkcount reset
			tdc = 0
			for sl := 7; sl < 19; sl++ {
				//filter all from trainingsset equals day and slot
				fsc := FilterSlice(wts[d], GetCurrentTimeSlot(sl))
				for st := 0; st < fsc+1; st++ {
					tdc += st
					q.state = vs.New(drinkcount{CoffeeCount: tdc, WaterCount: 0, MateCount: 0}, d, float64(sl))
					fb := su.UserMock(q.state)
					fmt.Println(q.state)
					q.learn(fb)
				}
			}
		}
		wa = append(wa /* q.sr.steps- */, q.sr.neg)
		q.sr.neg = 0
	}

	writeJsonFile(mapToString(q.qt))

	// fmt.Println("Q-Table \n", q.qt)
	fmt.Println("successratio", wa)
}

//learn one iteration of qlearning proccess
func (q *QLearning) learn(fb Action) {

	// q.learningRate = float64(1 / (q.sr.steps + 1))
	// if q.learningRate

	s := q.state.Get()
	q.AddState(s)

	greedyAction := q.EpsilonGreedy(s)
	reward, newstate := q.TakeAction(greedyAction, fb)
	q.AddState(newstate)

	// maxAction := q.GetAction(newstate)
	qsa := q.GetQ(greedyAction, s)
	// _qsa := q.GetQ(maxAction, newstate)

	qval := qsa + q.learningRate*(reward-qsa)
	q.SetQ(greedyAction, qval)

	q.state = newstate

	// fmt.Println(" ")
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
func (su *SimulatedUser) UserMock(s State) Action {
	return su.vsa[s.Get()]
}

//TakeAction exec given action and gets reward based on user action
func (q *QLearning) TakeAction(a Action, actionTook Action) (float64, State) {

	//socket connection: get which action user took
	reward := GetReward(a, actionTook)
	newstate := q.state.Update(actionTook)

	q.sr.steps++
	if reward >= 0 {
		fmt.Println("	Right Action Predicted", a)
	} else {
		q.sr.neg++
	}

	return reward, newstate
}

//EpsilonGreedy greedy-policy
func (q *QLearning) EpsilonGreedy(s State) Action {

	fmt.Println("		random: ", rand.Float64(), q.epsilon)
	if rand.Float64() < q.epsilon {
		act := rand.Intn(2)
		fmt.Println("		act", act, Action(act))
		return Action(act)
	}

	q.sr.greedy++
	a := q.GetAction(s)
	fmt.Println("		greedyAction: ", a)
	return a

}

//GetReward returns reward based on
//calculated action and user feedback
func GetReward(a Action, feedback Action) float64 {
	if a == feedback {
		return 1
	}
	return -1
}

// GetQ returns qval of given state action pair
func (q *QLearning) GetQ(a Action, s State) float64 {
	return q.qt[s][int(a)]
}

// SetQ sets qval of given state action pair
func (q *QLearning) SetQ(a Action, qv float64) {
	q.qt[q.state.Get()][a] = qv
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
		for slot, time := range wt {
			s := ss.New(drinkcount{CoffeeCount: slot, WaterCount: 0, MateCount: 0}, day, time)
			trs[s] = Coffee
		}
	}

	return wts, trs
}
