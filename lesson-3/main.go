// Arrays and Slices
package main

import "fmt"

func main() {
	// Arrays (fixed length)
	var agesOne [3]int = [3]int{20, 25, 30}
	var agesTwo = [3]int{20, 25, 30}

	fmt.Println(agesOne, len(agesOne))
	fmt.Println(agesTwo, len(agesTwo))

	names := [4]string{"yoshi", "mario", "peach", "bowser"}
	names[1] = "luigi"
	fmt.Println(names, len(names))

	// Slices (use arrays but are not fixed length)
	var scores = []int{100, 50, 60}
	scores[2] = 25
	scores = append(scores, 85)

	fmt.Println(scores, len(scores))

	// Slice ranges
	rangeOne := names[1:3]
	rangeTwo := names[2:]
	rangeThree := names[:3]

	fmt.Println(rangeOne)
	fmt.Println(rangeTwo)
	fmt.Println(rangeThree)

	rangeOne = append(rangeOne, "koopa")
	fmt.Println(rangeOne)

}
