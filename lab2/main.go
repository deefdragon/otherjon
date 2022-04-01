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
	return 0.0
}

func taxesProgressive(earnings float64) (taxes float64) {
	return 0.0
}

func accountGrowth(addedPerWeek int) (weeks int) {
	return 0
}
