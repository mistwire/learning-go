/*
● ARRAYS — fixed-size, ordered collections of a single type
● The SIZE is part of the type: [3]int and [2]int are DIFFERENT types
● Three ways to create: explicit size [3]int{}, compiler-counted [...]int{}, var declaration
● Arrays are VALUE types — assigning or passing copies the ENTIRE array
● Arrays are rarely used directly — SLICES (Lesson 098) are preferred
*/
package main

import "fmt"

func main() {

	// ARRAY LITERAL with explicit size
	a := [3]int{42, 43, 44}
	fmt.Println(a)

	// [...] lets the compiler COUNT the elements and set the size
	b := [...]string{"Hello", "Gophers!"}
	fmt.Println(b)

	// var declaration — creates array with ZERO VALUES, then assign by index
	var c [2]int
	c[0] = 7
	c[1] = 8
	fmt.Println(c)

	// %T shows the type — notice the size IS the type
	fmt.Printf("%T\n", a) // [3]int
	fmt.Printf("%T\n", c) // [2]int
	// c = a  ← won't compile! [3]int and [2]int are DIFFERENT types

	{
		var ni [7]int        // zero-value array: [0 0 0 0 0 0 0]
		ni[0] = 42           // assign by index
		fmt.Printf("%#v \t\t\t\t %T\n", ni, ni)

		ni2 := [4]int{55, 56, 57, 58} // explicit size
		fmt.Printf("%#v \t\t\t\t\t %T\n", ni2, ni2)

		ns := [...]string{"Chocolate", "Vanilla", "Strawberry"} // compiler-counted
		fmt.Printf("%#v \t %T\n", ns, ns)

		// len() returns the number of elements — builtin function
		// https://pkg.go.dev/builtin#len
		fmt.Println(len(ni))  // 7
		fmt.Println(len(ni2)) // 4
		fmt.Println(len(ns))  // 3
	}
}

/*
"Arrays have their place, but they're a bit inflexible,
so you don't see them too often in Go code.
Slices, though, are everywhere.
They build on arrays to provide
great power and convenience."
~ Go Slices: usage and internals - Go Blog - Andrew Gerrand
*/
// https://go.dev/blog/slices-intro

/*
REMARKS:
- Arrays are FIXED SIZE — the size is part of the type and cannot change
  [3]int ≠ [4]int ≠ [5]int — they are all DIFFERENT types
- Arrays are VALUE TYPES — copying or passing an array creates a FULL COPY
  (unlike slices which share an underlying array — Lesson 106)
- Three creation patterns:
    1. [3]int{1, 2, 3} — explicit size
    2. [...]int{1, 2, 3} — compiler counts for you (still fixed after creation)
    3. var x [3]int — zero-value array, assign elements by index
- %#v format verb prints the Go-syntax representation (useful for debugging)
- len() is a BUILTIN function — works on arrays, slices, maps, strings, channels
- Arrays are the FOUNDATION for slices (Lesson 098) — every slice has an underlying array
- In practice, SLICES are used almost everywhere instead of arrays because:
    ● Slices can grow/shrink dynamically
    ● Slices are passed as lightweight headers (pointer + len + cap), not copied entirely
    ● Most standard library functions work with slices, not arrays
- BUILDING ON LESSON 074: you can use for-range to iterate over arrays too
*/
