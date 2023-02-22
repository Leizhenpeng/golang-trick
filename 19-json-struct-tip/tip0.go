package main

import (
	"encoding/json"
	"fmt"
)

//  Tip1: ignore field 2 ways:

type Meta struct {
	Username string `json:"username" :"Username"`
	// ignore by downCase
	password string `json:"Password" :"Password"`
	// ignore by tag
	Score string `json:"-" :"Score"`
}

func main1() {
	m := Meta{
		Username: "river",
		password: "22",
		Score:    "100",
	}
	r, _ := json.Marshal(m)
	fmt.Println(string(r))
	// {"username":"river"}
}
