package main

type Book__ struct {
	cfg *Config
}
type Config struct {
	Title string
	ISBN  string
}

func NewBook__(cfg *Config) *Book__ {
	return &Book__{cfg}
}

func main__() {
	_ = NewBook__(&Config{
		Title: "早点下班的go语言学习法",
		ISBN:  "20230101",
	})
}
