package main

import (
	"fmt"
)

func main() {
	sum := 0
	z := 5
	for z >= 0 {
		sum += z
		z--
	}
	fmt.Println("Final sum -> ", sum)
}
