package main

type SellInfo struct {
	Price      float64
	OrderCount int
	TotalCount int
	MemberShip int
}

func main2() {
	var a = SellInfo{
		Price:      1.0,
		OrderCount: 1,
		TotalCount: 20,
		MemberShip: 1,
	}
	if a.Price > 0 {
		println("true")
	}
	if a.TotalCount > a.OrderCount {
		println("true")
	} else {
		println("false")
	}
	if a.MemberShip == 1 {
		println("true")
	}
	if a.Price < 100 && a.MemberShip == 2 {
		println("true")
	}
	// if ...else...
	// if ...else if ...else if ...else if ...else...
}

// How To Kill Disgusting If Statement?
