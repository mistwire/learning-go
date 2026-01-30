/*
● FOR LOOPS — Go's ONLY loop construct (no while, do-while, or foreach)
● Building on Lesson 066: this is the ITERATIVE control flow mechanism
● THREE forms: C-style (init;cond;post), while-style (cond only), infinite (no cond)
● BREAK exits the loop entirely; CONTINUE skips to the next iteration
*/
package main

import (
	"fmt"
)

func main() {

	/*
		The Go for loop unifies for and while and there is no do-while.
		There are three forms, only one of which has semicolons.

		// Like a C for
		for init; condition; post { }

		// Like a C while
		for condition { }

		// Like a C for(;;)
		for { }

	*/
	// https://go.dev/doc/effective_go#for

	// FORM 1: C-style — init; condition; post
	// i is block-scoped to the loop (Lesson 033)
	for i := 0; i < 5; i++ {
		fmt.Printf("counting to five: %v \t first  for loop\n", i)
	}

	y := 5

	// FORM 2: while-style — condition only (no init or post)
	// must manage the loop variable (y++) yourself
	for y < 10 {
		fmt.Printf("y is %v \t\t\t second for loop\n", y)
		y++
	}

	// FORM 3: infinite loop — no condition (loops forever until break)
	// BREAK exits the entire loop immediately
	for {
		fmt.Printf("y is %v \t\t third  for loop\n", y)
		if y > 20 {
			break // jumps out of the for loop
		}
		y++
	}

	// CONTINUE skips the rest of the current iteration and jumps to the next one
	// here: skip odd numbers, only print even ones
	for i := 0; i < 20; i++ {
		if i%2 != 0 { // modulus (Lesson 075) — i%2 != 0 means i is odd
			continue // skip to next iteration
		}
		fmt.Println("counting even numbers: ", i)
	}

}

/*
REMARKS:
- "for" is Go's ONLY loop keyword — it replaces while, do-while, foreach, and for from other languages
- THREE forms:
    1. for init; cond; post { }  — classic counted loop
    2. for cond { }              — like while
    3. for { }                   — infinite loop (use break to exit)
- The loop variable (i) in form 1 is BLOCK-SCOPED — it doesn't exist outside the loop
- BREAK exits the innermost loop; for nested loops, use labeled breaks:
    outer: for { for { break outer } } (Lesson 073)
- CONTINUE skips to the next iteration — useful for filtering (skip odd, process even)
- There is NO do-while in Go — if you need "execute at least once," use for { ... if cond { break } }
- BUILDING ON LESSON 066: for is the third control flow category (ITERATIVE)
  after SEQUENCE (066) and CONDITIONAL (067-070)
- The init statement uses := (short declaration from Lesson 013)
  and the semicolons mirror the comma-ok idiom structure (Lesson 069)
*/
