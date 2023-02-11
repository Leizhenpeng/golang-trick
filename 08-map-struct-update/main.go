package main

import "fmt"

type Book struct {
	Title  string
	Prince int
}

func main() {
	m := make(map[string]*Book)
	m["rust"] =
		&Book{Title: "rust", Prince: 20}

	//how to update prince?
	// m["rust"].Prince = 30
	b, ok := m["rust"]
	if ok {
		b.Prince = 30
		m["rust"] = b

	}
	// _ = &m["rust"]

	fmt.Printf(
		"prince: %d", m["rust"].Prince)
}
