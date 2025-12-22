package main

import (
	"fmt"
)

func main() {
	adams := 42
	fmt.Printf("42 as binary, %b \n", adams)
	fmt.Printf("42 as hexadecimal, %x \n", adams)

	// print these values as both binary & hexadecimal
	a, b, c, d, e, f, _ := 0, 1, 2, 3, 4, 5, 6
	fmt.Printf("%b %x \n", a, a)
	fmt.Printf("%b %x \n", b, b)
	fmt.Printf("%b %x \n", c, c)
	fmt.Printf("%b %x \n", d, d)
	fmt.Printf("%b %x \n", e, e)
	fmt.Printf("%b %x \n", f, f)
	// _ is a throwaway & doesn't bork the compiler
}
