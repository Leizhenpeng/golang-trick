package main

import "fmt"

type I interface{}
type AInt int
type AString string

func explain1() {
	m := map[string]int{
		"go": 2023,
	}
	var i I = m["go"]
	fmt.Printf("i is of type %T\n", i)
}

func explain2() {
	var a AInt = 5
	var aPtr *AInt

	var i I = a
	fmt.Printf("i is of type %T\n", i)
	// aPtr = &(i.(AInt))
	var x AInt = i.(AInt)
	aPtr = &x

	var b AString = "hello"
	i = b
	fmt.Printf("i is of type %T, aPtr is of type %T\n", i, aPtr)
}

// can't take the address of value in interface
// 1. some value is not addressable
// 2. compiler will confilct with type
