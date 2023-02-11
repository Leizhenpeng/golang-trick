package main

type Book___ struct {
	Title string
	Code  string
}
type Option func(*Book___)

func NewBook_(options ...Option) *Book___ {
	b := &Book___{}
	for _, option := range options {
		option(b)
	}
	return b
}

func WithTitle(title string) Option {
	return func(b *Book___) {
		b.Title = title
	}
}

func WithISBN(ISBN string) Option {
	return func(b *Book___) {
		b.Code = ISBN
	}
}

func main_() {
	_ = NewBook_(
		WithTitle("Go for Dummies"), WithISBN("1234567890"),
	)
}
