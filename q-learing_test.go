package main

import (
	"fmt"
	"testing"
)

func TestStartLearn(t *testing.T) {
	t.Parallel()
	fmt.Println("TestDrinksCount start")
	initLearner()

	fmt.Println("	")
}

func TestSetQ(t *testing.T) {
	t.Parallel()
	fmt.Println("TestSetQ start")
	q := QLearning{}
	ss := NewState(drinkcount{CoffeeCount: 0, WaterCount: 0, MateCount: 0}, int(Monday), 8.45)
	q.train = true
	q.qt = make(QTable)

	type Input struct {
		a Action
		v float64
		s State
	}

	var qavs = []struct {
		input    Input
		expected float64
	}{
		{input: Input{a: Coffee, v: -0.345, s: ss}, expected: -0.345},
		{input: Input{a: Nothing, v: -0.543, s: ss}, expected: -0.543},
		{input: Input{a: Nothing, v: -0.543, s: ss}, expected: -0.543},
	}

	fmt.Println("q.qt: ", q.qt)
	q.AddState(ss)

	for _, qav := range qavs {
		q.SetQ(qav.input.a, qav.input.v)

		if q.qt[ss][qav.input.a] != qav.expected {
			t.Error("Value does not match input val")
		}
	}
	fmt.Println("q.qt: ", q.qt)
}

func TestFilterSlice(t *testing.T) {
	t.Parallel()
	fmt.Println("TestFilterSlice start")

	type Input struct {
		tm []float64
		sl timeslot
	}

	var times = []struct {
		input    Input
		expected int
	}{
		{input: Input{tm: []float64{8.37, 10.15, 13.23, 14.57}, sl: timeslot(3)}, expected: 2},
		{input: Input{tm: []float64{8.37, 16.15, 15.23, 15.57}, sl: timeslot(4)}, expected: 3},
		{input: Input{tm: []float64{9.37, 10.15, 13.23, 14.57}, sl: timeslot(1)}, expected: 2},
		{input: Input{tm: []float64{1.37, 2.15, 3.23, 4.57}, sl: timeslot(6)}, expected: 4},
		{input: Input{tm: []float64{0, 0, 0, 0}, sl: timeslot(-1)}, expected: 0},
	}

	for _, ti := range times {
		if output := FilterSlice(ti.input.tm, ti.input.sl); output != ti.expected {
			t.Error("Wrong count", output, "Input", ti.input, ti.expected, ti)
		}
	}

}

func TestGetReward(t *testing.T) {
	t.Parallel()
	fmt.Println("Test Getreward start")

	var actionPairs = []struct {
		input    []Action
		expected float64
	}{
		{[]Action{Coffee, Nothing}, -1},
		{[]Action{Coffee, Coffee}, 1},
		{[]Action{Nothing, Coffee}, -1},
		{[]Action{Nothing, Nothing}, 0},
		{[]Action{Mate, Mate}, 0},
		{[]Action{Water, Water}, 0},
		{[]Action{Water, Coffee}, -1},
	}

	for _, ap := range actionPairs {
		if output := GetReward(ap.input[0], ap.input[1]); output != ap.expected {
			t.Error("Wrong reward given", output, "Input", ap.input, ap.expected, ap)
		}
	}

}

func TestUpdateState(t *testing.T) {
	t.Parallel()
	fmt.Println("TestUpdateState start")
	tq := QLearning{}
	sampleState := NewState(drinkcount{CoffeeCount: 0, WaterCount: 0, MateCount: 0}, -1, -1)

	var execActns = []struct {
		input    Action
		expected State
	}{
		{Coffee, NewState(drinkcount{CoffeeCount: 1, WaterCount: 0, MateCount: 0}, -1, -1)},
		{Nothing, NewState(drinkcount{CoffeeCount: 0, WaterCount: 0, MateCount: 0}, -1, -1)},
		{Mate, NewState(drinkcount{CoffeeCount: 0, WaterCount: 0, MateCount: 1}, -1, -1)},
		{Water, NewState(drinkcount{CoffeeCount: 0, WaterCount: 1, MateCount: 0}, -1, -1)},
		{4, NewState(drinkcount{CoffeeCount: 0, WaterCount: 0, MateCount: 0}, -1, -1)},
	}

	for _, ea := range execActns {
		tq.state = sampleState
		output := tq.UpdateState(ea.input)

		if output != ea.expected {
			t.Error("Wrong reward given", ea.input, ea.expected, ea)
		}
		fmt.Println("output: ", output.String())
	}

}

// func TestTakeAction(t *testing.T) {
// 	t.Parallel()
// 	fmt.Println("TestGetStateId start")
// 	var ssids = []struct {
// 		input    SSID
// 		expected [float64, State]
// 	}{
// 		{"eduroam", true},
// 		{"ipsum", false},
// 	}

// 	Q := QLearning{}
// 	// Q.InitStateSpace()
// 	Q.statemap = make(StateSpace)
// 	for index := 0; index < 6; index++ {
// 		sampleState := State{
// 			Weekday:  weekday(1),
// 			Timeslot: 1,
// 			Drinkcount: drinkcount{
// 				CoffeeCount: index,
// 				WaterCount:  0,
// 				MateCount:   0,
// 			},
// 		}
// 		Q.statemap[sampleState.String()] = index
// 	}

// }

// func TestInitStateSpace(t *testing.T) {

// }
// func TestGetState(t *testing.T) {

// }
// func TestUpdateState(t *testing.T) {

// }

// func TestSSID(t *testing.T) {
// 	t.Parallel()
// 	fmt.Println("TestSSID start")

// 	var ssids = []struct {
// 		input    SSID
// 		expected bool
// 	}{
// 		{"eduroam", true},
// 		{"ipsum", false},
// 	}

// 	for _, ssid := range ssids {
// 		if output := ssid.input.isEduroam(); output != ssid.expected {
// 			t.Error("Wrong wifi", ssid.input, ssid.expected, ssid)
// 		}
// 	}

// }
