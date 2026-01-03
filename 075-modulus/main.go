package main

import "fmt"

func main() {
	x := 83 / 40
	y := 83 % 40
	fmt.Printf("83 / 40 = %v\n83 %% 40 = %v\n", x, y)

	for i := range 100 {
		if i%2 == 0 {
			fmt.Printf("i is even: %v\n", i)
		}
	}
}
