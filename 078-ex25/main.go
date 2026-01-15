/*
● Modify the previous program to use one of these conditional logic statements

	○ a switch statement
	○ a select statement

● Which of the above conditional logic statements did you choose and why?
*/
package main

import (
	"fmt"
	"math/rand"
)

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
