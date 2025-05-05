package main

import "fmt"

func main() {
	x := 5
	x = x + 1
	x += 1
	x++
	x = x + 2 - 1
	fmt.Println("x after increments:", x)
}
