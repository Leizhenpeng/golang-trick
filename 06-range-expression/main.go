package main

import "fmt"

func main() {
	s := []string{"something"}
	for range s {
		s = append(s, "new_one")
		fmt.Println(s)
	}
}
