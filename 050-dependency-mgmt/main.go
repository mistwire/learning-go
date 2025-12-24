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

	// now the big dog stuff :-) 
	s3 := puppy.BigBark()
	s4 := pup.BigBark
	fmt.Println(s3)

	// or like this:
	fmt.Println(puppy.Bark())
	fmt.Println(puppy.Barks())
}
