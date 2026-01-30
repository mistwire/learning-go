/*
● SLICE INTERNALS (Part 3) — pass by value and function side effects
● Building on Lessons 106-107: passing a slice to a function copies the HEADER,
  but the function can still modify the caller's data through the shared pointer
● medianOne: SORTS THE ORIGINAL slice (side effect — modifies caller's data)
● medianTwo: uses make + copy to work on an INDEPENDENT copy (no side effect)
● "Everything in Go is pass by value" — but the slice value IS a header with a pointer
*/
package main

import (
	"fmt"
	"sort"
)

func main() {
	n := []float64{3, 1, 4, 2}

	fmt.Println(medianOne(n))
	fmt.Println("after medianOne", n) // n is NOW SORTED [1 2 3 4] — side effect!

	y := []float64{3, 1, 4, 2}
	fmt.Println(medianTwo(y))
	fmt.Println("after medianTwo", y) // y is UNCHANGED [3 1 4 2] — no side effect
}

// medianOne: sorts the CALLER'S slice because x shares the same underlying array
// the slice header is copied (pass by value), but the pointer still points to n's array
func medianOne(x []float64) float64 {
	sort.Float64s(x) // MODIFIES the underlying array — affects the caller!
	i := len(x) / 2
	if len(x)%2 == 1 {
		return x[i/2]
	}
	return (x[i-1] + x[i]) / 2
}

// medianTwo: makes an INDEPENDENT copy before sorting — caller's data is safe
func medianTwo(x []float64) float64 {
	// allocate a new underlying array and copy the data (Lesson 107 pattern)
	n := make([]float64, len(x))
	copy(n, x)

	sort.Float64s(n) // sorts the copy, NOT the original
	i := len(n) / 2
	if len(n)%2 == 1 {
		return n[i/2]
		// integer division truncates (Lesson 075)
	}
	return (n[i-1] + n[i]) / 2
}

/*
REMARKS:
- THIS LESSON TIES TOGETHER LESSONS 106 + 107 in a practical scenario
- Go is ALWAYS pass-by-value, but for slices the "value" is the header (pointer, len, cap)
    ● The pointer is copied — but it still points to the SAME underlying array
    ● So the function CAN modify the caller's data through that shared pointer
- medianOne demonstrates the PROBLEM: sort.Float64s(x) mutates the original slice
    ● The caller's data is permanently reordered — this is a SIDE EFFECT
    ● This is often a bug, not a feature
- medianTwo demonstrates the SOLUTION: make + copy before mutating
    ● Creates an independent slice with its own underlying array
    ● The caller's data is untouched
- RULE OF THUMB: if a function needs to modify a slice (sort, shuffle, etc.),
  make a copy first UNLESS you intentionally want to modify the caller's data
- This pattern (copy-then-mutate) appears frequently in real Go code
- BUILDING ON LESSON 075: integer division (len/2) truncates — used here for median index
- BUILDING ON LESSON 107: make + copy is the standard pattern for independent copies
- This is the last slice-internals lesson — the complete mental model is:
    Slices are headers → headers contain pointers → pointers share arrays → use copy() for independence
*/
