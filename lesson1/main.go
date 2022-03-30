package main

import (
	"fmt"
)

func main() {

	h := 10
	w := 3
	fmt.Printf("%d\n", int(float64(h)/float64(w)))

	fmt.Printf("PRIME DETECTOR: IN: %d, OUT: %t\n", 23, primeDector(23))
	// a := 3

	// // If statement example
	// if a > 10 {
	// 	fmt.Println("test ")
	// } else {
	// 	fmt.Println("false")
	// }

	// /*
	// 	block comment.
	// 	This is a for loop example
	// */
	// for i := 0; i <= 10; i++ {
	// 	fmt.Printf("the number %d\n", i)
	// }

	/*
		Problem: Print of the first 100 numbers, the primes
		Prime: A number that is evenly divisable by exactly 2 positive integers.

		for each number less than 100
			set p to true if number is prime
			if p is true, print number

	*/

	for i := 0; i <= 100; i++ {
		isPrime := primeDector(i)
		if isPrime {
			fmt.Printf("%d\n", i)
		}
	}
}

func primeDector(potentialPrime int) bool {
	//Prime: A number that is evenly divisable by exactly 2 positive integers.

	divisibleByCount := 0
	for i := 1; i <= potentialPrime; i++ {
		isEvenlyDivisible := (potentialPrime % i) == 0
		if isEvenlyDivisible {
			divisibleByCount++
		}
	}

	return divisibleByCount == 2

}
