package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func printTrip(mile float64, gallon float64) {
	fmt.Printf("%.2f | %.2f\n", mile, gallon)
}

func main() {
	trips := make(map[float64]float64)
	for {
		mile, galon := collectInput()
		trips[mile] = galon
		fmt.Print("Enter new trip? (y/n) : ")
		response, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		response = strings.ToLower(strings.TrimSpace(response))
		if response == "y" || response == "yes" {
			continue
		} else if response == "n" || response == "no" {
			fmt.Println("Mile | Gallon")
			fmt.Println("-------------")
			printTrips(trips)
			break
		} else {
			fmt.Println("Invalid input")
			continue
		}
	}

	var car Car = Car{"Camry", "Totota"}

	fmt.Println(car.Make())
}

func printTrips(trips map[float64]float64) {
	for mile, trip := range trips {
		printTrip(mile, trip)
	}
}

type Car struct {
	model string
	make  string
}

func (c Car) Make() string {
	return c.make
}

func (c Car) Model() string {
	return c.model
}

func (c Car) SetModel(model string) {
	c.model = model
}
