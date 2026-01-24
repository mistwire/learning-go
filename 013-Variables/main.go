// LESSON: Variables - short declaration :=, var keyword, blank identifier _, multiple assignment
package main

import "fmt"

func main() {
	a := 42
	fmt.Println(a)

	b, c, d, _, f := 0, 1, 2, 3, "happiness"
	fmt.Println(b, c, d, f)

	// this would not work
	/*
		b, c, d, e := 0, 1, 2, 3
		fmt.Println(b, c, d)
	*/

	// this does work

	var g int
	fmt.Println(g)
	g = 43
	fmt.Println(g)

	// declare a variable to hold a VALUE of a certain TYPE
	// initializing a variable
	var h int = 44
	fmt.Println(h)

}
