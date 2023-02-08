package main

type someStruct struct {
}

func NewPointer() interface{} {
	var p *someStruct = nil
	return p
}

func main() {
	p := NewPointer()

	if p == nil {
		println("p is nil")
	} else {
		println("p is not nil")
	}
}

/*
# which output?
# A. p is nil
# B. p is not nil
*/
