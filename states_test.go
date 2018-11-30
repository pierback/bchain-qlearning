package main

import (
	"fmt"
	"testing"
)

func TestDrinksCount(t *testing.T) {
	t.Parallel()
	fmt.Println("TestDrinksCount start")
	InitStateSpace()

	fmt.Println("	")
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
		if ts.GetCurrentTimeSlot(slot.input); ts != slot.expected {
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
