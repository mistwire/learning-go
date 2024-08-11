package main

import (
	"fmt"

	"github.com/mistwire/puppy"
)

func main() {
	fmt.Println("1st way to print")
	string1 := puppy.Bark()
	string2 := puppy.Barks()
	
	puppy.From11()
	fmt.Println(string1)
	fmt.Println(string2)
	
	// BigBark(s) comes from indirect module dependency:
	puppy.From12()
	string3 := puppy.BigBark()
	string4 := puppy.BigBarks()

	fmt.Println("Importing from indirect module:")
	fmt.Println(string3)
	fmt.Println(string4)
		
	fmt.Println("also can print this way:")
	fmt.Println(string1, string2) // notice no newline

	puppy.From13()
	fmt.Println("Printing a call to the func, no var:")
	fmt.Println(puppy.Bark())
	fmt.Println(puppy.Barks())

}