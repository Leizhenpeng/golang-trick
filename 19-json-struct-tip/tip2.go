package main

import (
	"encoding/json"
	"fmt"
)

//  Tip2: Omitting empty values

type Resp struct {
	Code     int    `json:"code"`
	Msg      string `json:"msg"`
	ErrorInf string `json:"error"`
}

type RespProxy struct {
	Resp
	ErrorInf string `json:"error,omitempty"`
}

func main2() {
	r := new(RespProxy)
	r.Code = 200
	r.Msg = "ok"
	r.ErrorInf = ""

	r2, _ := json.Marshal(r)
	fmt.Println(string(r2))
	// {"code":200,"msg":"ok"}

}
