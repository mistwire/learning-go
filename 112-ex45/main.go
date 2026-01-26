/*
Follow these steps:
● start with this slice

	○ x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

● append to that slice this value

	○ 52

● print out the slice
● in ONE STATEMENT append to that slice these values

	○ 53
	○ 54
	○ 55

● print out the slice
● append to the slice this slice

	○ y := []int{56, 57, 58, 59, 60}

● print out the slice
*/
package main

import "fmt"

func main (){
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	fmt.Println(x)
	x = append(x, 53, 54, 55)
	fmt.Println(x)
	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...) // ... is the unpacking operator
	fmt.Println(x)

}

/*
REMARKS:
- append() adds elements to the end of a slice and returns a NEW slice
- MUST reassign: x = append(x, ...) - append doesn't modify in place
- Can append multiple values: append(x, 1, 2, 3)
- To append another slice, use the ... (variadic/unfurl) operator: append(x, y...)
- The ... "unpacks" the slice into individual elements
- If capacity is exceeded, Go allocates a new, larger underlying array
- append() is the primary way to grow slices in Go
*/