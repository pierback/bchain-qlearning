package learning

import (
	"fmt"
	"testing"
	"time"
)

/* func TestDrinksCount(t *testing.T) {
	t.Parallel()
	fmt.Println("TestDrinksCount start")

	// fmt.Println(Q.statemap)
	fmt.Println("	")
}
*/
/* func TestGenerateTrainingSet(t *testing.T) {
	GenerateTrainingSet()
} */

func TestUpdate(t *testing.T) {
	t.Parallel()
	fmt.Println("TestUpdate start")
	tq := QLearning{}
	vs := VirtualState{}
	sampleState := vs.New(Drinkcount{CoffeeCount: 0, WaterCount: 0, MateCount: 0}, -1, -1)

	var execActns = []struct {
		input    Action
		expected State
	}{
		{Coffee, vs.New(Drinkcount{CoffeeCount: 1, WaterCount: 0, MateCount: 0}, -1, -1)},
		{Nothing, vs.New(Drinkcount{CoffeeCount: 0, WaterCount: 0, MateCount: 0}, -1, -1)},
		{Mate, vs.New(Drinkcount{CoffeeCount: 0, WaterCount: 0, MateCount: 1}, -1, -1)},
		{Water, vs.New(Drinkcount{CoffeeCount: 0, WaterCount: 1, MateCount: 0}, -1, -1)},
		{4, vs.New(Drinkcount{CoffeeCount: 0, WaterCount: 0, MateCount: 0}, -1, -1)},
	}

	for _, ea := range execActns {
		tq.State = sampleState
		output := tq.State.Update(ea.input)

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
				Timeslot:   GetCurrentTimeSlot(time.Now().Hour()),
				Drinkcount: Drinkcount{CoffeeCount: 1, WaterCount: 2, MateCount: 3},
			},
			expected: UserState{
				Weekday:    time.Now().Weekday(),
				Timeslot:   GetCurrentTimeSlot(time.Now().Hour()),
				Drinkcount: Drinkcount{CoffeeCount: 1, WaterCount: 2, MateCount: 3},
			},
		},
		{
			input: UserState{
				Weekday:    time.Now().Weekday() - 1,
				Timeslot:   GetCurrentTimeSlot(time.Now().Hour()),
				Drinkcount: Drinkcount{CoffeeCount: 4, WaterCount: 4, MateCount: 4},
			},
			expected: UserState{
				Weekday:    time.Now().Weekday(),
				Timeslot:   GetCurrentTimeSlot(time.Now().Hour()),
				Drinkcount: Drinkcount{CoffeeCount: 0, WaterCount: 0, MateCount: 0},
			},
		},
	}

	for _, us := range ussts {
		if output := us.input.Get(); output != us.expected {
			t.Error("Get Userstate not working properly", us.input, us.expected, output)
		}
	}

}

func TestIsNewDay(t *testing.T) {
	t.Parallel()
	fmt.Println("TestIsNewDay start")

	q := QLearning{}
	vs := VirtualState{}
	us := UserState{}
	q.State = us.New(Drinkcount{CoffeeCount: 0, WaterCount: 0, MateCount: 0}, time.Now().AddDate(0, 0, 2).Day(), 8.45)

	if q.isNewDay() == false {
		t.Error("isNewDay not working properly \n", q.State, time.Now().Day())
	} else {
		fmt.Printf("TestIsNewDay Current day %s differs from State day %v \n", time.Now().Weekday(), q.State.(UserState).Weekday)
	}

	q.State = us.New(Drinkcount{CoffeeCount: 0, WaterCount: 0, MateCount: 0}, time.Now().Day(), 8.45)

	if q.isNewDay() == true {
		t.Error("isNewDay not working properly \n", q.State, time.Now().Day())
	} else {
		fmt.Printf("TestIsNewDay Current day  %v and State day %v are the same \n", time.Now().Weekday(), q.State.(UserState).Weekday)
	}

	q.State = vs.New(Drinkcount{CoffeeCount: 0, WaterCount: 0, MateCount: 0}, time.Now().Day(), 8.45)

	if q.isNewDay() == true {
		t.Error("TestIsNewDay isNewDay not working properly \n", q.State, time.Now().Day())
	}

	/* q.State = us.New(Drinkcount{CoffeeCount: 0, WaterCount: 0, MateCount: 0}, time.Now().Day(), 8.45)

	if q.isNewDay() == true {
		t.Error("isNewDay not working properly", q.State, time.Now().Day())
	} */

}
