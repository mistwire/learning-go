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
	
	// BigBark(s) comes from indirect module dependency:
	string3 := puppy.BigBark()
	string4 := puppy.BigBarks()

	fmt.Println("Importing from indirect module:")
	fmt.Println(string3)
	fmt.Println(string4)
		
	fmt.Println("also can print this way:")
	fmt.Println(string1, string2) // notice no newline

	fmt.Println("also this way:")
	fmt.Println(puppy.Bark())
	fmt.Println(puppy.Barks())

}