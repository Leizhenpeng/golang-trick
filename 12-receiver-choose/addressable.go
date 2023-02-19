package main

import "fmt"

type I interface{}
type AInt int
type AString string

func main() {
	var a AInt = 5
	var i I = a
	fmt.Printf("i is of type %T\n", i)
	var aPtr *AInt
	// aPtr = &(i.(AInt))
	var b AString = "hello"
	i = b
	fmt.Printf("i is of type %T, aPtr is of type %T\n", i, aPtr)
}
