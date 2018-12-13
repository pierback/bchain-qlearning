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
	vsm        []State
	vs         State
	stateActns trainingsdata
}

type QLearning struct {
	Q [][]float64

	state    State
	actns    int
	statemap map[string]int

	workdays       int
	slots          int
	maxCoffeeCount int
	maxMateCount   int
	maxWaterCount  int
	dc             drinkcount

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
			s := StateFactory(drinkcount{CoffeeCount: slot, WaterCount: 0, MateCount: 0}, day, time)
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
	q.maxCoffeeCount = 7
	q.maxMateCount = 4
	q.maxWaterCount = 4
	q.slots = 7
	q.workdays = 5

	q.preveaction = 0

	q.state = q.GetState()
	q.InitStateSpace()
	q.Q = make([][]float64, len(q.statemap))

	for i := 0; i < len(q.statemap); i++ {
		q.Q[i] = make([]float64, q.actns)
	}

	for i := 0; i < len(q.statemap); i++ {
		for j := 0; j < q.actns; j++ {
			q.Q[i][j] = 0
		}
	}
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

	for reps := 0; reps < 6; reps++ {
		for d := 0; d < q.workdays; d++ {
			//drinkcount reset
			tdc = 0
			for sl := 7; sl < 19; sl += 2 {
				//filter all from trainingsset equals day and slot
				fsc := FilterSlice(wts[d], GetCurrentTimeSlot(sl))
				for st := 0; st < fsc+1; st++ {
					tdc += st
					q.tr.vs = StateFactory(drinkcount{CoffeeCount: tdc, WaterCount: 0, MateCount: 0}, d, float64(sl))
					fmt.Println(q.tr.vs)
					q.learn()
				}
			}
		}
	}

	fmt.Println("len(q.statemap)", q.Q)
	fmt.Println("successratio", q.sr)
	// fmt.Println("q.tr.stateActns", q.tr.stateActns)
}

func (q *QLearning) learn() {

	// q.learningRate = float64(1 / (1 + q.sr.steps))

	greedyAction := q.EpsilonGreedy()
	actionTook := q.UserMock()
	reward, newstate := q.TakeAction(greedyAction, actionTook)

	prevQ := q.GetQ(q.preveaction, q.GetStateId())
	q.state = newstate
	q.tr.vs = newstate
	curQ := q.GetQ(greedyAction, q.GetStateId())

	qval := curQ + q.learningRate*(reward+q.gamma*prevQ-curQ)
	q.SetQ(greedyAction, qval)

	q.preveaction = greedyAction
	q.prevstate = newstate
}

func (q *QLearning) GetAction() Action {
	q.state = q.GetState()
	stateID := q.GetStateId()
	action := 0
	max := q.Q[stateID][action]

	for i := 1; i < q.actns; i++ {
		if max < q.Q[stateID][i] {
			max = q.Q[stateID][i]
			action = i
		}
	}
	return Action(action)
}

func (q *QLearning) UserMock() Action {
	ma := q.tr.stateActns[q.tr.vs]
	fmt.Println("	UserMock: ", ma)
	return ma
}

func (q *QLearning) TakeAction(a Action, actionTook Action) (float64, State) {

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
		q.sr.greedy++
		return ra
	}
	return q.GetAction()
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

func (q *QLearning) GetQ(a Action, stateId int) float64 {
	return q.Q[stateId][int(a)]
}

func (q *QLearning) SetQ(a Action, f float64) {
	stateId := q.GetStateId()
	q.Q[stateId][int(a)] = f
}
