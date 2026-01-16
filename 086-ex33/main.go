/*
● use modulus and the continue statement in a loop to print all ODD numbers
● joke about the programmer and infinite loops
*/

package main

import "fmt"

func main() {
	for x := 0; x < 100; x++ {
		if x%2 == 0 {
			continue
		}
		fmt.Println(x)
	}
}
