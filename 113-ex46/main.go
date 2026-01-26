/*
To DELETE from a slice, we use APPEND along with SLICING. For this hands-on exercise,
follow these steps:
● start with this slice

	○ x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

● use APPEND & SLICING to get these values here which you should ASSIGN to the
variable “x” and then print:

	○ [42, 43, 44, 48, 49, 50, 51]
*/
package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x[:3], x[6:]...)
	fmt.Println(x)

}

/*
REMARKS:
- Go has NO built-in delete function for slices
- Delete by combining slicing + append: append(slice[:i], slice[i+n:]...)
- x[:3] keeps elements at index 0, 1, 2 (first 3 elements)
- x[6:] keeps elements from index 6 onward (skips indices 3, 4, 5)
- The ... unpacks x[6:] so append can use its elements
- This effectively removes elements at indices 3, 4, 5 (values 45, 46, 47)
- Warning: This modifies the underlying array - other slices sharing it will see changes
*/
