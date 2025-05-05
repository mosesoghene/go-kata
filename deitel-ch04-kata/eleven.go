package main

import "fmt"

func main() {
	x, y := 7, 3
	x = y
	y++
	fmt.Println("After x = y++, x:", x, "y:", y)
	y++
	x = y
	fmt.Println("After x = ++y, x:", x, "y:", y)
}
