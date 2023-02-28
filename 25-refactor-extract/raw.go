package main

import "math"

type Order struct {
	Quantity  float64
	ItemPrice float64
}

func price(order Order) float64 {
	// price formula:
	// base price - quantity discount + shipping
	return order.Quantity*order.ItemPrice -
		math.Max(0, order.Quantity-500)*order.ItemPrice*0.05 +
		math.Min(order.Quantity*order.ItemPrice*0.1, 100)
}

func main() {
	o := Order{
		Quantity:  1000,
		ItemPrice: 10,
	}
	println(price(o))
}
