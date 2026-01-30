/*
● SLICE FOR-RANGE — iterating over slices, index access, and len()
● Building on Lessons 074 & 098: applying for-range specifically to slices
● Access elements by INDEX: xs[0], xs[1], xs[2]
● Accessing beyond the length (xs[3]) causes a RUNTIME PANIC
● Comparing range-based vs C-style iteration
*/
package main

import "fmt"

func main() {
	xs := []string{"Almond Biscotti Café", "Banana Pudding ", "Balsamic Strawberry (GF)"}

	fmt.Println(xs)
	fmt.Printf("%T\n", xs) // []string — slice of strings

	// FOR-RANGE with blank identifier _ (Lesson 013) — discards the index, keeps the value
	for _, v := range xs {
		fmt.Printf("%v\n", v)
	}

	fmt.Println("---------------------")
	// DIRECT INDEX ACCESS: zero-based indexing
	fmt.Println(xs[0]) // first element
	fmt.Println(xs[1]) // second element
	fmt.Println(xs[2]) // third element

	// xs[3] would cause a RUNTIME PANIC: "index out of range"
	// Go does NOT silently return undefined/null — it crashes
	//fmt.Println(xs[3])

	fmt.Println("---get len of a slice------")
	fmt.Println(len(xs)) // 3 — number of elements

	// C-style loop — the "old way" before for-range
	// works but for-range is more idiomatic
	fmt.Println("---the old way------")
	for i := 0; i < len(xs); i++ {
		fmt.Println(i) // prints indices (0, 1, 2), not values
	}
}

/*
REMARKS:
- BUILDING ON LESSONS 074 & 098: this applies for-range to slice iteration
- INDEX ACCESS: xs[i] — zero-based, accessing beyond len causes a PANIC (not a compile error)
  — Go's bounds checking happens at RUNTIME, not compile time
- for-range is PREFERRED over C-style loops because:
    ● Cleaner syntax: for _, v := range xs vs for i := 0; i < len(xs); i++ { xs[i] }
    ● No off-by-one errors (range handles bounds automatically)
    ● Works the same way for slices, arrays, maps, strings, channels
- len(xs) returns the current NUMBER OF ELEMENTS (not capacity)
- The blank identifier _ discards the index — use it when you only need values
- You CAN use C-style loops when you need more control over iteration
  (e.g., iterating backwards, stepping by 2, modifying the slice during iteration)
*/
