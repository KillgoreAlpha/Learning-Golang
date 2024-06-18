// Maps
package main

import "fmt"

func main() {
	// Maps (Similar to python dictsionaries) -> map[key type]value type{key: value}
	menu := map[string]float64{
		"soup":           4.99,
		"pie":            7.99,
		"salad":          6.99,
		"toffee pudding": 3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	// looping through maps
	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	// Ints as key type
	phonebook := map[int]string{
		267584967: "mario",
		984759373: "luigi",
		845775485: "peach",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[267584967])

	// Update item in map
	phonebook[267584967] = "bowser"
	fmt.Println(phonebook)
	fmt.Println(phonebook[267584967])

	phonebook[845775485] = "yoshi"
	fmt.Println(phonebook)
	fmt.Println(phonebook[845775485])

}
