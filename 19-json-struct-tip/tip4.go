package main

import (
	"encoding/json"
	"fmt"
)

//  Tip3: type convert

type Base struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type BaseProxy struct {
	Base
	Code int `json:"code,string"`
}

func main3() {
	r := new(BaseProxy)
	r.Code = 200
	r.Msg = "ok"

	r2, _ := json.Marshal(r)
	fmt.Println(string(r2))
	// {"code":200,"msg":"ok","error":""}
	// {"code":200,"msg":"ok"}

}
