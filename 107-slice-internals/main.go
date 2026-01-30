/*
● SLICE INTERNALS (Part 2) — copy() for independent slices
● Building on Lesson 106: to AVOID shared underlying arrays, use make() + copy()
● copy(dst, src) copies elements from src into dst's OWN underlying array
● After copy, the two slices are fully INDEPENDENT — changes don't affect each other
*/
package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3, 4, 5}
	// make() creates a NEW slice with its OWN underlying array (Lesson 104)
	b := make([]int, 6)

	// copy(destination, source) — copies elements from a's array into b's array
	// returns the number of elements copied (min of len(dst), len(src))
	copy(b, a)

	fmt.Println("a ", a) // [0 1 2 3 4 5]
	fmt.Println("b ", b) // [0 1 2 3 4 5] — same values, DIFFERENT underlying array
	fmt.Println("--------------")

	a[0] = 7

	fmt.Println("a ", a) // [7 1 2 3 4 5]
	fmt.Println("b ", b) // [0 1 2 3 4 5] ← b is NOT affected — independent copy!
	fmt.Println("--------------")
}

/*
REMARKS:
- KEY DIFFERENCE FROM LESSON 106: b := a shares the array; make + copy creates independence
- THE PATTERN for a true copy:
    1. b := make([]int, len(a))  — allocate a new slice with its own array (Lesson 104)
    2. copy(b, a)                — copy element values from a into b
- copy() is a BUILTIN function — copies elements from src to dst
    ● Returns the number of elements copied: min(len(dst), len(src))
    ● If dst is shorter, only len(dst) elements are copied
    ● If dst is longer, extra elements keep their zero values
- After copy, the two slices are completely INDEPENDENT:
    ● Different underlying arrays
    ● Different pointers in their slice headers
    ● Modifying one does NOT affect the other
- WHEN TO USE copy vs simple assignment:
    ● b := a — when you WANT shared data (lightweight, no allocation)
    ● copy(b, a) — when you need ISOLATION (e.g., sorting without affecting the original — Lesson 108)
- BUILDING ON LESSON 104: make() creates the new underlying array; copy() fills it
*/