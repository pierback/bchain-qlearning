package main

import (
	"fmt"
	"time"
)

type timeslot int

type SSID string

//Drink Enum
const (
	CoffeeCount int8 = iota
	MateCount
	WaterCount
)

//timeslot
const (
	Slot0 timeslot = 0
	Slot1 timeslot = 1
	Slot2 timeslot = 2
	Slot3 timeslot = 3
	Slot4 timeslot = 4
	Slot5 timeslot = 5
	Slot6 timeslot = 6
	Slot7 timeslot = -1
)

type weekday int8

//days
const (
	Monday weekday = iota
	Tuesday
	Wednesday
	Thursday
	Friday
)

type drinkcount struct {
	CoffeeCount int `json: "coffeecount"`
	WaterCount  int `json: "watercount"`
	MateCount   int `json: "matecount"`
}
type StateSpace map[string]int

type QTable map[State][]float64

type State struct {
	Weekday    weekday    `json: "weekday"`
	Timeslot   timeslot   `json: "timeslot"`
	Drinkcount drinkcount `json: "drinkcount"`
}

//virtual state methods
type VirtualState interface {
	getState() State
}

func InitDrinksCountStates() []drinkcount {
	var drinksStates []drinkcount
	for i := 0; i < 8; i++ {
		// for j := 0; j < 4; j++ {
		// 	for k := 0; k < 5; k++ {
		// 		drinksStates = append(drinksStates, drinkcount{Coffee: i, Water: j, Mate: k})
		// 	}
		// }
		drinksStates = append(drinksStates, drinkcount{CoffeeCount: i, WaterCount: 0, MateCount: 0})
	}
	return drinksStates
}

func (q *QLearning) InitStateSpace() {
	drinksStates := InitDrinksCountStates()
	workdays := []weekday{Monday, Tuesday, Wednesday, Thursday, Friday}
	slots := []timeslot{Slot0, Slot1, Slot2, Slot3, Slot4, Slot5, Slot6}
	statemap := make(QTable)

	for _, wd := range workdays {
		for _, sl := range slots {
			for _, ds := range drinksStates {
				st := (State{wd, sl, ds})
				statemap[st] = []float64{0, 0}
			}
		}
	}
	q.qt = statemap
}

// func (q *QLearning) GetState() State {
// 	if q.train {
// 		return q.tr.vs
// 	}
// 	return q.state
// }

func (q *QLearning) UpdateState(a Action) State {
	var ts timeslot
	var wd weekday
	var cc, wc, mc int

	if q.train {
		ts = q.tr.vs.Timeslot
		wd = q.tr.vs.Weekday
		cc = q.tr.vs.Drinkcount.CoffeeCount
		wc = q.tr.vs.Drinkcount.WaterCount
		mc = q.tr.vs.Drinkcount.MateCount
	} else {
		ts = GetCurrentTimeSlot(time.Now().Hour())
		wd = weekday(time.Now().Weekday())
		cc = q.state.Drinkcount.CoffeeCount
		wc = q.state.Drinkcount.WaterCount
		mc = q.state.Drinkcount.MateCount
	}

	if a == Coffee {
		cc++
	} else if a == Water {
		wc++
	} else if a == Mate {
		mc++
	}

	return State{
		Weekday:  wd,
		Timeslot: ts,
		Drinkcount: drinkcount{
			CoffeeCount: cc,
			WaterCount:  wc,
			MateCount:   mc,
		},
	}

}

func (str *SSID) isEduroam() bool {
	if *str == "eduroam" {
		return true
	}
	return false
}

func GetCurrentTimeSlot(ch int) timeslot {
	switch ch {
	case 7, 8:
		return 0
	case 9, 10:
		return 1
	case 11, 12:
		return 2
	case 13, 14:
		return 3
	case 15, 16:
		return 4
	case 17, 18:
		return 5
	case 19, 20, 21, 22, 23, 0, 1, 2, 3, 4, 5, 6:
		return 6
	default:
		return -1
	}
}

func (q *QLearning) AddState(s State) State {
	if _, ok := q.qt[s]; !ok {
		fmt.Println("					Set State", s)
		q.qt[s] = []float64{0, 0}
	}
	return s
}

/* func (q *QLearning) GetState() VirtualState {
	if q.train {
		return q.tr.vs
	}

	var dc drinkcount
	wd := time.Now().Weekday()
	ts := GetCurrentTimeSlot(time.Now().Hour())

	if q.state.Weekday != weekday(wd) {
		dc = drinkcount{CoffeeCount: 0, WaterCount: 0, MateCount: 0}
	} else {
		dc = q.dc
	}

	return State{Weekday: weekday(wd), Timeslot: ts, Drinkcount: dc}
} */

func (q *QLearning) GetState() State {
	if q.train {
		return q.tr.vs
	}

	var dc drinkcount
	wd := time.Now().Weekday()
	ts := GetCurrentTimeSlot(time.Now().Hour())

	if q.state.Weekday != weekday(wd) {
		dc = drinkcount{CoffeeCount: 0, WaterCount: 0, MateCount: 0}
	} else {
		dc = q.dc
	}

	return State{Weekday: weekday(wd), Timeslot: ts, Drinkcount: dc}
}

func (day weekday) String() string {
	names := [...]string{
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
	}

	if day < Monday || day > Friday {
		return "Unknown"
	}
	return names[day]
}

// TimeSlotString
func (curTime timeslot) TimeSlotString() string {
	slots := [...]string{
		"Slot0",
		"Slot1",
		"Slot2",
		"Slot3",
		"Slot4",
		"Slot5",
		"Slot6",
	}
	if curTime < Slot0 || curTime > Slot6 {
		return "Unknown"
	}
	return slots[curTime]
}

func NewState(dc drinkcount, _wd int, _ct float64) State {
	var wd weekday
	var ct timeslot

	if _wd == -1 {
		wd = weekday(time.Now().Weekday())
		fmt.Println("wd: ", wd)
	} else {
		wd = weekday(_wd)
	}

	if _ct == -1 {
		ct = GetCurrentTimeSlot(time.Now().Hour())
	} else {
		ct = GetCurrentTimeSlot(int(_ct))
	}

	return State{
		Weekday:  wd,
		Timeslot: ct,
		Drinkcount: drinkcount{
			CoffeeCount: dc.CoffeeCount,
			WaterCount:  dc.WaterCount,
			MateCount:   dc.MateCount,
		},
	}
}
