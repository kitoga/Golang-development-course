package main

import "fmt"

func main() {

	shoppingList := make(map[string]int)
	shoppingList["eggs"] = 7
	shoppingList["milk"] = 2
	shoppingList["bread"] = 2

	shoppingList["eggs"] += 1
	fmt.Println(shoppingList)

	delete(shoppingList, "milk")
	fmt.Println("Milk deleted from shoppingList", shoppingList)

	fmt.Println("need", shoppingList["eggs"], "eggs")

	coco, found := shoppingList["coco"]
	fmt.Println("need coco?", coco)
	if !found {
		fmt.Println("coco not found!")
		return
	} else {
		fmt.Println("coco is", shoppingList["coco"], "on th shopping List")
	}

	// totalItem := 0
	// for _, amount := range shoppingList {
	// 	totalItem = totalItem + amount
	// }
	fmt.Println("total length", len(shoppingList))
}
