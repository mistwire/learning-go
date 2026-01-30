/*
● MAKE — pre-allocating slices with make([]T, length, capacity)
● Building on Lesson 098: make creates a slice with a specified CAPACITY
  to avoid repeated re-allocations when appending
● len = number of accessible elements; cap = total allocated space
● When len exceeds cap, Go allocates a NEW, LARGER underlying array (roughly 2x)
*/
package main

import "fmt"

func main() {
	// make(type, length, capacity)
	// length 0: no elements yet; capacity 10: room for 10 before reallocation
	xi := make([]int, 0, 10)
	fmt.Println(xi)
	fmt.Printf("length: %v\n", len(xi))  // 0 — no elements
	fmt.Printf("capacity: %v\n", cap(xi)) // 10 — pre-allocated space
	fmt.Println("------------")

	// append 10 elements — fits within capacity (10), no reallocation needed
	xi = append(xi, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(xi)
	fmt.Printf("length: %v\n", len(xi))  // 10
	fmt.Printf("capacity: %v\n", cap(xi)) // 10 — exactly filled
	fmt.Println("------------")

	// append 4 more — EXCEEDS capacity → Go allocates a new, larger underlying array
	xi = append(xi, 10, 11, 12, 13)
	fmt.Println(xi)
	fmt.Printf("length: %v\n", len(xi))  // 14
	fmt.Printf("capacity: %v\n", cap(xi)) // 20 (roughly doubled)
}

/*
REMARKS:
- make() is a BUILTIN that creates slices, maps, and channels
- For slices: make([]T, length, capacity)
    ● length: number of elements initialized to zero values
    ● capacity: size of the underlying array (pre-allocated space)
    ● capacity is optional: make([]int, 5) creates len=5, cap=5
- WHY USE MAKE: if you know approximately how many elements you'll need,
  pre-allocating avoids repeated re-allocations as the slice grows
- GROWTH STRATEGY: when append exceeds capacity, Go allocates a new array
    ● Roughly doubles capacity for small slices
    ● Growth factor decreases for very large slices (~1.25x)
    ● Old array is garbage collected once no slices reference it
- BUILDING ON LESSON 098: make directly controls the slice's three internal fields:
    ● Pointer → to a newly allocated array of size cap
    ● Len → set to the length argument
    ● Cap → set to the capacity argument
- BUILDING ON LESSON 101: append to a make'd slice with len=0 is efficient
  — it fills the pre-allocated space before needing to grow
- cap() is a builtin like len() — returns the capacity of a slice
- Common pattern: xs := make([]int, 0, expectedSize) then append in a loop
*/

