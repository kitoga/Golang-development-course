//--Summary:
//  Implement the classic "FizzBuzz" problem using a `for` loop.
//
//--Requirements:
//* Print integers 1 to 50, except:
//  - Print "Fizz" if the integer is divisible by 3
//  - Print "Buzz" if the integer is divisible by 5
//  - Print "FizzBuzz" if the integer is divisible by both 3 and 5
//
//--Notes:
//* The remainder operator (%) can be used to determine divisibility

package main

import "fmt"

func main() {

	sum := 0
	//* Print integers 1 to 50, except:
	for j := 1; j <= 10; j++ {
		sum = sum + 1

		//  - Print "Fizz" if the integer is divisible by 3
		if sum%3 == 0 {
			fmt.Println("fuzz")
			continue
		} else if sum%5 == 0 {
			//  - Print "Buzz" if the integer is divisible by 5
			fmt.Println("buzz")
		} else if (sum%3 == 0) && (sum&5 == 0) {
			//  - Print "FizzBuzz" if the integer is divisible by both 3 and 5
			fmt.Println("FuzzBuzz")
		} else {
			// not divisible
			fmt.Println("Not Divisible")
		}
	}

}
