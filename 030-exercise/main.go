package main

import "fmt"

func main() {
	var x int8 = 127
	var y uint8 = 255 

	fmt.Printf("%v is of type %T\n", x, x)
	fmt.Printf("%v is of type %T\n", y, y)
}
