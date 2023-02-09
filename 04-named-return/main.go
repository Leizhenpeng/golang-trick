package main

import "fmt"

func main() {
	a := []string{"go", "rust", "cpp"}
	b := CopyIt(a)
	fmt.Println(b)
}

func CopyIt(raw []string) []string {
	r := make([]string, 0)
	for i, v := range raw {
		r[i] = v
	}
	return r
}

/*
how to optimize function CopyIt
target: one line  code
*/
