package main

import (
	"fmt"
	"goog/22-timewheel/txt"
	"runtime"
	"time"

	"github.com/rfyiamcool/go-timewheel"
)

func rawTest(count int, loop int) {
	var m1, m2 runtime.MemStats
	runtime.ReadMemStats(&m1)
	var totalTime time.Duration
	var totalMemory uint64

	for index := 0; index < loop; index++ {
		start := time.Now()
		for index := 0; index < count; index++ {
			txt.NewSMS("Hello, world!")
		}

		loopTime := time.Since(start)
		runtime.ReadMemStats(&m2)
		totalMemory += m2.TotalAlloc - m1.TotalAlloc
		m1 = m2
		totalTime += loopTime
	}
	fmt.Println("-----timer result-----")
	fmt.Printf("loop: %d , time : %v,  memory : %dkb\n", 3, totalTime/3, totalMemory/3)
}

func twTest(count int, loop int) {
	tw, err :=
		timewheel.NewTimeWheel(1*time.Second, 20)
	if err != nil {
		panic(err)
	}
	tw.Start()
	defer tw.Stop()
	var m1, m2 runtime.MemStats
	runtime.ReadMemStats(&m1)
	var totalTime time.Duration
	var totalMemory uint64

	for index := 0; index < loop; index++ {
		start := time.Now()
		for index := 0; index < count; index++ {
			txt.NewSMS_("Hello, world!", tw)
		}
		loopTime := time.Since(start)
		runtime.ReadMemStats(&m2)
		totalMemory += m2.TotalAlloc - m1.TotalAlloc
		m1 = m2
		totalTime += loopTime
	}
	fmt.Println("-----timewheel result-----")
	fmt.Printf("loop: %d , time : %v,  memory : %dkb\n", 3, totalTime/3, totalMemory/3)
}

func main() {
	rawTest(500000, 3)
	twTest(500000, 3)
}
