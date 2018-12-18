package learning

import (
	"encoding/json"
	"fmt"
)

func FilterSlice(dayArr []float64, sl timeslot) int {
	var stateCount int //per slot
	for _, t := range dayArr {
		if GetCurrentTimeSlot(int(t)) == sl {
			stateCount++
		}
	}

	return stateCount
}

func stateToString(state State) string {
	out, err := json.Marshal(state)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return string(out)
}

func (us UserState) toString() string {
	return stateToString(us)
}

func (vs VirtualState) toString() string {
	return stateToString(vs)
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
