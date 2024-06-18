// Pointers
package main

import "fmt"

func updateName(x *string) {
	*x = "wedge"
}

func main() {
	name := "tifa"

	fmt.Println("name is", name)
	fmt.Println("memory address of name is", &name)

	m := &name
	fmt.Println("memory address:", m)
	fmt.Println("value at memory address:", *m)

	updateName(m)
	fmt.Println("name is", name)

}
