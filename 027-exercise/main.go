/*
Write a program that uses the following:
- for a var storing a value of type:
	- string
	- int
	- float64
- use print verbs (https://pkg.go.dev/fmt) to show:
	- value
	- type
*/


package main

import "fmt"

func main() {
	my_int := 42
	my_string := "Chris Williams"
	my_float64 := 42.00001
	
	fmt.Printf("%v is of type %T\n", my_string, my_string)
	fmt.Printf("%v is of type %T\n", my_int, my_int)
	fmt.Printf("%v is of type %T\n", my_float64, my_float64)
}