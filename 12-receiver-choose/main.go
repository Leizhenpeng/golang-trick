package main

import "fmt"

type Circle struct {
	Radius float64
}

type Shaper interface {
	Area() float64
	Scale(factor float64)
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (c Circle) Scale(factor float64) {
	c.Radius = c.Radius * factor
}

func main_() {
	// Common Sense:
	// 1.Value methods can be invoked on pointers and values,
	//2. pointer methods can only be invoked on pointers.
	//3. When the value is addressable, you can winvoking a pointer method on a value by inserting the address operator automatically.
	var c Shaper = Circle{10}
	ptr := &c
	fmt.Println(ptr)
	// addressable
	// var c Circle = Circle{10}
	// c.Scale(2)
	// c.Area()
	// (&c).Scale(2)
	// (&c).Area()
}
