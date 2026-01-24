// LESSON: Exercise - while-style for loop (condition only)
/*
● create a for loop using only a condition
● increment or decrement the variable being checked in the condition
*/

package main

import "fmt"

func main() {
	x := 0

	for x < 10 {
		fmt.Printf("x is %v\n", x)
		x++
	}

}
