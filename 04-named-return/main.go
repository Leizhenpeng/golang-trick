package main

import "fmt"

func main() {
	a := []string{"go", "rust", "cpp"}
	b := CopyIt(a)
	fmt.Println(b)
}

func CopyIt(raw []string) (r []string) {
	// r = append(r, raw...)
	copy(r, raw)
	return r
}

/*
how to optimize function CopyIt
target: one line  code
*/
