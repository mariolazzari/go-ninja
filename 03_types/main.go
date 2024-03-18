package main

import "fmt"

func main() {
	// string variables
	var nameOne string = "mario"
	var nameTwo = "luigi"
	var nameThree string

	fmt.Println(nameOne, nameTwo, nameThree)

	nameOne = "peach"
	nameThree = "bowser"

	fmt.Println(nameOne, nameTwo, nameThree)

	// the following is allowed inside functions only
	nameFour := "yoshi"
	fmt.Println(nameFour)

	// int variables
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40

	fmt.Println(ageOne, ageTwo, ageThree)

	// bits & memory
	// var numOne int8 = 25
	// var numTwo int8 = 128 // too large a number for 8-bit
	// var numTwo uint = -25 unsigned ints cannot be negative

	var scoreOne float32 = 25.98
	var scoreTwo float64 = 1965385877.5
	var scoreThree = 1.5 // inferred as float64

	fmt.Println(scoreOne, scoreTwo, scoreThree)

	// for more info see https://golang.org/ref/spec#Numeric_types
}
