package main

import "fmt"

func main() {

	sum := 0
	fmt.Println("Sum is", sum)

	for i := 1; i <= 10; i++ {
		sum = sum + i
		fmt.Println("sum: ", sum)
	}

	for sum > 10 {
		sum -= 5
		fmt.Println("decrement of sum", sum)
	}
}
