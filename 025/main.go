package main

import "fmt"

const (
	_ = iota // _ == 0
	a        // a == 1
	b        // b == 2
	c        // c == 3
	d        // d == 4
	e        // e == 5
	f        // f == 6
)

func main() { // go formatting verbs: https://pkg.go.dev/fmt
	fmt.Printf("%d \t %b\n", 1, 1)
	fmt.Printf("%d \t %b\n", 1<<a, 1<<a)
	fmt.Printf("%d \t %b\n", 1<<b, 1<<b)
	fmt.Printf("%d \t %b\n", 1<<c, 1<<c)
	fmt.Printf("%d \t %b\n", 1<<d, 1<<d)
	fmt.Printf("%d \t %b\n", 1<<e, 1<<e)
	fmt.Printf("%d \t %b\n", 1<<f, 1<<f)
}
