// //--Summary:
// //  Create a program to calculate the area and perimeter
// //  of a rectangle.
// //
// //--Requirements:
// //* Create a rectangle structure containing its coordinates
// //* Using functions, calculate the area and perimeter of a rectangle,
// //  - Print the results to the terminal
// //  - The functions must use the rectangle structure as the function parameter
// //* After performing the above requirements, double the size
// //  of the existing rectangle and repeat the calculations
// //  - Print the new results to the terminal
// //
// //--Notes:
// //* The area of a rectangle is length*width
// //* The perimeter of a rectangle is the sum of the lengths of all sides

// package main

// import (
// 	"fmt"
// 	//"math"
// )

// //* Create a rectangle structure containing its coordinates
// type Rectangle struct {
// 	a, b Coordinates
// }

// type Coordinates struct {
// 	x, y int
// }

// //  - The functions must use the rectangle structure as the function parameter
// func width(rect Rectangle) int {
// 	return rect.b.x - rect.a.x
// }
// func length(rect Rectangle) int {
// 	return rect.b.y - rect.a.y
// }

// //* Using functions, calculate the area and perimeter of a rectangle,
// //  - Print the results to the terminal
// func area(rect Rectangle) int {
// 	return length(rect) * width(rect)
// }

// func parameter(rect Rectangle) int {
// 	return (width(rect) * 2) + (length(rect) * 2)
// }

// //  - The functions must use the rectangle structure as the function parameter
// //* After performing the above requirements, double the size
// //  of the existing rectangle and repeat the calculations
// //  - Print the new results to the terminal
// func PrintInfo(rect Rectangle) {
// 	fmt.Println("Area is", area(rect))
// 	fmt.Println("parimeter is", parameter(rect))
// }

// //exercise # 2 employees
// type Employee struct {
// 	FirstName  string
// 	employeeID int
// }

// func changeId(e Employee, newId int) {
// 	e.employeeID = newId
// }

// func pointerChangeId(e *Employee, newId int) {
// 	e.employeeID = newId
// }

// func (e Employee) PrintGreeting() {
// 	fmt.Println("My name is", e.FirstName, "and my employee ID is", e.employeeID)
// }

// type Address struct {
// 	city    string
// 	country string
// }

// type User struct {
// 	name string
// 	age  int
// 	Address
// }

// func (u User) userPrintGreeting() {
// 	fmt.Println("My name is", u.name, "and I live in", u.city, u.country, "I am ", u.age, "years of age")
// }

// //returning pointer to structs
// func NewEmployee(name string, employeeID int) *Employee {
// 	if employeeID <= 0 {
// 		return nil
// 	} else {
// 		return &Employee{name, employeeID}
// 	}
// }

// ///
// ///
// ///
// ///
// func main() {
// 	//  - Print the new results to the terminal
// 	// rect := Rectangle{a: Coordinates{30, 0}, b: Coordinates{0, 10}}
// 	// PrintInfo(rect)
// 	// //* After performing the above requirements, double the size
// 	// //  of the existing rectangle and repeat the calculations
// 	// rect.a.y *= 2
// 	// rect.b.x *= 2
// 	// //  - Print the new results to the terminal
// 	// PrintInfo(rect)

// 	////////////////////////////exercise#2 Employee////////////////////////
// 	//simple struct
// 	var nathan Employee
// 	fmt.Println(nathan)
// 	nathan.FirstName = "Nathan"
// 	nathan.employeeID = 12345
// 	fmt.Println(nathan)

// 	heather := Employee{"heather", 22334455}
// 	fmt.Println(heather)

// 	aaron := Employee{FirstName: "aaron"}
// 	fmt.Println(aaron)

// 	//comparing structs
// 	employee1 := Employee{"moses", 12345}
// 	employee2 := employee1
// 	fmt.Println(employee2 == employee1)

// 	//accessing fields
// 	jumbo := employee1
// 	fmt.Println("my name is", jumbo.FirstName, "and my employee ID is", jumbo.employeeID)

// 	//value semantics with function
// 	mojojo := Employee{"Mojojo", 122345}
// 	fmt.Println(mojojo)
// 	changeId(mojojo, 112233)
// 	fmt.Println(mojojo)

// 	//using pointers in structs
// 	var employeePointer1 *Employee = &Employee{"Nathan", 123456}
// 	fmt.Println("getting a specific field: ", (*employeePointer1).FirstName)
// 	fmt.Println("with implicit dereferencing: ", employeePointer1.FirstName)

// 	employeePointer12 := &Employee{"nate", 11111}
// 	fmt.Println(*employeePointer12)
// 	pointerChangeId(employeePointer12, 22222)
// 	fmt.Println(*employeePointer12)

// 	//method with structs
// 	employ1 := Employee{"nAtte", 334455}
// 	employ1.PrintGreeting()

// 	//go structs promoted field
// 	p := User{
// 		name: "Joseph",
// 		age:  26,
// 		Address: Address{
// 			city:    "Lagos",
// 			country: "Nigeria",
// 		},
// 	}
// 	fmt.Println("this is the imformation on p: ", p)
// 	//p is = to the User Structs that has Address struts within
// 	p.userPrintGreeting()

// 	//returning pointer to structs
// 	newEmployee := NewEmployee("nathen", 2346566)
// 	fmt.Println(*newEmployee)

// }
