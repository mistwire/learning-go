/*
● SLICES — dynamic, flexible views into underlying arrays
● KEY DIFFERENCE FROM LESSON 096: slices have NO fixed size — []string vs [3]string
● A slice is a HEADER with three fields: pointer, length, capacity
● Slices are the workhorse data structure in Go — used far more than arrays
*/
package main

import "fmt"

func main() {
	// []string (no size) = SLICE — dynamic, flexible
	// [2]string (with size) = ARRAY — fixed, rigid (Lesson 096)
	xs := []string{"hello", "world"}
	fmt.Println(xs)
}

/*
	A slice is a struct with THREE fields internally:
	-- pointer to an underlying array (where the data actually lives)
	-- len (number of elements currently in use)
	-- cap (total capacity of the underlying array)

	type slice struct {
		array unsafe.Pointer
		len   int
		cap   int
	}
*/
// src/runtime/slice.go

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
- KEY DIFFERENCE FROM ARRAYS (Lesson 096): no size in the type — []int vs [3]int
    ● []int is a SLICE — can grow with append() (Lesson 101)
    ● [3]int is an ARRAY — fixed at compile time
- A slice is a lightweight HEADER (24 bytes on 64-bit) containing:
    1. Pointer — to the first element in the underlying array
    2. Length — number of elements currently accessible (len())
    3. Capacity — total elements the underlying array can hold before reallocation (cap())
- Slices are REFERENCE types — assigning a slice creates a new header pointing to the SAME array
  (this has important implications — see Lesson 106)
- Multiple slices can SHARE the same underlying array (Lesson 106)
- Common operations (upcoming lessons):
    ● for-range (100), append (101), slicing (102), delete (103), make (104), multidimensional (105)
- BUILDING ON LESSON 096: arrays are the foundation; slices sit on top providing flexibility
- Zero value of a slice is nil (not an empty slice): var xs []int → xs == nil is true
*/
