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

func TestUpdate(t *testing.T) {
	t.Parallel()
	fmt.Println("TestUpdate start")
	tq := QLearning{}
	vs := VirtualState{}
	sampleState := vs.New(drinkcount{CoffeeCount: 0, WaterCount: 0, MateCount: 0}, -1, -1)

	var execActns = []struct {
		input    Action
		expected State
	}{
		{Coffee, vs.New(drinkcount{CoffeeCount: 1, WaterCount: 0, MateCount: 0}, -1, -1)},
		{Nothing, vs.New(drinkcount{CoffeeCount: 0, WaterCount: 0, MateCount: 0}, -1, -1)},
		{Mate, vs.New(drinkcount{CoffeeCount: 0, WaterCount: 0, MateCount: 1}, -1, -1)},
		{Water, vs.New(drinkcount{CoffeeCount: 0, WaterCount: 1, MateCount: 0}, -1, -1)},
		{4, vs.New(drinkcount{CoffeeCount: 0, WaterCount: 0, MateCount: 0}, -1, -1)},
	}

	for _, ea := range execActns {
		tq.state = sampleState
		output := tq.state.Update(ea.input)

		if output != ea.expected {
			t.Error("Wrong reward given", ea.input, ea.expected, ea)
		}
		fmt.Println("output: ", output)
	}
}

func TestGet(t *testing.T) {
	t.Parallel()
	fmt.Println("TestGet start")

	var ussts = []struct {
		input    UserState
		expected State
	}{
		{
			input: UserState{
				Weekday:    time.Now().Weekday(),
				Timeslot:   getCurrentTimeSlot(time.Now().Hour()),
				Drinkcount: drinkcount{CoffeeCount: 1, WaterCount: 2, MateCount: 3},
			},
			expected: UserState{
				Weekday:    time.Now().Weekday(),
				Timeslot:   getCurrentTimeSlot(time.Now().Hour()),
				Drinkcount: drinkcount{CoffeeCount: 1, WaterCount: 2, MateCount: 3},
			},
		},
		{
			input: UserState{
				Weekday:    time.Now().Weekday() - 1,
				Timeslot:   getCurrentTimeSlot(time.Now().Hour()),
				Drinkcount: drinkcount{CoffeeCount: 4, WaterCount: 4, MateCount: 4},
			},
			expected: UserState{
				Weekday:    time.Now().Weekday(),
				Timeslot:   getCurrentTimeSlot(time.Now().Hour()),
				Drinkcount: drinkcount{CoffeeCount: 0, WaterCount: 0, MateCount: 0},
			},
		},
	}

	for _, us := range ussts {
		if output := us.input.Get(); output != us.expected {
			t.Error("Get Userstate not working properly", us.input, us.expected, output)
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
				Weekday:  time.Now().Weekday(),
				Timeslot: getCurrentTimeSlot(time.Now().Hour()),
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
				Weekday:  3,
				Timeslot: getCurrentTimeSlot(int(3)),
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
