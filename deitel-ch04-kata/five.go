package main

import (
	"fmt"
)

func main() {
	product := 1
	c := 1
	for c <= 5 {
		product *= c
		c++
	}
	fmt.Println("Final product -> ", product)
}
