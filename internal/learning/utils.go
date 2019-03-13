package learning

import (
	"encoding/json"
	"fmt"
	"os"
	"os/user"
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

func MapToString(Qt QTable) []byte {
	sm := map[string][]float64{}

	for st, v := range Qt {
		sm[st.toString()] = v
	}

	jsonData, err := json.Marshal(sm)
	if err != nil {
		panic(err)
	}
	return jsonData
}

func writeJsonFile(jsonData []byte) {
	usr, err := user.Current()

	if err != nil {
		panic(err)
	}

	jsonFile, err := os.Create(usr.HomeDir + "/qtable.json")

	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	jsonFile.Write(jsonData)
	jsonFile.Close()
	fmt.Println("JSON data written to ", jsonFile.Name())
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

//String converts action to string
func (a Action) String() string {
	actions := [...]string{
		"Nothing",
		"Coffee",
		"Mate",
		"Water",
	}

	if a > Action(3) || a < Action(0) {
		return "Unknown"
	}

	return actions[a]
}

//GetAction converts string to action
func GetAction(a string) Action {
	actions := map[string]Action{
		"nothing": Nothing,
		"coffee":  Coffee,
		"mate":    Mate,
		"water":   Water,
	}
	return actions[a]
}
