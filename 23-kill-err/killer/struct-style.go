package killer

import "fmt"

func someError() error {
	return fmt.Errorf("some error")
}

type Book2 struct {
	Name   string
	Price  int
	Store  int
	Member int
	error  error
}

func clientExample2() {
	book2 := &Book2{
		Name:   "golang",
		Price:  100,
		Store:  100,
		Member: 1,
	}
	book2.CalcDiscount(99).
		JuedeIfStore().
		IfSale().
		AddPay().
		AddOrder().
		CheckAll()

	fmt.Println("ok")

}

func (b *Book2) CalcDiscount(count int) *Book2 {
	if b.error != nil {
		return b
	}
	// 具体的业务逻辑..
	b.error = someError()
	return b
}

func (b *Book2) JuedeIfStore() *Book2 {
	if b.error != nil {
		return b
	}
	// 具体的业务逻辑
	return b
}
func (b *Book2) IfSale() *Book2 {
	if b.error != nil {
		return b
	}
	return b
}
func (b *Book2) CalcTotal() *Book2 {
	if b.error != nil {
		return b
	}
	return b
}
func (b *Book2) AddPay() *Book2 {
	if b.error != nil {
		return b
	}
	return b
}
func (b *Book2) AddOrder() *Book2 {
	if b.error != nil {
		return b
	}
	return b
}

func (b *Book2) CheckAll() *Book2 {
	if b.error != nil {
		return b
	}
	return b
}
