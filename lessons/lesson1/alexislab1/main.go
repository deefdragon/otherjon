package main

import (
	"fmt"
)

func main() {

	// h := 10
	// w := 3

	// fmt.Printf("%d\n", int(float64(h)/float64(w)))

	// // fmt.Printf("%d\n", int(float64(h)/float64(w)))

	fmt.Println("abcd"[0:3])

	// // printDiv10(500, 600)

	fmt.Println("Negatives")
	fmt.Println(10)
	fmt.Println(Negatives(10))

	fmt.Println(35)
	fmt.Println(Negatives(35))

	fmt.Println(-221)
	fmt.Println(Negatives(-221))

	fmt.Println(48)
	fmt.Println(Negatives(48))

	fmt.Println("Pennies")
	fmt.Println(3.12)
	fmt.Println(Pennies(3.12))

	fmt.Println(1.00)
	fmt.Println(Pennies(1.00))

	fmt.Println(1.23)
	fmt.Println(Pennies(1.23))

	fmt.Println("CodeGen")
	fmt.Println("x", "5")
	fmt.Println(CodeGen("x", "5")) //should print "x := 5"

	fmt.Println("abs", "500")
	fmt.Println(CodeGen("abs", "500"))

	fmt.Println("foo", "test")
	fmt.Println(CodeGen("foo", "test"))

	fmt.Println("Fractions")
	fmt.Println(97, 52)
	fmt.Println(Fractions(97, 52))

	fmt.Println(8, 6)
	fmt.Println(Fractions(8, 6)) // 1.333

	fmt.Println(13, 7)
	fmt.Println(Fractions(13, 7))

	fmt.Println("SecondHalf")
	fmt.Println("this")
	fmt.Println(SecondHalf("this")) // is

	fmt.Println("potato")
	fmt.Println(SecondHalf("potato")) // ato

	fmt.Println("asdfghjl")
	fmt.Println(SecondHalf("asdfghjl")) //ghjl

	//len()
	//[startInclusive:endExclusive]
	//if missing, assume 0/end

}

func Negatives(input int) (output int) {
	//Do some stuff
	number := input * -1
	return number
}

func SecondHalf(input string) (output string) {

	return ""
}

func Pennies(input float64) (output int) {
	return 0
}

func CodeGen(input string, input2 string) (output string) {
	return ""
}

func Fractions(numerator int, denomernator int) (output float64) {
	number := numerator / denomernator
	return float64(number)
}

// func foo(text string) {

// 	if true {
// 		fmt.Println("I did the true thing")
// 	} else {
// 		fmt.Println("I did the not true thing")
// 	}

// 	if false {
// 		fmt.Println("I did the false thing")
// 	} else {
// 		fmt.Println("I did the not false thing")
// 	}

// 	for i := 0; i < 10; i++ {
// 		fmt.Println(i)
// 	}

// 	fmt.Println(text)

// }

// func printDiv10(minimum int, maximum int) {

// 	// someNumber := minimum

// 	for i := minimum; i <= maximum; i++ {
// 		isDivisibleBy10 := isDiv10(i)
// 		if isDivisibleBy10 {
// 			fmt.Println(i)
// 		}
// 	}

// 	/*
// 	* need to know starting and stopping points
// 	* how do we determine what numbers is divisible by 10
// 	* how do we go through each number.
// 	 */
// }

// func isDiv10(number int) (is bool) {
// 	var remainder int
// 	remainder = number % 10

// 	//when remainder is 7, return false.
// 	//when remainder is 0, return true.

// 	return remainder == 0
// }
