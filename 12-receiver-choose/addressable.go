package main

import "fmt"

type animal interface {
	speak() string
}

type dog struct{}

func (d dog) speak() string {
	return "Woof!"
}

func main() {
	d := dog{}
	var a animal = d
	fmt.Println(a) // this will output "&{}"
	fmt.Println(&a)
	// this will result in a compile-time error, as the concrete value stored in 'a' is not addressable
}
