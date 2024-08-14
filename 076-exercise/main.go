package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(251)
	fmt.Printf("The random integer is %v\n", x)
	if x <= 100 {
		fmt.Println("between 0 and 100")
	} else if x >= 101 && x <= 200 {
		fmt.Println("between 101 and 200")
	} else if x <= 250 {
		fmt.Println("between 201 and 250")
	}

	y := rand.Intn(4)
	fmt.Printf("The random integer is %v\n", y)
}
