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
	Drinkcount Drinkcount   `json:"Drinkcount"`
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
	New(dc Drinkcount, _wd int, _ct float64) State
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
		Drinkcount: Drinkcount{
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

	ts = GetCurrentTimeSlot(time.Now().Hour())
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
		Drinkcount: Drinkcount{
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
		fmt.Printf("getTS I don't know about type %T! %v \n", v, st)
		return -1
	}
}

//CurrentWeekday returns current weekday based on state type
func CurrentWeekday(st State) time.Weekday {
	switch v := st.(type) {
	case VirtualState:
		return st.(VirtualState).Weekday
	case UserState:
		return time.Now().Weekday()
	default:
		fmt.Printf("CurrentWeekday I don't know about type %T! %v\n", v, st)
		return -1
	}
}

//GetWeekday returns weekday based on state type
// which is currently set in state
func GetWeekday(st State) time.Weekday {
	switch v := st.(type) {
	case VirtualState:
		return st.(VirtualState).Weekday
	case UserState:
		return st.(UserState).Weekday
	default:
		fmt.Printf("GetWeekday I don't know about type %T!\n", v)
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
		fmt.Printf("getCC I don't know about type %T! %v\n", v, st)
		return -1
	}
}

var wg sync.WaitGroup

//AddState to q-table
func (q *QLearning) AddState(s State) {
	if _, ok := q.Qt[s]; !ok {
		// fmt.Println("					Set State", s)
		q.Qt[s] = []float64{0, 0}
	}
}

//Get returns current Virtualstate
func (vs VirtualState) Get() State {
	return vs
}

//Get returns current UserState
func (s UserState) Get() State {
	var dc Drinkcount
	wd := time.Now().Weekday()
	ts := GetCurrentTimeSlot(time.Now().Hour())

	//if new day set Drinkcount to 0
	if s.Weekday != wd {
		dc = Drinkcount{CoffeeCount: 0, WaterCount: 0, MateCount: 0}
	} else {
		dc = s.Drinkcount
	}

	return UserState{Weekday: wd, Timeslot: ts, Drinkcount: dc}
}

//New returns a new Userstate object
// func (s UserState) New(dc Drinkcount, _wd int, _ct int) State {
func (s UserState) New(dc Drinkcount, _wd int, _ct float64) State {
	var wd time.Weekday
	var ct timeslot

	if _wd == -1 {
		wd = time.Now().Weekday()
	} else {
		wd = time.Weekday(_wd)
	}

	if _ct == -1 {
		ct = GetCurrentTimeSlot(time.Now().Hour())
	} else {
		ct = GetCurrentTimeSlot(int(_ct))
	}

	return UserState{
		Weekday:  wd,
		Timeslot: ct,
		Drinkcount: Drinkcount{
			CoffeeCount: dc.CoffeeCount,
			WaterCount:  dc.WaterCount,
			MateCount:   dc.MateCount,
		},
	}
}

//New returns a new VirtualState object
//ct currentime not slot
// func (vs VirtualState) New(dc Drinkcount, _wd int, _ct int) State {
func (vs VirtualState) New(dc Drinkcount, _wd int, _ct float64) State {
	var wd time.Weekday
	var ct timeslot

	if _wd == -1 {
		wd = time.Now().Weekday()
	} else {
		wd = time.Weekday(_wd)
	}

	if _ct == -1 {
		ct = GetCurrentTimeSlot(time.Now().Hour())
	} else {
		ct = GetCurrentTimeSlot(int(_ct))
	}

	return VirtualState{
		Weekday:  wd,
		Timeslot: ct,
		Drinkcount: Drinkcount{
			CoffeeCount: dc.CoffeeCount,
			WaterCount:  dc.WaterCount,
			MateCount:   dc.MateCount,
		},
	}
}

func NewState() State {
	wd := time.Now().Weekday()
	ct := GetCurrentTimeSlot(time.Now().Hour())

	return UserState{
		Weekday:  wd,
		Timeslot: ct,
		Drinkcount: Drinkcount{
			CoffeeCount: 0,
			WaterCount:  0,
			MateCount:   0,
		},
	}
}
