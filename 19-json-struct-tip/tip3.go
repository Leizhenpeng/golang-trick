package main

import (
	"encoding/json"
	"fmt"
)

//  Tip3: type convert

type Base struct {
	Code int    `json:"code,string"`
	Msg  string `json:"msg"`
}

func main3() {
	r := new(Base)
	r.Code = 200
	r.Msg = "ok"

	r2, _ := json.Marshal(r)
	fmt.Println(string(r2))
	// {"code":"200","msg":"ok"}

}
