// Pass By Value
package main

import "fmt"

func updateName(x string) string {
	x = "wedge"
	return x
}

func updateMenu(y map[string]float64) {
	y["coffee"] = 2.99
}

func main() {
	// Group A types -> strings, ints, bools, floats, arrays, structs (non-pointer values)
	// (data is copied to new memory location, original is unchanged)
	name := "tifa"

	updateName(name)
	fmt.Println(name)

	// Must update variable
	name = updateName(name)
	fmt.Println(name)

	// Group B types -> slices, maps, functions (pointer wrapper values)
	// (pointer to data is copied, not the data pointed to so referenced data is changed)
	menu := map[string]float64{
		"pie":       5.95,
		"ice cream": 3.99,
	}

	// Don't need to update variable
	updateMenu(menu)
	fmt.Println(menu)

}
