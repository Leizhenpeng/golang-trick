package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Result struct {
	Name string
	time.Time
}

func (r Result) MarshalJSON() ([]byte, error) {
	//{"Name":"river","Time":"2023-02-21T23:04:34.042027+08:00"}
	return json.Marshal(struct {
		Name string
		Time string
	}{
		Name: r.Name,
		Time: r.Time.Format(time.RFC3339),
	})
}

func main() {
	r := Result{
		Name: "river",
		Time: time.Now(),
	}
	b, err := json.Marshal(r)
	if err != nil {
		return
	}
	fmt.Printf("json: %s\n", string(b))
	//json: "2023-02-21T23:04:34.042027+08:00" why??
	//wahy not

}
