package main

import (
	"fmt"
)

func main() {

	num := 20

	for i := 1; i < num; i++ {
		if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")

		} else {
			fmt.Println("FizzBuzz")
		}

	}
}
