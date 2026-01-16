/*
● Create one random int between 0 inclusive and 5 exclusive
	○ assign the value to a variable with the identifier x
● Use a switch statement to print a description of the variable and value
● run the code 42 times and print the iteration number
*/

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := range(42) {
		x := rand.Intn(5)
		fmt.Printf("%v iteration\t", i)
		switch x {
		case 0:
			fmt.Println("x is zero")
		case 1:
			fmt.Println("x is one")
		case 2:
			fmt.Println("x is two")
		case 3:
			fmt.Println("x is three")
		case 4:
			fmt.Println("x is four")
		}
	}

}
