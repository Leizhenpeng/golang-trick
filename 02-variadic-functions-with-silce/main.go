package main

func ShowName(book ...string) {
	for _, content := range book {
		println(content)
	}
}

func main() {
	ShowName("rust", "go", "cpp")
}
