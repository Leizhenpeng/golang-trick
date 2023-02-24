package main

import (
	"encoding/json"
	"fmt"
)

// Tip0:
// Outer Struct Override Inner Struct

type Inner struct {
	Name  string `json:"name" `
	Price int    `json:"price"`
}

type Outer struct {
	Inner
	Price string `json:"price"`
}

func main0() {
	m := new(Outer)

	m.Name = "river"
	m.Price = "100"

	r, _ := json.Marshal(m)
	fmt.Println(string(r))
}
