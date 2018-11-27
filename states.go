package main

import (
	"fmt"
	"strconv"
)

type TimeSlot int
type Weekday int8
type SSID string
type Drink int8

const (
	Coffee Drink = 0
	Mate   Drink = 1
	Water  Drink = 2
)

//timeslot
const (
	Slot0 TimeSlot = 0
	Slot1 TimeSlot = 1
	Slot2 TimeSlot = 2
	Slot3 TimeSlot = 3
	Slot4 TimeSlot = 4
	Slot5 TimeSlot = 5
	Slot6 TimeSlot = 6
	Slot7 TimeSlot = -1
)
const (
	Sunday    Weekday = 0
	Monday    Weekday = 1
	Tuesday   Weekday = 2
	Wednesday Weekday = 3
	Thursday  Weekday = 4
	Friday    Weekday = 5
	Saturday  Weekday = 6
)

type DrinksCount map[Drink]int

func (dc DrinksCount) DrinksCountString() string {
	drinks := [...]string{
		"Coffee",
		"Mate",
		"Water",
	}

	var drinkStr string
	for i := 0; i < len(drinks); i++ {
		drinkStr += drinks[i] + ": " + strconv.Itoa(dc[Drink(i)]) + " "
	}
	return drinkStr
}

func (str *SSID) isEduroam() bool {
	if *str == "eduroam" {
		return true
	}
	return false
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func contains(s []int, e int) bool {
	fmt.Println("s", s)
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func (slot *TimeSlot) GetCurrentTimeSlot(ch int) {
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

func (day Weekday) String() string {
	names := [...]string{
		"Sunday",
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
	}

	if day < Sunday || day > Saturday {
		return "Unknown"
	}
	return names[day]
}

func (curTime TimeSlot) TimeSlotString() string {
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
