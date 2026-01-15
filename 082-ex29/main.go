/*
there are two parts ot this hands on exercise
○ create a program that has a loop that prints every number from 0 to 99
○ modify the program from the previous hands on exercise to run 100 times
*/

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := range 100 {
		fmt.Printf("i is %v\n", i)
	}

	for i := range 100 {
		x := rand.Intn(10)
		y := rand.Intn(10)

		fmt.Printf("%v:\t x is %v - y is %v\t",i + 1, x, y)

		switch {
		case x < 4 && y < 4:
			fmt.Println("Both are less than 4")
		case x > 6 && y > 6:
			fmt.Println("both are greater than 6")
		case x >= 4 && x <= 6:
			fmt.Println("x is between 4 and 6")
		case y != 5:
			fmt.Println("y is not 5")
		default:
			fmt.Println("None of the previous cases were met")
		}
	}
}
