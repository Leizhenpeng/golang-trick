package main

import "fmt"

func main() {
	s := []string{"something"}
	for i := 0; i < len(s); i++ {
		s = append(s, "new_one")

		fmt.Println(s)
	}
}

// Which Output is correct:
// A. [something new_one]

// B. [something new_one],
//    [something new_one new_one ]
//    [something new_one new_one new_one] ...

// C. panic: runtime error: index out of range
