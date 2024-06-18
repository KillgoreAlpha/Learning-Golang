// Printing and Forrmatting
package main

import "fmt"

func main() {
	age := 35
	name := "shaun"

	// Print (doesn't add new line)
	fmt.Print("Hello, ")
	fmt.Print("world! \n")
	fmt.Print("new line \n")

	// Println (adds new line)
	fmt.Println("Hello, world!")
	fmt.Println("Bye, world!")
	fmt.Println("my age is", age, "and my name is", name)

	// Printf (formatted string) %_ = format specifier
	fmt.Printf("my age is %v and my name is %v \n", age, name)
	fmt.Printf("my age is %q and my name is %q \n", age, name)
	fmt.Printf("age is of type %T \n", age)
	fmt.Printf("you scored %f points! \n", 225.55)
	fmt.Printf("you scored %0.1f points! \n", 225.55)

	// Sprintf (save formatted strings)
	var str = fmt.Sprintf("my age is %v and my name is %v \n", age, name)
	fmt.Println("the saved string is:", str)

}
