package main

type Book struct {
	Title string
	Code  string
}
type Option func(*Book)

func NewBook_(options ...Option) *Book {
	b := &Book{}
	for _, option := range options {
		option(b)
	}
	return b
}

func WithTitle(title string) Option {
	return func(b *Book) {
		b.Title = title
	}
}

func WithISBN(ISBN string) Option {
	return func(b *Book) {
		b.Code = ISBN
	}
}

func main_() {
	_ = NewBook_(
		WithTitle("Go for Dummies"), WithISBN("1234567890"),
	)
}
