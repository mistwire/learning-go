/*
● NESTED LOOPS — loops inside loops
● Building on Lesson 072: the inner loop runs COMPLETELY for each iteration of the outer loop
● Also introduces FOR-RANGE — Go's idiomatic way to iterate over data structures
● Preview of SLICES (Lesson 098) and MAPS (Lesson 116)
*/
package main

import (
	"fmt"
)

func main() {
	// NESTED LOOPS: outer runs 5 times, inner runs 5 times PER outer iteration
	// total iterations: 5 × 5 = 25
	for i := 0; i < 5; i++ {
		fmt.Println("--")
		for j := 0; j < 5; j++ {
			fmt.Printf("outer loop %v \t inner loop %v\n", i, j)
		}
	}

	// FOR-RANGE over a SLICE — returns (index, value) for each element
	xi := []int{42, 43, 44, 45, 46, 47}
	for i, v := range xi {
		fmt.Println("ranging over a slice", i, v) // i=index, v=copy of value
	}

	// FOR-RANGE over a MAP — returns (key, value) for each entry
	// NOTE: map iteration order is NOT guaranteed (randomized by Go)
	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}
	for k, v := range m {
		fmt.Println("ranging over a map", k, v) // k=key, v=copy of value
	}

}

/*
REMARKS:
- NESTED LOOPS: the inner loop completes ALL iterations before the outer loop increments
  — O(n×m) total iterations for two nested loops
- break/continue in nested loops only affect the INNERMOST loop
  — to break out of an outer loop, use a LABEL: outer: for { for { break outer } }
- FOR-RANGE is the idiomatic way to iterate over Go data structures:
    ● Slices: for i, v := range slice { }   — i=index, v=value
    ● Maps: for k, v := range m { }         — k=key, v=value
    ● Strings: for i, r := range s { }      — i=byte index, r=rune (Unicode code point)
    ● Channels: for v := range ch { }       — receives until channel is closed
- range returns COPIES of values — modifying v doesn't change the original slice/map
- Use _ (blank identifier, Lesson 013) to discard unwanted values:
    for _, v := range slice { }   — index discarded
    for k := range m { }          — value discarded (just keys)
- BUILDING ON LESSON 072: for-range is the fourth form of for loop
- Map iteration order is INTENTIONALLY RANDOMIZED by Go — never rely on it
*/
