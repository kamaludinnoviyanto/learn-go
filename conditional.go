package main

import "fmt"

func main() {
	number := 10

	if number == 0 {
		fmt.Println("Number is zero.")
	} else if number%2 == 0 {
		fmt.Printf("Number %d is even.\n", number)
	} else {
		fmt.Printf("Number %d is odd.\n", number)
	}
}
