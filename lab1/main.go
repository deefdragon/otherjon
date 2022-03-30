package main

import "fmt"

// Lab Assignment 1: Negatives - Given an integer, calculate and return it negative.

/*To start, I've made a basic function to calculate my level of sleep loss to accomodate for my self inflicted extra work load
I've declared a function, then set a variable to represent my usual amount of sleep every night
Then, I've factored in whether I've been keeping up on my practice. For now, this is true.
I've made an if statement to check to see if hasBeenCoding is true, in this case yes.
Because the boolean was true, we calculate sleep loss by reversing the amount of sleep represented by the variable "jonsSleepHrs"
Then we print this to the console and sha-bang! A fun little assignment to play with and reinforce some basic concepts.
*/
// Declared main function
func main() {
	// Printed the result using the fmt package, via a string, the variable, and another string
	fmt.Printf("Jon gets %d hours of sleep tonight!", Negatives(13))
	// Prints "Jon has (jonsSavings amount) pennies in his savings!" to the console, while specifying to printf what the various data types it needs to print is
	fmt.Printf("Jon has %f pennies in his savings!\n", jonsSavings)
	// Printed the message to the console using Printf and specifying the data types of the variables it was supposed to print
	fmt.Printf("%s %d %s", s, c, v)
	// Printf is used to print a statement to the console, a string, followed by a variable divided by another variable, followed by another string
	fmt.Printf("\nJon's attention span is working at %f %s", (numer / denom), "efficiency!")
	// Used Println to print text to console, "\n" to put it on a new line, then printed variable thanos, but sliced the value to only the back half
	fmt.Println("\n", thanos[2:4])
}

// // Set variable to integer 12
// var jonsSleepHrs int = 12

// Set variable to boolean that is true
var hasBeenCoding bool = true

func Negatives(jonsSleepHrs int) int {

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
// Set variable jonsSavings to 73.35
var jonsSavings = 73.35

// Set variable brokePggyBnk to true
var brokePggyBnk bool = true

func Pennies() {
	// Created If statement, if brokePggyBnk is true, jonsSavings multiplies by 100
	if brokePggyBnk {
		jonsSavings *= 100
	}
}

// Lab Assignment 3: Code generation - Given a string representing a variable and one representing its value, return the code string that would be used to initalize that variable to that value

/* This assignment seemed so trivial that I feel like I'm missing an instruction, I've created several variables, set them to strings and an integer,
and printed a message using the variables. (If theres something Im missing, just tack this on the next assignment and I'll make it up as extra credit)
*/
// Created variable s, set to a string representing the first half of the message
var s string = "Jon has had"

// Created variable c, set to an integer
var c int = 6

// Created variable v, set to a string representing the second half of the message
var v string = "cups of coffee tonight!"

// Just a quick note on this one, after moving variables out, I realized this function went unused... not sure how to implement it to be honest lmao but leaving it here for learning's sake
func codeGeneration() {

}

// Lab Assignment 4: Fractions - Given a numerator and a denominator, calculate and return the decimal representation of the fraction
/* Oof. Just oof. This ones gonna take some going over during the lesson if I did it wrong (which I'm sure I did) but 'ere we go.
   I defined two variables as float types, with the numerator set to 5 and the denominator set to 10.
   Then I did a Printf and threw in some strings to print a message, in the middle I divide numerator by denominator to get the fraction in decimal format.
*/
// Assigned variable "numer" to type float64 and gave it value 5
var numer float64 = 5

// Assigned variable "denom" to type float64 and gave it value 10
var denom float64 = 10

// Actually another quick note, same for fractions, all I did was create variables inside the function and print a line with those variables, this function is also unused
// not that I dont think they should be, I definately flubbed somewhere in reading/understanding instructions, or maybe theres a tool I'm failing to recall
func fractions() {

}

//Lab Assignment 5: Second Half - Given a string, return the second half of the string

/* You wouldn't believe how long it took me to remember semicolins exist. I've set a variable equal to a string, then printed the variable, specifying I only wanted
the last 2 characters in the 4 character array :) (and I experienced another off by one trying 3:4 for those wondering)
*/
// Assigned variable thanos to string value "Snap"
var thanos string = "Snap"

// Alllllright definately need to go over what I missed in the instructions lmao, Im seeing a pattern
func secondHalf() {

}
