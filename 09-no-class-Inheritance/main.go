package main

import "fmt"

type People struct {
	name string
}

func (p *People) getName() string {
	return p.name
}

type Namer interface {
	getName() string
}

func sayHello(p Namer) {
	fmt.Printf("Hello %s", p.getName())
}

type Boy struct {
	People
}

func main() {
	b := &Boy{
		People: People{name: "John"},
	}
	fmt.Printf("name is %s \n", b.name)
	fmt.Printf("hi to %s \n", b.getName())

	sayHello(b)
}
