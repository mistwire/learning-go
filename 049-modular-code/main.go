package main

import (
	"fmt"

	"github.com/mistwire/puppy"
)

func main() {
	s1 := puppy.Bark()
	s2 := puppy.Barks()
	fmt.Println(s1)
	fmt.Println(s2)

	// or like this:
	fmt.Println(puppy.Bark())
	fmt.Println(puppy.Barks())
}
