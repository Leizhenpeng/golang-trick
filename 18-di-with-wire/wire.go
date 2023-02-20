//go:build wireinject
// +build wireinject

package main

import "github.com/google/wire"

func InitAPeople(msg string) People {
	wire.Build(
		NewPeople,
		NewMouth,
		NewHand,
		//..
		NewMessage,
		NewMeta)
	return People{}
}
