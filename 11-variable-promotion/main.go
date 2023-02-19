package main

import "fmt"

func getCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {
	counter := getCounter()
	fmt.Println(counter()) // 1
	fmt.Println(counter()) // 2
	fmt.Println(counter()) // 3
}

// 为何闭包会导致记忆效应？
// 逃逸分析试试看
// go build -gcflags='-m'
