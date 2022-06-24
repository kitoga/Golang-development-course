//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

type Products struct {
	name  string
	price int
}

func printStats(list [4]Products) {
	var cost, totalItem int

	for i := 0; i < len(list); i++ {
		item := list[i]
		cost = cost + item.price

		if item.name != "" {
			totalItem += 1
		}
	}
	fmt.Println("last item on the list:", list[totalItem-1])
	fmt.Println("total item on the list:", totalItem)
	fmt.Println("Total cost:", cost)
}

func main() {

	//solution #1

	// //--Requirements:
	// //* Using an array, create a shopping list with enough room
	// //  for 4 products
	var shoppingList = [4]Products{
		{"Beans", 17},
		{"Bread", 8},
		{"Milk", 23},
	}

	printStats(shoppingList)

	// //* Add a fourth product to the list and print out the
	// //  information again
	shoppingList[3] = Products{"sardine", 17}

	printStats(shoppingList)

	// //  - Products must include the price and the name
	// fmt.Println("4 products inserted")
	// //* Insert 3 products into the array
	// //* Print to the terminal:
	// //  - The last item on the list
	// fmt.Println("The last item on the kist: ", myArray[2])
	// //  - The total number of items
	// fmt.Println("The total number of Items: ", len(myArray))
	// //  - The total cost of the items
	// fmt.Println("The total cost of all the arrays: ", myArray[0].price+myArray[1].price+myArray[2].price)
	// //* Add a fourth product to the list and print out the
	// //  information again
	// myArray[3] = Products{"biscuits", 17}

	// fmt.Println(myArray)

}
