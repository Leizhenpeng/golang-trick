package main

type someInterface interface {
	DoSomething()
	DoAnotherThing()
}

type someStruct struct {
}

// 如何快速实现接口 - how to quick impl someInterface for someStruct

// 1. imple someInterface for someStruct{}

// 2.
var _ someInterface = (*someStruct)(nil)

// DoAnotherThing implements someInterface
func (*someStruct) DoAnotherThing() {
	panic("unimplemented")
}

// DoSomething implements someInterface
func (*someStruct) DoSomething() {
	panic("unimplemented")
}
