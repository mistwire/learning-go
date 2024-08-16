package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(401)
	fmt.Printf("The random integer is %v\n", x)
	switch {
	case x <= 100:
		fmt.Println("between 0 and 100")
	case x >= 101 && x <= 200:
		fmt.Println("between 101 and 200")
	case x <= 250:
		fmt.Println("between 201 and 250")
	default:
		fmt.Println("The number is above 250")
	}
}
