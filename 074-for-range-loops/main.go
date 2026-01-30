/*
● FOR-RANGE LOOPS — focused lesson on ranging over data structures
● Building on Lesson 073: dedicated practice with range over slices and maps
● range returns TWO values: (index, value) for slices, (key, value) for maps
● Use _ (blank identifier) to discard either the index/key or value
*/
package main

import (
	"fmt"
)

func main() {
	// RANGE OVER A SLICE: returns (index, value) pairs
	// []int is a slice literal (covered in depth in Lesson 098)
	xi := []int{42, 43, 44, 45, 46, 47}
	for i, v := range xi {
		fmt.Println("ranging over a slice", i, v)
	}

	// RANGE OVER A MAP: returns (key, value) pairs
	// map[string]int is a map literal (covered in depth in Lesson 116)
	// iteration order is NOT deterministic
	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}
	for k, v := range m {
		fmt.Println("ranging over a map", k, v)
	}

}

/*
REMARKS:
- BUILDING ON LESSON 073: this lesson focuses exclusively on for-range
- range is the IDIOMATIC way to iterate in Go — prefer it over C-style index loops
- SLICE range returns:
    ● Two values: for i, v := range xs — index + copy of value
    ● One value: for i := range xs — index only
    ● Discard index: for _, v := range xs — value only (using blank identifier from Lesson 013)
- MAP range returns:
    ● Two values: for k, v := range m — key + copy of value
    ● One value: for k := range m — key only
    ● Map order is RANDOMIZED — never depend on iteration order
- v is always a COPY — modifying v does NOT modify the original data structure
  (this connects to pass-by-value, covered in Lesson 108 for slices)
- range also works with strings (iterates runes), channels, and integers (Go 1.22+)
- This is a foundation for the array/slice lessons (096-108) and map lessons (116-119)
*/
