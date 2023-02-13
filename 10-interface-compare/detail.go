package main

import "reflect"

func main_() {
	var b = true
	var i interface{} = true

	if b == i {
		println("Comparable are equal")
	}

	if reflect.TypeOf(b) == reflect.TypeOf(i) {
		println("Dynamic types are equal")
	}

	if reflect.ValueOf(b) == reflect.ValueOf(i) {
		println("Dynamic values are equal")
	}
}
