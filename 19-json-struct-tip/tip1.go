package main

import (
	"encoding/json"
	"fmt"
)

//  Tip1: ignore field 2 ways:

type Meta struct {
	Username string `json:"username" :"Username"`
	Password string `json:"password" :"Password"`
	// many more fields…
}

type MetaProxy struct {
	Meta
	// ignore by downCase
	score string `json:"score" :"Score"`
	// ignore by tag
	Password string `json:"-"`
}

func main1() {
	m := new(MetaProxy)

	m.Username = "river"
	m.score = "100"
	m.Password = "123456"

	r, _ := json.Marshal(m)
	fmt.Println(string(r))
	// {"username":"river",Password:""}
}
