package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func collectInput() (float64, float64) {
	scanner := bufio.NewReader(os.Stdin)
	fmt.Print("Enter mile: ")
	mileInput, _ := scanner.ReadString('\n')
	fmt.Print("Enter gallon: ")
	gallonInput, _ := scanner.ReadString('\n')

	mile, _ := strconv.ParseFloat(mileInput[:len(mileInput)-1], 64)
	gallon, _ := strconv.ParseFloat(gallonInput[:len(gallonInput)-1], 64)

	return mile, gallon
}

func main() {

	mile, galon := collectInput()
	fmt.Println("Mile:", mile)
	fmt.Println("Galon:", galon)
}
