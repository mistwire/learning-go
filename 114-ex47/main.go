/*
For this exercise, do the following:
● Create a slice to store the names of all of the states in the United States of America.

	○ Use make and append to do this.
	○ Goal: do not have the array that underlies the slice created more than once.

● Print out

	○ the len
	○ the cap
	○ the values, along with their index position, without using the range clause.

● Here is a list of the 50 states:
` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, `
Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, `
Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, `
Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`,
` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`,
` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, `
Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`,
*/

package main

import "fmt"

func main() {
	xi := make([]string, 0, 50) //0 length, 10 capacity
	xi = append(xi, ` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, 
		` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, 
		` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, 
		` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, 
		` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, 
		` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`,` Pennsylvania`, ` Rhode Island`, 
		` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, 
		` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`)

	fmt.Println(len(xi))
	fmt.Println(cap(xi))
	for i:= 0; i < len(xi); i++{
		fmt.Printf("Index %v:\t Value:%v\n", i, xi[i])
	}
}

/*
REMARKS:
- make([]T, length, capacity) creates a slice with specified length and capacity
- Length: number of elements currently in the slice
- Capacity: size of the underlying array (max elements before reallocation)
- Using make with capacity=50 means the underlying array is allocated ONCE
- Without pre-allocating, append() may reallocate multiple times as slice grows
- Pre-allocating capacity improves performance when you know the final size
- len(slice) returns current length, cap(slice) returns capacity
- Traditional for loop (not range) used here to iterate by index
*/
