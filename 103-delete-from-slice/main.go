// LESSON: Delete from Slice - using append with slice syntax to remove elements
package main

import "fmt"

func main() {
	xi := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("xi - %#v\n", xi)
	fmt.Println("-------------")

	// remove the value '4' from the slice
	xi = append(xi[:4], xi[5:]...)
	fmt.Printf("xi - %#v\n", xi)
	fmt.Println("-------------")
}
