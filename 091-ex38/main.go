/*
● use the "statement statement" idiom to

	○ initialize x with and random int between 0 inclusive and 5 exclusive
	○ if x is 3
		■ print "x is 3"

● run that code 100 times
● what's the benefit of using the "statement statement" idiom?
*/
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	c := 1
	for i := range 100 {
		if x := rand.Intn(5); x == 3 {
			fmt.Printf("Iteration %v\t Total count %v\t x is %v\n", i, c, x)
		c++
		}
	}
}
