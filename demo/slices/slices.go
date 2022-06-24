package main

import "fmt"

func printSlice(title string, slice []string) {
	fmt.Println()
	fmt.Println("---", title, "---")

	for i := 0; i < len(slice); i++ {
		element := slice[i]
		fmt.Println(element)
	}
}

func main() {

	routeSlice := []string{"benin", "delta", "Ph", "osun"}
	printSlice("Route1", routeSlice)

	routeSlice = append(routeSlice, "abj", "ibadan", "ekiti", "ogun")
	printSlice("Route2", routeSlice)

	//visited routes
	fmt.Println(routeSlice[1], "visited")
	fmt.Println(routeSlice[5], "visited")
	fmt.Println(routeSlice[7], "visited")

	//reslice an existing slice
	notVisitedSlice := routeSlice[2:]
	fmt.Println(notVisitedSlice, "not visited")

	//fmt.Println(routeSlice)

}
