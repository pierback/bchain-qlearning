package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	// handleRequests()
	// initLearner()
	var ts TimeSlot = 9
	h := time.Now().Hour()
	ts.GetCurrentTimeSlot(h)
	fmt.Println("&slot", ts)

	fmt.Printf("Which day it is? %s\n", time.Now().Weekday())
	fmt.Printf("Slot? %s\n", reflect.TypeOf(time.Now().Hour()))
	// Which day it is? Sunday
	// insertNRetrieve()
	// wsInit()
}
