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
