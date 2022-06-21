package main

import "fmt"

func price() int {
	return 12
}

const (
	Economy    = 0
	Business   = 1
	FirstClass = 2
)

func main() {

	switch p := price(); {
	case p < 2:
		fmt.Println("cheaper item")
	case p < 10:
		fmt.Println("moderately Priced item")
	default:
		fmt.Println("expensive shiihh")
	}

	c := Economy
	switch c {
	case Economy:
		fmt.Println("economy")
	case Business:
		fmt.Println("business")
	case FirstClass:
		fmt.Println("first class")
	default:
		fmt.Println("Others")
	}

}
