//Summary:
//  Print basic information to the terminal using various variable
//  creation techniques. The information may be printed using any
//  formatting you like.
//
//Requirements:
//* Store your favorite color in a variable using the `var` keyword
//* Store your birth year and age (in years) in two variables using
//  compound assignment
//* Store your first & last initials in two variables using block assignment
//* Declare (but don't assign!) a variable for your age (in days),
//  then assign it on the next line by multiplying 365 with the age
// 	variable created earlier
//
//Notes:
//* Use fmt.Println() to print out information
//* Basic math operations are:
//    Subtraction    -
// 	  Addition       +
// 	  Multiplication *
// 	  Division       /

package main

import "fmt"

func main() {

	//* Store your favorite color in a variable using the `var` keyword
	var favColor = "black"
	fmt.Println("my favorite color is", favColor)
	//* Store your birth year and age (in years) in two variables using
	//  compound assignment
	bYear, age := 1996, 26
	fmt.Println("BirthYear & Age", bYear, "&", age, "respectively")
	//* Store your first & last initials in two variables using block assignment
	var (
		fInitial = "JK"
		lInitial = "TJ"
	)
	fmt.Println("My first and last initials are", fInitial, "&", lInitial, "respectively")
	//* Declare (but don't assign!) a variable for your age (in days),
	//  then assign it on the next line by multiplying 365 with the age
	// 	variable created earlier
	var myAgeInDays int

	myAgeInDays = age * 365
	fmt.Println("my age in days", myAgeInDays)

}
