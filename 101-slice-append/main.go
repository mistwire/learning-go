// LESSON: Slice Append - append() builtin for growing slices dynamically
package main

import "fmt"

func main() {
	xi := []int{42, 43, 44}
	fmt.Println(xi)
	fmt.Println("------------")

	xi = append(xi, 44, 48, 99, 100, 777)
	fmt.Println(xi)
	fmt.Println("------------")

}
