/*
● Using a COMPOSITE LITERAL:
● create an ARRAY which holds 5 VALUES of TYPE int
● assign VALUES to each index position.
● Range over the array and print the values out.

	○ print out the VALUE and the TYPE
*/
package main

import "fmt"

func main (){
	my_array := [...]int{13, 42, 69, 73, 1492}

	for i, v := range my_array {
		fmt.Printf("For index %v: the value is %v\n", i, v)
	}
}

/*
REMARKS:
- Arrays in Go have a FIXED size that is part of their type ([5]int is different from [6]int)
- The [...] syntax lets the compiler count the elements and set the array size automatically
- Composite literal: declaring and initializing in one statement using {values}
- Arrays are VALUE types - assigning or passing copies the entire array
- Use range to iterate: "i" gets the index, "v" gets a copy of the value at that index
- Arrays are rarely used directly in Go; slices are preferred for flexibility
*/