//	write a program that uses the following:
//
// ● for a VARIABLE storing a VALUE of TYPE
// ○ string
// ○ int
// ○ float64
// ● use print verbs to show
// ○ value
// ○ type
package main

import "fmt"

func main() {
	x, y, z := "Yurt", 123, 123.123
	fmt.Printf("%v is of type %T\n", x, x)
	fmt.Printf("%v is of type %T\n", y, y)
	fmt.Printf("%v is of type %T\n", z, z)
}
