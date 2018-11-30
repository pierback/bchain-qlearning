package main

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
	Q [][]float64

	st    State
	actns int

	workdays       int
	slots          int
	maxCoffeeCount int
	maxMateCount   int
	maxWaterCount  int

	learningRate float64
	epsilon      float64
	gamma        float64
}

func (q *QLearning) Initialize() {

	q.learningRate = 0.5
	q.epsilon = 0.1
	q.gamma = 1

	Actions := 4

	var ts timeslot
	ts.GetCurrentTimeSlot(time.Now().Hour())
	wd := time.Now().Weekday()
	dc := drinkcount{Coffee: 0, Mate: 0, Water: 0}
	q.st = State{Weekday: weekday(wd), Timeslot: ts, Drinkcount: dc}
	fmt.Println("q.st: ", q.st.String())

	q.actns = Actions
	q.maxCoffeeCount = 7
	q.maxMateCount = 4
	q.maxWaterCount = 4
	q.slots = 7
	q.workdays = 5

	q.Q = make([][]float64, Actions)

	stateSpaceLength := q.maxCoffeeCount * q.maxMateCount * q.maxWaterCount * q.slots * q.workdays

	for i := 0; i < stateSpaceLength; i++ {
		for j := 0; j < Actions; j++ {
			q.Q[i][j] = 0
		}
	}

	// q.SetQAll(q.ter_n, q.ter_m, 0)

}

// func (q *QLearning) Pi() {
// 	for i := q.Sn - 1; i >= 0; i-- {
// 		for j := 0; j < q.Sm; j++ {
// 			if i == q.ter_n && j == q.ter_m {
// 				fmt.Print(" G")
// 			} else {
// 				switch q.GetAction(i, j) {
// 				case Coffee:
// 					fmt.Print(" Coffee")
// 				case Mate:
// 					fmt.Print(" Mate")
// 				case Water:
// 					fmt.Print(" Water")
// 				case Nothing:
// 					fmt.Print(" Nothing")
// 				}
// 			}
// 		}
// 		fmt.Println("")
// 	}

// 	fmt.Println("")
// }

func initLearner() {

	t, _ := time.Parse("2006 01 02 15 04", "2015 11 11 16 50")
	fmt.Println(t.YearDay()) // 315
	fmt.Println(t.Weekday()) // Wednesday

	t, _ = time.Parse("2006 01 02 15 04", "2011 01 01 0 00")
	fmt.Println(t.YearDay())
	fmt.Println(t.Weekday())

	rand.Seed(time.Now().Unix())
	Q := QLearning{}
	Q.Initialize()
	// Q.Start()
	// Q.Pi()
}

// func PrintAction(action int) {
// 	switch action {
// 	case Coffee:
// 		fmt.Print(" Coffee")
// 	case Mate:
// 		fmt.Print(" Mate")
// 	case Water:
// 		fmt.Print(" Water")
// 	case Nothing:
// 		fmt.Print(" Nothing")
// 	}
// }

// func (q *QLearning) Start() {

// 	episodes := 1000
// 	for i := 0; i < episodes; i++ {
// 		Sn := q.ini_n
// 		Sm := q.ini_m

// 		ep := 0

// 		for Sn != q.ter_n || Sm != q.ter_m {
// 			ep++
// Action := q.epsilon_greedy(Sn, Sm)
// r, _Sn, _Sm := q.TakeAction(Action, Sn, Sm)
// QSA := q.GetQ(Sn, Sm, Action)
// MaxAction := q.GetAction(_Sn, _Sm)
// _QSA := q.GetQ(_Sn, _Sm, MaxAction)

// Q := QSA + q.learningRate*(r+q.gamma*_QSA-QSA)
// q.SetQ(Sn, Sm, Action, Q)

// Sn = _Sn
// Sm = _Sm
// 		}
// 	}

// }

func (q *QLearning) GetAction(n, m int) int {

	// Idx := q.GetIdx(n, m)
	// max := q.Q[0][Idx]
	// Action := 0
	// for i := 1; i < q.Qn; i++ {
	// 	if max < q.Q[i][Idx] {
	// 		max = q.Q[i][Idx]
	// 		Action = i
	// 	}
	// }
	return 0
	// return Action

}

func (q *QLearning) epsilon_greedy(n, m int) int {

	Action := q.GetAction(n, m)

	if rand.Float64() < 1-q.epsilon {
		return Action
	}
	return 0
	// return rand.Intn(q.Qn)

}

// func (q *QLearning) TakeAction(a, n, m int) (float64, int, int) {

// 	_n := n
// 	_m := m

// 	switch a {
// 	case Coffee:
// 		if n != q.Sn-1 {
// 			_n = n + 1
// 		}
// 	case Mate:
// 		if n != 0 {
// 			_n = n - 1
// 		}
// 	case Water:
// 		if m != 0 {
// 			_m = m - 1
// 		}
// 	case Nothing:
// 		if m != q.Sm-1 {
// 			_m = m + 1
// 		}
// 	}

// 	if _n == q.ter_n && _m == q.ter_m {
// 		return 0.0, q.ter_n, q.ter_m
// 	} else if _n == 0 && _m >= 1 && _m < q.Sm-1 {
// 		return -100.0, 0, 0
// 	}

// 	return -1.0, _n, _m
// }

func (q *QLearning) GetQ(n, m, a int) float64 {

	return q.Q[a][q.GetIdx(n, m)]
}
func (q *QLearning) SetQ(n, m, a int, f float64) {
	q.Q[a][q.GetIdx(n, m)] = f
}

func (q *QLearning) GetIdx(n, m int) int {
	return 0
}
