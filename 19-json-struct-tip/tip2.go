package main

import (
	"encoding/json"
	"fmt"
)

//  Tip1: ignore field 2 ways:

type Meta struct {
	Username string `json:"username" :"Username"`
	// ignore by downCase
	password string `json:"password" :"Password"`
	// ignore by tag
	Score string `json:"-" :"Score"`
}

const jsonStr = `{
	"username": "gopher",
	"password": "123456",
	"score": "100"
}`

func main() {
	var meta Meta
	json.Unmarshal([]byte(jsonStr), &meta)
	fmt.Printf("%#v", meta)
	//{Username: "gopher", password: "", Score: ""}
}
