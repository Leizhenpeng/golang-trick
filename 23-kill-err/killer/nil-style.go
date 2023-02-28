package killer

import "fmt"

type Book struct {
	Name   string
	Price  int
	Store  int
	Member int
}

func clientExample() {
	book := &Book{
		Name:   "golang",
		Price:  100,
		Store:  100,
		Member: 1,
	}
	err := book.CalcDiscount(99)
	if err != nil {
		return
	}
	err = book.JuedeIfStore()
	if err != nil {
		return
	}
	err = book.IfSale()
	if err != nil {
		return
	}
	err = book.CalcTotal()
	if err != nil {
		return
	}
	fmt.Println("ok")
}
func (b *Book) CalcDiscount(count int) error {
	// ...
	b.Price = b.Price * count / 100
	err := SomeError()
	if err != nil {
		return err
	}
	return nil
}
func (b *Book) JuedeIfStore() error {

	return nil
}
func (b *Book) IfSale() error {
	// ...
	return nil
}
func (b *Book) CalcTotal() error {
	// ...
	return nil
}
func (b *Book) AddPay() error {
	// ...
	return nil
}
func (b *Book) AddOrder() error {
	// ...
	return nil
}

func SomeError() error {
	return fmt.Errorf("some error")
}
