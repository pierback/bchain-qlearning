package learning

import (
	"fmt"
	"testing"
	"time"
)

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
		{9, 0},
		{8, 0},
		{10, 1},
		{12, 1},
		{14, 2},
		{16, 3},
		{18, 3},
		{22, 4},
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
		input    time.Weekday
		expected string
	}{
		{1, "Monday"},
		{2, "Tuesday"},
		{3, "Wednesday"},
		{4, "Thursday"},
		{5, "Friday"},
	}

	for _, weekday := range weekdays {
		if output := weekday.input.String(); output != weekday.expected {
			t.Error("weekday Failed: {} inputted, {} expected, recieved: {}", weekday.input, weekday.expected, output)
		}
	}
}
