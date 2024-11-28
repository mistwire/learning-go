package main

import (
	"fmt"
)

func main() {
	// the short declaration operator only works inside of a function:
	y := 42
	z := 42.0
	fmt.Printf("%v of type %T \n", y, y)
	fmt.Printf("%v of type %T \n", z, z)

	var m float32 = 43.742
	fmt.Printf("%v of type %T \n", m, m)

	/*
		// this does not work!
		// in go you can't take a VALUE that is float32 and store it
		// in a variable that is declared to hold a VALUE of float64
		z = m
		fmt.Printf("%v of type %T \n", z, z)
	*/
	// you CAN re-cast (convert) it as a float64
	// https://go.dev/ref/spec#Conversions
	z = float64(m)
	fmt.Printf("%v of type %T \n", z, z)
}
