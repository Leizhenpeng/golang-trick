package main

import "fmt"

type Meta struct{}
type Message struct {
	msg string
	r   Meta
}
type Mouth struct{ Message Message }
type Hand struct{ Message Message }
type People struct {
	Mouth Mouth
	Hand  Hand
	//..
}

func main_() {
	meta := NewMeta()
	message := NewMessage("A Hi To U", meta)
	aMouth := NewMouth(message)
	aHand := NewHand(message)
	people := NewPeople(aMouth, aHand)
	people.Hi()
}

//use wire
func main() {
	p := InitAPeople("Nice to meet you!")
	p.Hi()
}
func NewMeta() Meta {
	return Meta{}
}
func NewMessage(msg string, r Meta) Message {
	return Message{
		msg: msg,
		r:   r,
	}
}
func NewMouth(m Message) Mouth {
	return Mouth{Message: m}
}
func NewHand(m Message) Hand {
	return Hand{Message: m}
}
func NewPeople(g Mouth, t Hand) People {
	return People{Mouth: g, Hand: t}
}

func (p People) Hi() {
	msg := p.Mouth.Greet()
	fmt.Println(msg)
}
func (g Mouth) Greet() string {
	return g.Message.msg
}
