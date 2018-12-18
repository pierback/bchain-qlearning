package learning

import (
	"fmt"
	"testing"
	"time"
)

func TestDrinksCount(t *testing.T) {
	t.Parallel()
	fmt.Println("TestDrinksCount start")

	// fmt.Println(Q.statemap)
	fmt.Println("	")
}

func TestGenerateTrainingSet(t *testing.T) {
	GenerateTrainingSet()
}
func TestSSID(t *testing.T) {
	t.Parallel()
	fmt.Println("TestSSID start")

	var ssids = []struct {
		input    SSID
		expected bool
	}{
		{"eduroam", true},
		{"ipsum", false},
	}

	for _, ssid := range ssids {
		if output := ssid.input.isEduroam(); output != ssid.expected {
			t.Error("Wrong wifi", ssid.input, ssid.expected, ssid)
		}
	}

}

func TestNew(t *testing.T) {
	t.Parallel()
	fmt.Println("TestNew start")

	type Input struct {
		dc drinkcount
		wd int
		ct float64
	}

	var stfs = []struct {
		input    Input
		expected State
	}{
		{
			input: Input{dc: drinkcount{CoffeeCount: 4, WaterCount: 0, MateCount: 0}, wd: -1, ct: -1},
			expected: VirtualState{
				Weekday:  weekday(time.Now().Weekday()),
				Timeslot: GetCurrentTimeSlot(time.Now().Hour()),
				Drinkcount: drinkcount{
					CoffeeCount: 4,
					WaterCount:  0,
					MateCount:   0,
				},
			},
		},
		{
			input: Input{dc: drinkcount{CoffeeCount: 4, WaterCount: 0, MateCount: 0}, wd: 3, ct: 3},
			expected: VirtualState{
				Weekday:  weekday(3),
				Timeslot: GetCurrentTimeSlot(int(3)),
				Drinkcount: drinkcount{
					CoffeeCount: 4,
					WaterCount:  0,
					MateCount:   0,
				},
			},
		},
	}

	vs := VirtualState{}
	for _, s := range stfs {
		if output := vs.New(s.input.dc, s.input.wd, s.input.ct); output != s.expected {
			t.Error("New not working properly", s.input, s.expected, output)
		}
	}

}

func TestTimeSlot(t *testing.T) {
	t.Parallel()
	fmt.Println("TestTimeSlot start")

	var slotStrings = []struct {
		input    timeslot
		expected string
	}{
		{0, "Slot0"},
		{1, "Slot1"},
		{2, "Slot2"},
		{3, "Slot3"},
		{4, "Slot4"},
		{5, "Slot5"},
		{6, "Slot6"},
		{777, "Unknown"},
	}

	var slotIds = []struct {
		input    int
		expected timeslot
	}{
		{7, 0},
		{9, 1},
		{8, 0},
		{10, 1},
		{12, 2},
		{14, 3},
		{16, 4},
		{18, 5},
		{22, 6},
		{26, -1},
	}
	var ts timeslot

	for _, slot := range slotStrings {
		if output := slot.input.TimeSlotString(); output != slot.expected {
			t.Error("Slot Strings", slot.input, slot.expected, output)
		}
	}

	for _, slot := range slotIds {
		if out := GetCurrentTimeSlot(slot.input); out != slot.expected {
			t.Error("Input Hour", slot.input, slot.expected, ts)
		}
	}
}

func TestWeekdayString(t *testing.T) {
	t.Parallel()
	fmt.Println("TestWeekdayString start")

	var weekdays = []struct {
		input    weekday
		expected string
	}{
		{0, "Monday"},
		{1, "Tuesday"},
		{2, "Wednesday"},
		{3, "Thursday"},
		{4, "Friday"},
	}

	for _, weekday := range weekdays {
		if output := weekday.input.String(); output != weekday.expected {
			t.Error("weekday Failed: {} inputted, {} expected, recieved: {}", weekday.input, weekday.expected, output)
		}
	}
}
