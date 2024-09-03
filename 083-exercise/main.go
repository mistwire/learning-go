package main

import "fmt"

func main() {
	x := 10
	for x >= 0 {
		fmt.Printf("For this loop, x = %v\n", x)
		x--
	}
}
