package main

import "fmt"

func main() {
	//absolute taxes tests.
	in := 10000.0
	fmt.Printf("Absolute taxes for %f is %f\n", in, taxesAbsolute(in))
	in = 100000.0
	fmt.Printf("Absolute taxes for %f is %f\n", in, taxesAbsolute(in))
	in = 1000000.0
	fmt.Printf("Absolute taxes for %f is %f\n", in, taxesAbsolute(in))

	//progressive taxes tests.
	in = 10000.0
	fmt.Printf("Progressive taxes for %f is %f\n", in, taxesProgressive(in))
	in = 100000.0
	fmt.Printf("Progressive taxes for %f is %f\n", in, taxesProgressive(in))
	in = 1000000.0
	fmt.Printf("Progressive taxes for %f is %f\n", in, taxesProgressive(in))

	//progressive taxes tests.
	N := 1000
	fmt.Printf("%d dollars per week passed in %d weeks\n", N, accountGrowth(N))
	N = 5000
	fmt.Printf("%d dollars per week passed in %d weeks\n", N, accountGrowth(N))
	N = 25165
	fmt.Printf("%d dollars per week passed in %d weeks\n", N, accountGrowth(N))
	N = 65536
	fmt.Printf("%d dollars per week passed in %d weeks\n", N, accountGrowth(N))
	// N = random value
	// fmt.Printf("%d dollars per week passed in %d weeks\n", N, accountGrowth(N))

}

func taxesAbsolute(earnings float64) (taxes float64) {
	// Step by step checks earnings ranges and declares the tax value based on the answer
	if earnings >= 0 && earnings <= 9275 {
		taxes = .10
	} else {
		// utilizes stacking If statements to go through step by step
		if earnings >= 9275 && earnings <= 37650 {
			taxes = .15
		} else {
			if earnings >= 37650 && earnings <= 91150 {
				taxes = .25

			} else if earnings >= 91150 && earnings <= 190150 {
				taxes = .28
			} else if earnings >= 190150 && earnings <= 413350 {
				taxes = .33
			} else if earnings >= 413350 && earnings <= 415050 {
				taxes = .35
			} else if earnings >= 415050 {
				taxes = .396
			}
		}
	}
	earnings *= taxes
	return earnings
}

func taxesProgressive(earnings float64) (taxes float64) {
	return 0.0
}

func accountGrowth(addedPerWeek int) (weeks int) {
	return 0
}
