package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3, 4, 5}
	// make a new slice 'b' instead of copying 'a' - which also has a different underlying array
	b := make([]int, 6)
	
	// use a copy function to copy the info from a's underlying array to b's
	copy(b, a)

	fmt.Println("a ", a)
	fmt.Println("b ", b)
	fmt.Println("--------------")

	a[0] = 7

	fmt.Println("a ", a)
	fmt.Println("b ", b)
	fmt.Println("--------------")
	// now slice b isn't affected by the change made in slice a
}