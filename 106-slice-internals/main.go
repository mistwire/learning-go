/*
● SLICE INTERNALS (Part 1) — shared underlying arrays
● Building on Lesson 098: assigning a slice to another variable copies the HEADER, not the data
● Both slices point to the SAME underlying array — changes through one affect the other
● This is the key behavioral difference from arrays (Lesson 096), which copy all data
*/
// https://pkg.go.dev/runtime@go1.25.6
// https://cs.opensource.google/go/go/+/refs/tags/go1.25.6:src/runtime/slice.go

package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3, 4, 5}
	b := a // copies the slice HEADER (pointer, len, cap) — NOT the underlying array

	fmt.Println("a ", a) // [0 1 2 3 4 5]
	fmt.Println("b ", b) // [0 1 2 3 4 5]
	fmt.Println("---------------")

	a[0] = 7 // modify through slice a

	fmt.Println("a ", a) // [7 1 2 3 4 5]
	fmt.Println("b ", b) // [7 1 2 3 4 5] ← b ALSO changed!
	fmt.Println("---------------")
	// both slice 'a' and 'b' are looking at the same underlying array

}

/*
REMARKS:
- THIS IS THE MOST IMPORTANT SLICE CONCEPT: slices are REFERENCE-LIKE types
- b := a copies THREE values: pointer, len, cap — but the POINTER points to the same array
  — changing a[0] changes b[0] because they share the same memory
- CONTRAST WITH ARRAYS (Lesson 096): b := a copies the ENTIRE array — they are independent
- This is why "everything in Go is pass by value" is nuanced:
    ● The slice header IS passed by value (copied)
    ● But the header contains a POINTER — so the pointed-to data is shared
- This behavior affects:
    ● Slicing (Lesson 102): sub-slices share the original's underlying array
    ● Passing slices to functions (Lesson 108): the function can modify the caller's data
    ● Delete operations (Lesson 103): modifying the underlying array in place
- To create a truly INDEPENDENT copy, use copy() — see Lesson 107
- BUILDING ON LESSON 098: this is why understanding the slice struct (pointer, len, cap) matters
  — the pointer is what creates the shared reference behavior
*/
