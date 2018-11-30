package main

import (
	"encoding/json"
	"fmt"
)

func (state *State) String() string {
	out, err := json.Marshal(*state)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return string(out)
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
