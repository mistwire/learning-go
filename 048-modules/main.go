package main

import (
	"fmt"

	"github.com/mistwire/puppy"
)

func main() {
	fmt.Println("1st way to print")
	string1 := puppy.Bark()
	string2 := puppy.Barks()
	fmt.Println(string1)
	fmt.Println(string2)

	fmt.Println("also can print this way:")
	fmt.Println(string1, string2) // notice no newline

	fmt.Println("also this way:")
	fmt.Println(puppy.Bark())
	fmt.Println(puppy.Barks())

}
