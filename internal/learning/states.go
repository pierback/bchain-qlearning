package learning

import (
	"fmt"
	"sync"
	"time"
)

//stateType type definition of a state
type stateType struct {
	Weekday    time.Weekday `json:"weekday"`
	Timeslot   timeslot     `json:"timeslot"`
	Drinkcount drinkcount   `json:"drinkcount"`
}

//UserState representing state of a user
type UserState stateType

//VirtualState representing state of dummy-user
type VirtualState stateType

//State interface methods
type State interface {
	Get() State
	toString() string
	Update(a Action) State
	New(dc drinkcount, _wd int, _ct float64) State
}

func (vs *VirtualState) getVS() *VirtualState {
	return vs
}

//Update virtualstate
func (vs VirtualState) Update(a Action) State {
	var ts timeslot
	var cc, wc, mc int

	ts = vs.Timeslot
	wd := vs.Weekday
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

//Update userstate
func (s UserState) Update(a Action) State {
	var ts timeslot

	var cc, wc, mc int

	ts = getCurrentTimeSlot(time.Now().Hour())
	wd := time.Now().Weekday()
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

func getTS(st State) timeslot {
	switch v := st.(type) {
	case VirtualState:
		return st.(VirtualState).Timeslot
	case UserState:
		return st.(UserState).Timeslot
	default:
		fmt.Printf("I don't know about type %T!\n", v)
		return -1
	}
}

func getCC(st State) int {
	switch v := st.(type) {
	case VirtualState:
		return st.(VirtualState).Drinkcount.CoffeeCount
	case UserState:
		return st.(UserState).Drinkcount.CoffeeCount
	default:
		fmt.Printf("I don't know about type %T!\n", v)
		return -1
	}
}

var wg sync.WaitGroup

//AddState to q-table
func (q *QLearning) AddState(s State) {
	if _, ok := q.qt[s]; !ok {
		// fmt.Println("					Set State", s)
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
	ts := getCurrentTimeSlot(time.Now().Hour())

	//if new day set drinkcount to 0
	if s.Weekday != wd {
		dc = drinkcount{CoffeeCount: 0, WaterCount: 0, MateCount: 0}
	} else {
		dc = s.Drinkcount
	}

	return UserState{Weekday: wd, Timeslot: ts, Drinkcount: dc}
}

//New returns a new Userstate object
func (s UserState) New(dc drinkcount, _wd int, _ct float64) State {
	var wd time.Weekday
	var ct timeslot

	if _wd == -1 {
		wd = time.Now().Weekday()
	} else {
		wd = time.Weekday(_wd)
	}

	if _ct == -1 {
		ct = getCurrentTimeSlot(time.Now().Hour())
	} else {
		ct = getCurrentTimeSlot(int(_ct))
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
	var wd time.Weekday
	var ct timeslot

	if _wd == -1 {
		wd = time.Now().Weekday()
	} else {
		wd = time.Weekday(_wd)
	}

	if _ct == -1 {
		ct = getCurrentTimeSlot(time.Now().Hour())
	} else {
		ct = getCurrentTimeSlot(int(_ct))
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

func NewState() State {
	wd := time.Now().Weekday()
	ct := getCurrentTimeSlot(time.Now().Hour())

	return UserState{
		Weekday:  wd,
		Timeslot: ct,
		Drinkcount: drinkcount{
			CoffeeCount: 0,
			WaterCount:  0,
			MateCount:   0,
		},
	}
}
