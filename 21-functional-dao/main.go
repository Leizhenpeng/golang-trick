package main

import "functional-dao/model"

func main() {
	model.InitDB()
	model.AddFakeBook()
	model.ClientExample2()
}
