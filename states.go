package main

import (
	"fmt"
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
	Coffee int
	Water  int
	Mate   int
}
type StateSpace map[string]int

type State struct {
	Weekday    weekday
	Timeslot   timeslot
	Drinkcount drinkcount
}

func InitDrinksCountStates() []drinkcount {
	drinksStates := []drinkcount{}
	for i := 0; i < 8; i++ {
		// for j := 0; j < 4; j++ {
		// 	for k := 0; k < 5; k++ {
		// 		drinksStates = append(drinksStates, drinkcount{Coffee: i, Water: j, Mate: k})
		// 	}
		// }
		drinksStates = append(drinksStates, drinkcount{Coffee: i, Water: 0, Mate: 0})
	}
	return drinksStates
}

func InitStateSpace() StateSpace {
	drinksStates := InitDrinksCountStates()
	workdays := []weekday{Monday, Tuesday, Wednesday, Thursday, Friday}
	slots := []timeslot{Slot0, Slot1, Slot2, Slot3, Slot4, Slot5, Slot6}

	var index int
	statemap := make(StateSpace)

	for _, ds := range drinksStates {
		for _, sl := range slots {
			for _, wd := range workdays {
				statemap[(&State{wd, sl, ds}).String()] = index
				index++
			}
		}
	}
	fmt.Println("statemap: ", statemap, len(statemap))
	return statemap
}

// Returns Count of all Drinks  as string
// func (dc drinkcount) DrinksCountString() string {
// 	drinks := [...]string{
// 		"CoffeeCount",
// 		"MateCount",
// 		"WaterCount",
// 	}

// 	var drinkStr string
// 	for i := 0; i < len(drinks); i++ {
// 		drinkStr += drinks[i] + ": " + strconv.Itoa(dc[drinkcount(i)]) + " "
// 	}
// 	return drinkStr
// }

func (str *SSID) isEduroam() bool {
	if *str == "eduroam" {
		return true
	}
	return false
}

func (slot *timeslot) GetCurrentTimeSlot(ch int) {
	switch ch {
	case 7, 8:
		*slot = 0
		break
	case 9, 10:
		*slot = 1
		break
	case 11, 12:
		*slot = 2
		break
	case 13, 14:
		*slot = 3
		break
	case 15, 16:
		*slot = 4
		break
	case 17, 18:
		*slot = 5
		break
	case 19, 20, 21, 22, 23, 0, 1, 2, 3, 4, 5, 6:
		*slot = 6
		break
	default:
		*slot = -1
		break
	}
	// if ch >= 7 && ch < 9 {
	// 	*slot = 0
	// } else if ch >= 9 && ch < 11 {
	// 	*slot = 1
	// } else if ch >= 11 && ch < 13 {
	// 	*slot = 2
	// } else if ch >= 13 && ch < 15 {
	// 	*slot = 3
	// } else if ch >= 15 && ch < 17 {
	// 	*slot = 4
	// } else if ch >= 17 && ch < 19 {
	// 	*slot = 5
	// } else if ch > 19 && ch < 25 || ch >= 0 && ch < 6 {
	// 	*slot = 6
	// } else {
	// 	*slot = -1
	// }
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
