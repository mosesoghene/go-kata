package main

import (
	"fmt"
)

func main() {
	product := 5
	x := 5
	product *= x
	x++
	fmt.Println("Product -> ", product)
	fmt.Println("x -> ", x)
}
