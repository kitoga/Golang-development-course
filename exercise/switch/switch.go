//--Summary:
//  Create a program to display a classification based on age.
//
//--Requirements:
//* Use a `switch` statement to print the following:
//  - "newborn" when age is 0
//  - "toddler" when age is 1, 2, or 3
//  - "child" when age is 4 through 12
//  - "teenager" when age is 13 through 17
//  - "adult" when age is 18+

package main

import "fmt"

func age() int {
	return 12
}

func main() {

	switch ageNum := age(); {
	case ageNum == 0:
		fmt.Println("newborn")
	case ageNum >= 1 && ageNum <= 3:
		fmt.Println("toddler")
	case ageNum >= 4 && ageNum <= 12:
		fmt.Println("child")
	case ageNum >= 13 && ageNum <= 17:
		fmt.Println("Teenager")
	default:
		fmt.Println("adult")
	}
}
