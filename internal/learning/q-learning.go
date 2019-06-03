package learning

import (
	"log"
	"math/rand"
	"time"

	en "github.com/pierback/bchain-qlearning/cmd/environment"
)

//Action type
type Action int8

//Actions a user can take
const (
	Nothing Action = iota
	Coffee
	Mate
	Water
)

//QLearning struct
type QLearning struct {
	Qt QTable `json:"qt"`

	State State `json:"state"`
	Actns int   `json:"actns"`

	Workdays int        `json:"workdays"`
	Dc       Drinkcount `json:"dc"`

	LearningRate float64 `json:"learningRate"`
	Epsilon      float64 `json:"epsilon"`
	EpsilonDecay float64 `json:"epsilonDecay"`

	Gamma float64 `json:"gamma"`

	Reward float64      `json:"reward"`
	Sr     Successratio `json:"sr"`

	Prediction Action `json:"prediction"`
}

type Successratio struct {
	Steps  int
	Neg    int
	Greedy int
	Wa     []int
}

// QTable type
type QTable map[State][]float64

type QLearner interface {
	InitLearner()
}

type SimulatedStateActions map[State]Action

// Initialize set init vals of qlearning object
func (q *QLearning) Initialize() {
	q.LearningRate = 0.7
	q.Epsilon = 0.99

	q.Gamma = 0.8

	q.Actns = 2
	q.Workdays = 5

	q.EpsilonDecay = 0.9995

	q.Qt = make(QTable)

	q.State = NewState()
}

func (q *QLearning) isNewDay() bool {
	curWd := CurrentWeekday(q.State)
	if GetWeekday(q.State) != curWd {
		return true
	}
	return false
}

//Learn one iteration of qlearning proccess
func (q *QLearning) Learn(fb Action) {
	//
	// q.Epsilon = float64(1 / (q.Sr.steps + 1))
	if q.isNewDay() {
		q.State = NewState()
	} else {
		//no need to eval prediction if new day
		//if not eval
		q.EvalPrediction(fb)
	}
	q.MakePrediction()
	// fmt.Println("   ")
}

func (q *QLearning) EvalPrediction(fb Action) {
	qsa := q.GetQ(q.Prediction, q.GetState())
	reward, newstate := q.TakeAction(q.Prediction, fb)
	q.AddState(newstate)

	qval := qsa + q.LearningRate*(reward-qsa)

	/* maxAction := q.GetAction(newstate)
		_qsa := q.GetQ(maxAction, newstate)

	  qval := qsa + q.LearningRate*(reward+q.Gamma*_qsa-qsa) */

	q.SetQ(q.Prediction, qval)
	if *en.SimFlag < 1 {
		// log.Printf("Make prediction: %s --> state: %s", q.Prediction, s.toString())
		log.Println("eval prediction: ", q.State, "user: ", fb, "ql: ", q.Prediction)
	}
	q.State = newstate
}

func (q *QLearning) MakePrediction() {
	s := q.GetState()
	q.Prediction = q.EpsilonGreedy(s)
	q.Epsilon *= q.EpsilonDecay

	if *en.SimFlag < 1 {
		// log.Printf("Watched User Action:  %s <-> Prediction: %s \n", fb, q.Prediction)
		log.Printf("Make prediction: %s --> state: %s \n", q.Prediction, s.toString())
	}
}

//GetAction returns action with highest qval on given state
func (q *QLearning) GetAction(s State) Action {
	action := Action(0)
	max := q.Qt[s][0]

	for i := 1; i < q.Actns; i++ {
		if max < q.Qt[s][Action(i)] {
			max = q.Qt[s][Action(i)]
			action = Action(i)
		}
	}
	return action
}

//TakeAction exec given action and gets reward based on user action
func (q *QLearning) TakeAction(a Action, actionTook Action) (float64, State) {
	reward := GetReward(a, actionTook)
	newstate := q.State.Update(actionTook)

	q.Sr.Steps++
	if reward < 0 {
		q.Sr.Neg++
	}

	return reward, newstate
}

//EpsilonGreedy greedy-policy
func (q *QLearning) EpsilonGreedy(s State) Action {
	if rand.Float64() < q.Epsilon {
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
	return q.Qt[s][int(a)]
}

// SetQ sets qval of given state action pair
func (q *QLearning) SetQ(a Action, qv float64) {

	// bc.SetQValue(stateToString(s), fmt.Sprintf("%f", qv))
	// q.Qt[s][a] = qv

	defer func() {
		if r := recover(); r != nil {
			log.Println("SetQ panic")
			q.State = NewState()
			q.AddState(q.State)
		}
	}()

	if q.State == nil {
		q.State = NewState()
	}

	s := q.State.Get()
	q.AddState(s)
	q.Qt[s][a] = qv
}

func (q *QLearning) GetState() State {
	defer func() {
		if r := recover(); r != nil {
			log.Println("panic panic panic")
			q.State = NewState()
		}
	}()

	if q.State == nil {
		log.Println("state is nill")
		q.State = NewState()
	}

	s := q.State.Get()
	q.AddState(s)
	return s
}

func (q *QLearning) SetState(d int, sl int) {
	// return vs.New(Drinkcount{CoffeeCount: cnt, WaterCount: 0, MateCount: 0}, d, sl)
	q.State.New(GetCurDrinkCount(q.State), int(d), float64(sl))
}

func (q *QLearning) SetNewUsrtState(vs VirtualState, d int, sl int) State {
	cnt := getCC(q.State)
	// return vs.New(Drinkcount{CoffeeCount: cnt, WaterCount: 0, MateCount: 0}, d, sl)
	return vs.New(Drinkcount{CoffeeCount: cnt, WaterCount: 0, MateCount: 0}, d, float64(sl))
}

func (q *QLearning) SetNewVirtState(vs VirtualState, d int, sl int) State {
	cnt := getCC(q.State)
	// return vs.New(Drinkcount{CoffeeCount: cnt, WaterCount: 0, MateCount: 0}, d, sl)
	return vs.New(Drinkcount{CoffeeCount: cnt, WaterCount: 0, MateCount: 0}, d, float64(sl))
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
	q.Qt = statemap
}
