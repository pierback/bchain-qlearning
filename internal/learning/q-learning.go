package learning

import (
	"fmt"
	"math/rand"
	"time"
)

type Action int8

const (
	Nothing Action = iota
	Coffee
	Mate
	Water
)

type QLearning struct {
	qt QTable

	state State
	actns int

	workdays int
	dc       drinkcount

	learningRate float64
	epsilon      float64
	epsilonDecay float64

	gamma float64

	reward float64
	sr     successratio

	prediction Action
}

// QTable type
type QTable map[State][]float64

type QLearner interface {
	InitLearner()
}

type SimulatedStateActions map[State]Action

// Initialize set init vals of qlearning object
func (q *QLearning) Initialize() {
	q.learningRate = 0.7
	q.epsilon = 0.99

	q.gamma = 0.8

	q.actns = 2
	q.workdays = 5

	q.epsilonDecay = 0.9995

	q.qt = make(QTable)
}

//learn one iteration of qlearning proccess
func (q *QLearning) learn(fb Action) {
	//
	// q.epsilon = float64(1 / (q.sr.steps + 1))
	q.evalPrediction(fb)
	q.makePrediction()

	fmt.Println("   ")
}

func (q *QLearning) evalPrediction(fb Action) {
	qsa := q.GetQ(q.prediction, q.getState())
	reward, newstate := q.TakeAction(q.prediction, fb)
	q.AddState(newstate)

	qval := qsa + q.learningRate*(reward-qsa)

	/* maxAction := q.GetAction(newstate)
		_qsa := q.GetQ(maxAction, newstate)

	  qval := qsa + q.learningRate*(reward+q.gamma*_qsa-qsa) */

	q.SetQ(q.prediction, qval)
	fmt.Println("eval prediction: ", q.state, "user: ", fb, "ql: ", q.prediction)

	q.state = newstate
}

func (q *QLearning) makePrediction() {
	s := q.getState()

	q.prediction = q.EpsilonGreedy(s)

	q.epsilon *= q.epsilonDecay
	fmt.Println("make prediction: ", s, "ql:   ", q.prediction.String())
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

//TakeAction exec given action and gets reward based on user action
func (q *QLearning) TakeAction(a Action, actionTook Action) (float64, State) {
	reward := GetReward(a, actionTook)
	newstate := q.state.Update(actionTook)

	q.sr.steps++
	if reward < 0 {
		q.sr.neg++
	}

	return reward, newstate
}

//EpsilonGreedy greedy-policy
func (q *QLearning) EpsilonGreedy(s State) Action {
	if rand.Float64() < q.epsilon {
		act := rand.Intn(2)
		return Action(act)
	}

	a := q.GetAction(s)
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
	s := q.state.Get()
	// bc.SetQValue(stateToString(s), fmt.Sprintf("%f", qv))
	q.qt[s][a] = qv
}

func (q *QLearning) getState() State {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic panic panic")
			q.state = NewState()
		}
	}()

	if q.state == nil {
		fmt.Println("No panic")
		q.state = NewState()
	}

	s := q.state.Get()
	q.AddState(s)
	return s
}

func (q *QLearning) setNewState(vs VirtualState, d int, sl int) State {
	cnt := getCC(q.state)
	return vs.New(drinkcount{CoffeeCount: cnt, WaterCount: 0, MateCount: 0}, d, float64(sl))
}

func (q *QLearning) initStateSpace() {
	drinksStates := initDrinksCountStates()
	workdays := []time.Weekday{time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday}
	slots := []timeslot{Slot0, Slot1, Slot2, Slot3, Slot4, Slot5, Slot6}
	statemap := make(QTable)

	for _, wd := range workdays {
		for _, sl := range slots {
			for _, ds := range drinksStates {
				st := (&UserState{wd, sl, ds})
				statemap[st] = []float64{0, 0}
			}
		}
	}
	q.qt = statemap
}
