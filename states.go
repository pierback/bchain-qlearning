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

//drinkcount of a user
type drinkcount struct {
	CoffeeCount int `json:"coffeeCount"`
	WaterCount  int `json:"waterCount"`
	MateCount   int `json:"mateCount"`
}

//StateStruct type definition of a state
type StateStruct struct {
	Weekday    weekday    `json:"weekday"`
	Timeslot   timeslot   `json:"timeslot"`
	Drinkcount drinkcount `json:"drinkcount"`
}

//UserState representing state of a user
type UserState StateStruct

//VirtualState representing state of dummy-user
type VirtualState StateStruct

//State interface methods
type State interface {
	Get() State
	toString() string
	Update(a Action) State
	New(dc drinkcount, _wd int, _ct float64) State
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
				st := (&UserState{wd, sl, ds})
				statemap[st] = []float64{0, 0}
			}
		}
	}
	q.qt = statemap
}

// func (q *QLearning) Get() State {
// 	if q.train {
// 		return q.tr.vs
// 	}
// 	return q.state
// }

func (vs VirtualState) Update(a Action) State {
	var ts timeslot
	var wd weekday
	var cc, wc, mc int

	ts = vs.Timeslot
	wd = vs.Weekday
	cc = vs.Drinkcount.CoffeeCount
	wc = vs.Drinkcount.WaterCount
	mc = vs.Drinkcount.MateCount

	if a == Coffee {
		cc++
	} else if a == Water {
		wc++
	} else if a == Mate {
		mc++
	}

	return VirtualState{
		Weekday:  wd,
		Timeslot: ts,
		Drinkcount: drinkcount{
			CoffeeCount: cc,
			WaterCount:  wc,
			MateCount:   mc,
		},
	}

}

func (s UserState) Update(a Action) State {
	var ts timeslot
	var wd weekday
	var cc, wc, mc int

	ts = GetCurrentTimeSlot(time.Now().Hour())
	wd = weekday(time.Now().Weekday())
	cc = s.Drinkcount.CoffeeCount
	wc = s.Drinkcount.WaterCount
	mc = s.Drinkcount.MateCount

	if a == Coffee {
		cc++
	} else if a == Water {
		wc++
	} else if a == Mate {
		mc++
	}

	return UserState{
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

func (q *QLearning) AddState(s State) {
	if _, ok := q.qt[s]; !ok {
		fmt.Println("					Set State", s)
		q.qt[s] = []float64{0, 0}
	}
}

//Get returns current Virtualstate
func (vs VirtualState) Get() State {
	return vs
}

//Get returns current UserState
func (s UserState) Get() State {
	var dc drinkcount
	wd := time.Now().Weekday()
	ts := GetCurrentTimeSlot(time.Now().Hour())

	if s.Weekday != weekday(wd) {
		dc = drinkcount{CoffeeCount: 0, WaterCount: 0, MateCount: 0}
	} else {
		dc = s.Drinkcount
	}

	return UserState{Weekday: weekday(wd), Timeslot: ts, Drinkcount: dc}
}

//String converts weekday to string
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

//New returns a new Userstate object
func (s UserState) New(dc drinkcount, _wd int, _ct float64) State {
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

	return UserState{
		Weekday:  wd,
		Timeslot: ct,
		Drinkcount: drinkcount{
			CoffeeCount: dc.CoffeeCount,
			WaterCount:  dc.WaterCount,
			MateCount:   dc.MateCount,
		},
	}
}

//New returns a new VirtualState object
func (vs VirtualState) New(dc drinkcount, _wd int, _ct float64) State {
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

	return VirtualState{
		Weekday:  wd,
		Timeslot: ct,
		Drinkcount: drinkcount{
			CoffeeCount: dc.CoffeeCount,
			WaterCount:  dc.WaterCount,
			MateCount:   dc.MateCount,
		},
	}
}
