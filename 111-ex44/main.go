/*
Using the code from the previous example, use SLICING to create the following new slices
which are then printed:
● [42 43 44 45 46]
● [47 48 49 50 51]
● [44 45 46 47 48]
● [43 44 45 46 47]
*/

package main

import "fmt"

func main (){
	xs := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	// hint: [inclusive:exclusive]
	fmt.Println(xs[:5])
	fmt.Println(xs[5:])
	fmt.Println(xs[2:7])
	fmt.Println(xs[1:6])

}

/*
REMARKS:
- Slicing syntax: slice[start:end] where start is INCLUSIVE, end is EXCLUSIVE
- slice[:n] - from beginning to index n (exclusive) - "first n elements"
- slice[n:] - from index n to the end - "everything after first n"
- slice[a:b] - from index a (inclusive) to index b (exclusive)
- The new slice SHARES the same underlying array as the original
- Modifying elements in a sub-slice affects the original slice!
- Length of result: end - start (e.g., xs[2:7] has length 5)
*/