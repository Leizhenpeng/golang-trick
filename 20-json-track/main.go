package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Score struct {
	Name string
	time.Time
}

func main() {
	s := Score{
		Name: "river",
		Time: time.Now(),
	}
	r, err := json.Marshal(s)
	if err != nil {
		return
	}
	fmt.Printf("json: %s\n", string(r))

	/*
		guess the output

		Is { name: "river",
			time: "2023-02-21T23:04:34.042027+08:00"}
		??
	*/

}
