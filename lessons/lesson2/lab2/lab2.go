package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	// test()
	// return
	//absolute taxes tests.
	in := 9276.0
	fmt.Printf("Absolute taxes for %f is %f\n", in, taxesAbsolute(in))
	in = 100000.0
	fmt.Printf("Absolute taxes for %f is %f\n", in, taxesAbsolute(in))
	in = 1000000.0
	fmt.Printf("Absolute taxes for %f is %f\n", in, taxesAbsolute(in))

	//progressive taxes tests.
	in = 9276.0
	fmt.Printf("Progressive taxes for %f is %f\n", in, taxesProgressive(in))
	in = 100000.0
	fmt.Printf("Progressive taxes for %f is %f\n", in, taxesProgressive(in))
	in = 1000000.0
	fmt.Printf("Progressive taxes for %f is %f\n", in, taxesProgressive(in))

	//progressive taxes tests.
	N := 1000
	fmt.Printf("%d dollars per week passed in %d weeks\n", N, accountGrowth(N))
	N = rand.Intn(4001) + 1000
	fmt.Printf("%d dollars per week passed in %d weeks\n", N, accountGrowth(N))
	N = rand.Intn(4001) + 1000
	fmt.Printf("%d dollars per week passed in %d weeks\n", N, accountGrowth(N))
	N = rand.Intn(4001) + 1000
	fmt.Printf("%d dollars per week passed in %d weeks\n", N, accountGrowth(N))
	// N = random value
	// fmt.Printf("%d dollars per week passed in %d weeks\n", N, accountGrowth(N))

}

func taxesAbsolute(earnings float64) (taxes float64) {
	// Step by step checks earnings ranges and declares the tax value based on the answer
	if earnings >= 0 && earnings <= 9275 {
		taxes = .10
	} else if earnings >= 9275 && earnings <= 37650 {
		taxes = .15
	} else if earnings >= 37650 && earnings <= 91150 {
		taxes = .25
	} else if earnings >= 91150 && earnings <= 190150 {
		taxes = .28
	} else if earnings >= 190150 && earnings <= 413350 {
		taxes = .33
	} else if earnings >= 413350 && earnings <= 415050 {
		taxes = .35
	} else {
		taxes = .396
	}
	earnings *= taxes
	return earnings
}

func taxesProgressive(earnings float64) (taxes float64) {
	taxes = 0

	// var tr taxRange
	/*
		low:  9275,
		high: 37650,
		rate: .15,
	*/
	for _, tr := range taxRanges {
		if earnings >= tr.low && earnings <= tr.high {
			portion := (earnings - tr.low) * tr.rate
			taxes += portion
			break
		} else if earnings > tr.high {
			F := (tr.high - tr.low) * tr.rate
			taxes += F
			// do the progressive calculation of that range.
		}
	}

	return taxes
}

type taxRange struct {
	low  float64
	high float64
	rate float64
}

var taxRanges = []taxRange{
	{
		low:  0,
		high: 9275,
		rate: .10,
	},
	{
		low:  9275,
		high: 37650,
		rate: .15,
	},
	{
		low:  37650,
		high: 91150,
		rate: .25,
	},
	{
		low:  91150,
		high: 190150,
		rate: .28,
	},
	{
		low:  190150,
		high: 413350,
		rate: .33,
	},
	{
		low:  413350,
		high: 415050,
		rate: .35,
	},
	{
		low:  415050,
		high: math.MaxFloat64,
		rate: .396,
	},
}

// Calculate each bracket's deductions up to point of input
// input earnings
// Find amount taken on every tax bracket in between start and earnings
// Add all amounts together to find total
// output taxes

func accountGrowth(addedPerWeek int) (weeks int) {
	accA := 0.01
	accB := 0.0
	weeks = 0

	for {
		accA *= 2
		accB += float64(addedPerWeek)
		weeks++

		if accA > accB {
			break
		}
	}

	return weeks
}

// func test() {
// 	r := rand.New(rand.NewSource(99))
// 	r2 := rand.New(rand.NewSource(99))
// 	fmt.Println(r2.Int(1, 10))
// 	fmt.Println(r.Int())
// 	fmt.Println(r.Int())
// 	fmt.Println(r2.Int())
// }
