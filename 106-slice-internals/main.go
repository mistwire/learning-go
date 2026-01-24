// https://pkg.go.dev/runtime@go1.25.6
// https://pkg.go.dev/runtime@go1.25.6#section-sourcefiles -> "view all"
// https://cs.opensource.google/go/go/+/refs/tags/go1.25.6:src/runtime/
// https://cs.opensource.google/go/go/+/refs/tags/go1.25.6:src/runtime/slice.go

package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3, 4, 5}
	b := a 

	fmt.Println("a ", a)
	fmt.Println("b ", b)
	fmt.Println("---------------")
	
	a[0] = 7
	
	fmt.Println("a ", a)
	fmt.Println("b ", b)
	fmt.Println("---------------")
	// both slice 'a' and 'b' are looking at the same underlying array

}
