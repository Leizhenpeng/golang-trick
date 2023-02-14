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
	fmt.Println(counter()) // Output: 1
	fmt.Println(counter()) // Output: 2
	fmt.Println(counter()) // Output: 3
}

// go build -gcflags='-m'
