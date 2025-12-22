// write a program that declares two variables
// 	● one variable to store a VALUE of TYPE int8
// 		○ assign to it the largest number possible, then print it
// 	● one variable to store a VALUE of TYPE uint8
// 		○ assign to it the largest number possible, then print it

package main

import "fmt"

func main() {
	var x int8 = 127
	var y uint8 = 255
	fmt.Printf("%v is of type %T\n", x, x)
	fmt.Printf("%v is of type %T\n", y, y)

}
