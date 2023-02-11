package main

type Book struct {
	Title string
	ISBN  string
}

func NewBook(t string) *Book {
	return &Book{
		Title: t,
		ISBN:  "",
	}
}

func NewBook2(t string, i string) *Book {
	return &Book{
		Title: t,
		ISBN:  i,
	}
}

func main() {
	_ = NewBook("早点下班的go语言学习法")
	_ = NewBook2("早点下班的go语言学习法", "20230101")
}

// How To Make Constructor Better
