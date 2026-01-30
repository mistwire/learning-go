/*
● DELETE FROM SLICE — Go has no built-in "delete" for slices
● Building on Lessons 101 & 102: combines APPEND + SLICING to remove elements
● Pattern: slice = append(slice[:i], slice[i+1:]...)
● The ... operator UNPACKS the second slice into individual arguments for append
*/
package main

import "fmt"

func main() {
	xi := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("xi - %#v\n", xi)
	fmt.Println("-------------")

	// DELETE element at index 4 (value '4'):
	// xi[:4]   → [0, 1, 2, 3]         (everything BEFORE index 4)
	// xi[5:]   → [5, 6, 7, 8, 9]      (everything AFTER index 4)
	// xi[5:]...→ unpacks into 5, 6, 7, 8, 9 (variadic args for append)
	// result   → [0, 1, 2, 3, 5, 6, 7, 8, 9]
	xi = append(xi[:4], xi[5:]...)
	fmt.Printf("xi - %#v\n", xi)
	fmt.Println("-------------")
}

/*
REMARKS:
- Go has NO built-in "delete" or "remove" function for slices
  (maps DO have delete() — Lesson 118 — but slices don't)
- DELETE PATTERN: slice = append(slice[:i], slice[i+1:]...)
    1. slice[:i]   — everything before the element to delete (Lesson 102)
    2. slice[i+1:] — everything after the element to delete (Lesson 102)
    3. ...          — unpack the second slice into individual values for append (Lesson 101)
    4. append joins them together, skipping the deleted element
- The ... (variadic/spread) operator is REQUIRED because append takes individual values,
  not a slice — xi[5:]... unpacks [5,6,7,8,9] into 5, 6, 7, 8, 9
- WARNING: this modifies the UNDERLYING ARRAY — elements after the deletion point shift left
  — other slices sharing the same underlying array may be affected (Lesson 106)
- This does NOT preserve order — for order-preserving delete, this pattern is correct
  For unordered delete (faster), swap with last element and shrink: xs[i] = xs[len-1]; xs = xs[:len-1]
- BUILDING ON LESSONS 101 + 102: this is the combination of append + slicing syntax
- Go 1.21+ added slices.Delete() in the standard library: slices.Delete(xi, 4, 5)
*/
