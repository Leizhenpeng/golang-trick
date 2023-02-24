package main

import "fmt"

func main() {
	m := map[int]string{
		1: "golang",
		2: "rust",
		3: "ts",
	}

	for k := range m {
		fmt.Println(k, m[k])
	}

}

// why map iteration is random?
// 1. 汇编 go build -gcflags="-l -S -N" ./tip1.go  > initMap.s 2>&1
// 2. 查看核心调用代码 mapiterinit
// 3. 查看 mapiterinit 源代码
