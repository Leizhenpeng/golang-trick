package main

import (
	"encoding/json"
	"fmt"
)

//  Tip4: Store Raw Data Without Type

type Raw struct {
	Code int             `json:"code"`
	Resp json.RawMessage `json:"resp"`
}

func main() {
	rawStr := `{
		"code":200,
		"resp":{"name":"river", "age":18}}`
	var raw Raw
	json.Unmarshal([]byte(rawStr), &raw)
	fmt.Println(string(raw.Resp))
	//{"name":"river", "age":18}
}
