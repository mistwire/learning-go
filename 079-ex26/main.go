// LESSON: Exercise - using init() function for initialization
/*
Modify the previous program to have your program use the init func to print
○ "This is where initialization for my program occurs"
● read about the init function in effective go
○ What does niladic mean?
*/
package main

import (
	"fmt"
	"math/rand"
)

func init (){
	fmt.Println("This is where initialization for my program occurs")
}

func main() {
	x := rand.Intn(250)
	x = 0
	fmt.Printf("The variable x is: %v\n", x)
	if x <= 100 {
		fmt.Println("between 0 and 100")
	} else if x <= 200 {
		fmt.Println("between 101 and 200")
	} else if x <= 250 {
		fmt.Println("between 201 and 250")
	}
}
