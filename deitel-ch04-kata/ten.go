package main

import (
	"fmt"
)

func main() {
	num := 5
	factorial := 1

	for i := 1; i <= num; i++ {
		factorial *= i
	}
	fmt.Println("Factorial of 5:", factorial)

	var input string
	for {
		fmt.Print("Enter text (or 'exit' to stop): ")
		fmt.Scanln(&input)
		if input == "exit" {
			break
		}
		fmt.Println("You entered:", input)
	}
}
