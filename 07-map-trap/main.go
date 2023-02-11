package main

import "fmt"

func main() {
	m := map[string]int{
		"rust": 1,
		"go":   2,
		"cpp":  3,
	}
	//copy one
	m_copy := make(map[string]int, len(m))
	for k, v := range m {
		m_copy[k] = v
	}

	for k, v := range m_copy {
		m[k+"_new"] = v + 100
	}

	fmt.Printf("map length: %d\n", len(m))
}

// Output
// A: 3
// B: 6
// C: I don't know
