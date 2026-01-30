/*
● SLICING A SLICE — creating sub-slices with [low:high] syntax
● Building on Lesson 098: slicing creates a NEW slice header pointing to the SAME underlying array
● Syntax: slice[inclusive:exclusive] — low index included, high index excluded
● Shorthand forms: [:n], [n:], [:]
*/
package main

import "fmt"

func main() {
	xi := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("xi - %#v\n", xi)
	fmt.Println("-------------")

	// [inclusive:exclusive] — elements at index 0, 1, 2, 3 (NOT 4)
	fmt.Printf("xi - %#v\n", xi[0:4]) // [0, 1, 2, 3]
	fmt.Println("-------------")

	// [:exclusive] — from start (0) to index 6
	fmt.Printf("xi - %#v\n", xi[:7]) // [0, 1, 2, 3, 4, 5, 6]
	fmt.Println("-------------")

	// [inclusive:] — from index 4 to end
	fmt.Printf("xi - %#v\n", xi[4:]) // [4, 5, 6, 7, 8, 9]
	fmt.Println("-------------")

	// [:] — entire slice (same as xi itself)
	fmt.Printf("xi - %#v\n", xi[:]) // [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
	fmt.Println("-------------")
}

/*
REMARKS:
- Syntax: slice[low:high] — INCLUSIVE low, EXCLUSIVE high (like Python)
  ● xi[0:4] returns elements at indices 0, 1, 2, 3 (length = high - low = 4)
- Shorthand forms:
    ● xi[:n]  — from beginning to n (same as xi[0:n])
    ● xi[n:]  — from n to end (same as xi[n:len(xi)])
    ● xi[:]   — entire slice (same as xi[0:len(xi)])
- The result is a NEW SLICE HEADER, but it shares the SAME UNDERLYING ARRAY
  — modifying the sub-slice affects the original! (demonstrated in Lesson 106)
- This is used in Lesson 103 for DELETING elements: append(xi[:i], xi[i+1:]...)
  — "everything before index i" + "everything after index i"
- The sub-slice's CAPACITY extends to the end of the underlying array
  — cap(xi[2:5]) includes elements from index 2 to the end of xi's underlying array
- BUILDING ON LESSON 098: slicing is possible because slices are headers over arrays
  — creating a sub-slice is O(1) — just creates a new header, no data is copied
*/
