/*
● SLICE APPEND — growing slices dynamically with the append() builtin
● Building on Lesson 098: slices can GROW (unlike fixed-size arrays from Lesson 096)
● append() returns a NEW slice — you MUST reassign: xs = append(xs, values...)
● Can append one or MULTIPLE values in a single call
*/
package main

import "fmt"

func main() {
	xi := []int{42, 43, 44}
	fmt.Println(xi) // [42 43 44] — len=3
	fmt.Println("------------")

	// append multiple values at once — MUST reassign to xi
	xi = append(xi, 44, 48, 99, 100, 777)
	fmt.Println(xi) // [42 43 44 44 48 99 100 777] — len=8
	fmt.Println("------------")

}

/*
REMARKS:
- append() is a BUILTIN function — no import needed
- Signature: append(slice, elements...) — takes a slice and one or more values
- You MUST reassign the result: xi = append(xi, ...)
  — append MAY return a new underlying array if capacity is exceeded (Lesson 104)
  — forgetting to reassign is a common bug
- HOW IT WORKS INTERNALLY (ties to Lesson 098's slice struct):
    ● If len < cap: appends in place, increments len
    ● If len == cap: allocates a NEW, LARGER underlying array, copies data, returns new slice header
- To append one slice to another, use the ... (variadic) operator:
    xs = append(xs, ys...) — unpacks ys into individual elements
- BUILDING ON LESSON 096: this is WHY slices are preferred over arrays
  — arrays can't grow; slices grow dynamically with append
- BUILDING ON LESSON 098: append may change the underlying array pointer and capacity
  — this is why understanding the slice header (pointer, len, cap) matters
*/
