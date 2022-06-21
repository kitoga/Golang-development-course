//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import "fmt"

//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
func name(name string) {
	fmt.Println("hello there", name)
}

//* Write a function that returns any message, and call it from within
//  fmt.Println()
func message() string {
	return "Hello there function exercise"
}

//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
func addThreeNumber(num1, num2, num3 int) int {
	return num1 + num2 + num3
}

//* Write a function that returns any number
func anyNumber() int {
	return 5
}

//* Write a function that returns any two numbers
func returnTwoNumbers() (int, int) {
	return 2, 2
}

func main() {
	name("Joseph")
	fmt.Println(message())
	a, b := returnTwoNumbers()
	//* Add three numbers together using any combination of the existing functions.
	answer := addThreeNumber(anyNumber(), a, b) //5 + 2+2=9

	//  * Print the result
	fmt.Println("add three numbers will be ", answer)

	//* Call every function at least once

}
