package main

import (
	"fmt"
	"sync"
)

func (o Order) BasePrice() float64 {
	return o.Quantity * o.ItemPrice
}

func QuantityDiscount(o Order) float64 {
	discount := 0.0
	if o.Quantity > 500 {
		discount = 0.05
	}
	return o.BasePrice() * discount
}

func Shipping(o Order) float64 {
	shipping := o.BasePrice() * 0.1
	if shipping > 100 {
		shipping = 100
	}
	return shipping
}

func Price(o Order) float64 {
	return o.BasePrice() -
		QuantityDiscount(o) +
		Shipping(o)
}

func PriceQuick(o Order) float64 {
	var wg sync.WaitGroup
	var quantityDiscount float64
	var shipping float64

	basePrice := o.BasePrice()
	go func() {
		defer wg.Done()
		quantityDiscount = QuantityDiscount(o)
	}()
	go func() {
		defer wg.Done()
		shipping = Shipping(o)
	}()
	wg.Wait()

	return basePrice - quantityDiscount + shipping
}

func main_() {
	order := Order{Quantity: 600, ItemPrice: 10.0}
	fmt.Println(Price(order))
}
