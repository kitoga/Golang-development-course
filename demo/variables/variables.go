package main

import "fmt"

func main() {
	var name = "joseph"
	fmt.Println(name)

	//type notation
	var myName string = "Theo"
	fmt.Println("name", myName)

	myAdmin := "Admin"
	fmt.Println("admin", myAdmin)

	part1, part2 := 1, 5
	fmt.Println("starting part is", part1, "ending part is", part2)

	var sum int
	fmt.Println(sum)

	sum = part1 * part2
	fmt.Println("sum of the parts is", sum)

	var (
		lessonName = "Variables"
		lessonType = "demo"
	)
	fmt.Println("lessonName", lessonName)
	fmt.Println("lessonType", lessonType)

	word1, word2, _ := "hello", "world", "!"
	fmt.Println(word1, word2)
}
