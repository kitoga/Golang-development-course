//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing its coordinates
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import "fmt"

//* Create a rectangle structure containing its coordinates
type Rectangle struct {
	a, b Coordinates
}

type Coordinates struct {
	x, y int
}

//  - The functions must use the rectangle structure as the function parameter
func width(rect Rectangle) int {
	return rect.b.x - rect.a.x
}
func length(rect Rectangle) int {
	return rect.b.y - rect.a.y
}

//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
func area(rect Rectangle) int {
	return length(rect) * width(rect)
}

func parameter(rect Rectangle) int {
	return (width(rect) * 2) + (length(rect) * 2)
}

//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
func PrintInfo(rect Rectangle) {
	fmt.Println("Area is", area(rect))
	fmt.Println("parimeter is", parameter(rect))
}

func main() {
	//  - Print the new results to the terminal
	rect := Rectangle{a: Coordinates{30, 0}, b: Coordinates{0, 10}}
	PrintInfo(rect)
	//* After performing the above requirements, double the size
	//  of the existing rectangle and repeat the calculations
	rect.a.y *= 2
	rect.b.x *= 2
	//  - Print the new results to the terminal
	PrintInfo(rect)

}
