package learning

type timeslot int
type ssid string

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

//Drinkcount of a user
type Drinkcount struct {
	CoffeeCount int `json:"coffeeCount"`
	WaterCount  int `json:"waterCount"`
	MateCount   int `json:"mateCount"`
}

func initDrinksCountStates() []Drinkcount {
	var drinksStates []Drinkcount
	for i := 0; i < 8; i++ {
		// for j := 0; j < 4; j++ {
		// 	for k := 0; k < 5; k++ {
		// 		drinksStates = append(drinksStates, Drinkcount{Coffee: i, Water: j, Mate: k})
		// 	}
		// }
		drinksStates = append(drinksStates, Drinkcount{CoffeeCount: i, WaterCount: 0, MateCount: 0})
	}
	return drinksStates
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

func (str *ssid) isEduroam() bool {
	if *str == "eduroam" {
		return true
	}
	return false
}

func TsBoundaries(ts uint) (int8, int8) {
	switch ts {
	case 0:
		return 7, 9
	case 1:
		return 10, 12
	case 2:
		return 13, 15
	case 3:
		return 16, 18
	case 4:
		return 19, 6
	default:
		return -1, -1
	}
}

//GetCurrentTimeSlot returns current timeslot
func GetCurrentTimeSlot(ch int) timeslot {
	switch ch {
	case 7, 8, 9:
		return 0
	case 10, 11, 12:
		return 1
	case 13, 14, 15:
		return 2
	case 16, 17, 18:
		return 3
	case 19, 20, 21, 22, 23, 0, 1, 2, 3, 4, 5, 6:
		return 4
	default:
		return -1
	}
}
