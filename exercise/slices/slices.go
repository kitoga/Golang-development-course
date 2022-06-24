//--Summary:
//  Create a program to manage parts on an assembly line.
//
//--Requirements:
//* Using a slice, create an assembly line that contains type Part
//* Create a function to print out the contents of the assembly line
//* Perform the following:
//  - Create an assembly line having any three parts
//  - Add two new parts to the line
//  - Slice the assembly line so it contains only the two new parts
//  - Print out the contents of the assembly line at each step
//--Notes:
//* Your program output should list 3 parts, then 5 parts, then 2 parts

package main

import "fmt"

type Part string

//* Create a function to print out the contents of the assembly line
func assemblyLinePrint(line []Part) {
	for i := 0; i < len(line); i++ {
		lineItem := line[i]
		fmt.Println("line item: ", lineItem)
	}
}

func main() {
	//fmt.Println("jbvbv")

	//* Perform the following:
	//  - Create an assembly line having any three parts
	line := []Part{"automobile-", "brakes-", "wheels-"}
	fmt.Println("3 parts:")
	assemblyLinePrint(line)
	//  - Add two new parts to the line
	line = append(line, "appliances-", "electronic")
	fmt.Println("\n5 parts:")
	assemblyLinePrint(line)
	//  - Slice the assembly line so it contains only the two new parts
	line = line[3:]
	fmt.Println("\n2 parts:")
	assemblyLinePrint(line)
	//  - Print out the contents of the assembly line at each step
	fmt.Println("Done")
}
