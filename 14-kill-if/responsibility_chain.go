package main

import "fmt"

type Rule interface {
	Check(sellInfo *SellInfo) bool
}

func Chain(sellInfo *SellInfo, rules ...Rule) bool {
	for _, r := range rules {
		if !r.Check(sellInfo) {
			return false
		}
	}
	return true
}

func main() {
	sellInfo := &SellInfo{
		Price:      1.0,
		OrderCount: 1,
		TotalCount: 20,
		MemberShip: 1,
	}

	rules := []Rule{
		&PriceRule{},
		&OrderCountRule{},
		&MemberShipRule{},
		&DiscountRule{},
		//...

	}
	r := Chain(sellInfo, rules...)
	fmt.Println(r)

}

type PriceRule struct{}

func (pr *PriceRule) Check(sellInfo *SellInfo) bool {
	return sellInfo.Price > 0
}

type OrderCountRule struct{}

func (ocr *OrderCountRule) Check(sellInfo *SellInfo) bool {
	return sellInfo.TotalCount > sellInfo.OrderCount
}

type MemberShipRule struct{}

func (msr *MemberShipRule) Check(sellInfo *SellInfo) bool {
	return sellInfo.MemberShip == 1
}

type DiscountRule struct{}

func (dr *DiscountRule) Check(sellInfo *SellInfo) bool {
	return sellInfo.Price < 100 && sellInfo.MemberShip == 2
}
