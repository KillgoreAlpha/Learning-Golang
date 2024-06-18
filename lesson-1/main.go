// Variables, Strings, and Numbers
package main

import "fmt"

var someName string

func main() {
	// Multiple ways to assign a variable

	// Strings
	var nameOne string = "mario"
	var nameTwo = "luigi"
	var nameThree string

	fmt.Println(nameOne, nameTwo, nameThree)

	nameOne = "peach"
	nameThree = "bowser"

	fmt.Println(nameOne, nameTwo, nameThree)

	// Only use := on first var init and only inside of a function
	nameFour := "yoshi"

	fmt.Println(nameFour)

	// Integers
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40

	fmt.Println(ageOne, ageTwo, ageThree)

	// Bits and memory
	var numOne int8 = 25
	var numTwo int8 = -128
	var numThree uint16 = 256

	fmt.Println(numOne, numTwo, numThree)

	// Floats
	var scoreOne float32 = -1.5
	var scoreTwo float64 = 1234567890.0987654321
	scoreThree := 1.5

	fmt.Println(scoreOne, scoreTwo, scoreThree)

}
