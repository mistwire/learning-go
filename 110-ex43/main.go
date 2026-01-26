/*
● Using a COMPOSITE LITERAL:
● create a SLICE of TYPE int
● assign these 10 VALUES
42, 43, 44, 45, 46, 47, 48, 49, 50, 51
● Range over the slice and print the values out.

	○ print out the VALUE and the TYPE
*/
package main

import "fmt"

func main (){
	xs := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	for i := range(xs){
		fmt.Printf("VALUE: %v \t TYPE:%T\n", i, i)
	}
}

/*
REMARKS:
- Slices are DYNAMIC and can grow/shrink - no fixed size in the type declaration
- []int (no size) = slice, [5]int (with size) = array
- Slices are REFERENCE types - they point to an underlying array
- Using only one variable with range (i := range xs) gives you just the INDEX
- To get both index and value, use: for i, v := range xs
- Note: This code prints the index twice, not the actual slice values!
- %T format verb prints the TYPE of the value
*/