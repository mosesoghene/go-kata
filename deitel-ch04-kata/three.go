package main

import (
	"fmt"
)

func main() {
	sum := 0
	x := 1
	for x <= 10 {
		sum += x
		x++
	}
	fmt.Println("The sum of numbers from 1 to 10 is -> ", sum)
}
