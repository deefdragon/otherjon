package main

import "fmt"

// Lab Assignment 1: Negatives - Given an integer, calculate and return it negative.

// Declared main function
func main() {
	// Printed the result using the fmt package, via a string, the variable, and another string
	fmt.Printf("Jon gets %d hours of sleep tonight!\n", Negatives(12))
	// Prints "Jon has (jonsSavings amount) pennies in his savings!" to the console, while specifying to printf what the various data types it needs to print is
	fmt.Printf("Jon has %d pennies in his savings!\n", Pennies(73.34))
	// Printed the message to the console using Printf and specifying the data types of the variables it was supposed to print
	fmt.Printf("Jon has had %d cups of coffee tonight!", codeGeneration(3))
	// Printf is used to print a statement to the console, a string, followed by a variable divided by another variable, followed by another string
	fmt.Printf("\nJon's attention span is working at %f efficiency!", fractions(5, 10))
	// Used Println to print text to console, "\n" to put it on a new line, then printed variable thanos, but sliced the value to only the back half
	fmt.Println("\n", secondHalf("Snap"))
}

func Negatives(jonsSleepHrs int) int {
	// Set variable to boolean that is true
	var hasBeenCoding bool = true
	// Created If statement tied to boolean variable hasBeenCoding
	if hasBeenCoding {
		// If true, multiplied jonsSleepHrs by -1
		jonsSleepHrs *= -1
	}
	return jonsSleepHrs
}

// Lab Assignment 2: Pennies - Given a floating point number of dollars and cents ($3.14), calculate and return the number of pennies

/* In this assignment, I've set up variables to start with, my savings and a boolean declaring that I took it out in the form of change, brokePggyBnk.
   I've then created an if statement using the boolean stating that if I did break my piggy bank open, multiply my savings by 100, to get the amount of pennies I'd have.
   After that, I've printed the result to the console with a string, a float, and another string using printf, and specifying the data types it should be looking for.
*/

func Pennies(jonsSavings float64) int {
	// Created If statement, if brokePggyBnk is true, jonsSavings multiplies by 100
	// Set variable brokePggyBnk to true
	var brokePiggyBank bool = true
	if brokePiggyBank {
		jonsSavings *= 100
	}
	return int(jonsSavings)
}

// Lab Assignment 3: Code generation - Given a string representing a variable and one representing its value, return the code string that would be used to initalize that variable to that value

func codeGeneration(cups int) int {
	addiction := cups * 2
	return addiction
}

//Given "x" and "5" return "x := 5"
//Given "y" and "10" return "y := 10"
/*
func codeGeneration(name, value string) string {
	output := fmt.Sprintf("%s := %s", name, value)
	output = name + " := " + value
	return output
}
*/

// Lab Assignment 4: Fractions - Given a numerator and a denominator, calculate and return the decimal representation of the fraction

func fractions(numer float64, denom float64) float64 {

	result := numer / denom
	return result
}

//Lab Assignment 5: Second Half - Given a string, return the second half of the string

/* You wouldn't believe how long it took me to remember semicolins exist. I've set a variable equal to a string, then printed the variable, specifying I only wanted
the last 2 characters in the 4 character array :) (and I experienced another off by one trying 3:4 for those wondering)
*/

func secondHalf(thanos string) string {
	number := len(thanos) / 2
	snap := thanos[number:]
	return snap
}
