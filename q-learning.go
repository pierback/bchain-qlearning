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
	Q [][]float64

	qt QTable

	state    State
	actns    int
	statemap map[string]float64

	workdays int
	slots    int
	dc       drinkcount

	stateSpaceLength int

	learningRate float64
	epsilon      float64
	gamma        float64

	reward      float64
	prevstate   State
	preveaction Action

	train bool

	tr Training

	sr successratio
}

type trainingsdata map[State]Action

type successratio struct {
	steps  int
	posr   int
	greedy int
}

var sr successratio

var wts [][]float64

func GenerateTrainingSet() trainingsdata {
	mondayTimes := []float64{8.33, 10, 15}
	tuesdayTimes := []float64{8.49, 10.30, 12.30}
	wednesdayTimes := []float64{8.15, 10, 14.37}
	thursdayTimes := []float64{8.30, 9.30, 13, 15.20}
	fridayTimes := []float64{8.37, 10.15, 13.23, 15.57}

	wts = [][]float64{mondayTimes, tuesdayTimes, wednesdayTimes, thursdayTimes, fridayTimes}
	trs := map[State]Action{}

	for day, wt := range wts {
		for slot, time := range wt {
			s := NewState(drinkcount{CoffeeCount: slot, WaterCount: 0, MateCount: 0}, day, time)
			trs[s] = Coffee
		}
	}

	return trs
}

func (q *QLearning) Initialize(training bool) {
	q.train = training

	q.learningRate = 0.3
	q.epsilon = 0.1
	q.gamma = 1

	q.actns = 2
	q.slots = 7
	q.workdays = 5

	q.preveaction = 0

	q.state = q.GetState()

	q.qt = make(QTable)

	if training {
		q.tr.stateActns = GenerateTrainingSet()
	}
}

func initLearner() {
	q := QLearning{}
	q.Initialize(true)
	q.Start()
}

func (q *QLearning) Start() {
	q.train = true
	var tdc int

	for reps := 0; reps < 12; reps++ {
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

	fmt.Println("len(q.statemap)", q.qt)
	fmt.Println("successratio", q.sr)
	// fmt.Println("q.tr.stateActns", q.tr.stateActns)
}

func (q *QLearning) learn() {

	// q.learningRate = float64(1 / (1 + q.sr.steps))

	prevQ := q.GetQ(q.preveaction)

	greedyAction := q.EpsilonGreedy()
	actionTook := q.UserMock()
	reward, newstate := q.TakeAction(greedyAction, actionTook)
	q.SetState(*newstate)

	curQ := q.GetQ(greedyAction)
	qval := curQ + q.learningRate*(reward+q.gamma*prevQ-curQ)
	q.SetQ(greedyAction, qval)

	q.preveaction = greedyAction
	q.prevstate = q.state
}

func (q *QLearning) GetAction() Action {
	q.state = q.GetState()
	action := Action(0)
	max := q.qt[q.state][action]

	for i := 1; i < q.actns; i++ {
		if max < q.qt[q.state][Action(i)] {
			max = q.qt[q.state][Action(i)]
			action = Action(i)
		}
	}
	return action
}

func (q *QLearning) UserMock() Action {
	ma := q.tr.stateActns[q.tr.vs]
	fmt.Println("		UserMock:     ", ma)
	return ma
}

func (q *QLearning) TakeAction(a Action, actionTook Action) (float64, *State) {

	//socket connection: get which action user took
	reward := GetReward(a, actionTook)
	newstate := q.UpdateState(actionTook)

	q.sr.steps++
	if reward >= 0 {
		q.sr.posr++
		fmt.Println("	Right Action Predicted", a, actionTook)
	}

	return reward, newstate
}

func (q *QLearning) EpsilonGreedy() Action {
	if rand.Float64() < q.epsilon {
		ra := Action(rand.Intn(q.actns))
		return ra
	}
	q.sr.greedy++
	a := q.GetAction()
	fmt.Println("		greedyAction: ", a)
	return a
}

func GetReward(a Action, feedback Action) float64 {
	if a == feedback {
		if a == Coffee {
			return 1
		}
		return 0
	}
	return -1
}

func (q *QLearning) GetQ(a Action) float64 {
	s := q.GetState()
	if _, ok := q.qt[s]; !ok {
		q.qt[s] = []float64{0, 0}
	}
	return q.qt[s][int(a)]
}

func (q *QLearning) SetQ(a Action, f float64) {
	s := q.GetState()
	if _, ok := q.qt[s]; !ok {
		q.qt[s] = []float64{0, 0}
	}
	q.qt[s][a] = f
}
