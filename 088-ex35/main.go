/*
● below is the code to create a data structure called a slice of ints
● put this code into a program
xi := []int{42, 43, 44, 45, 46, 47}
● use a for range loop to print each value and the index position of each value
😃🌴🌞🔆🌻🙌🌴🌴🌴🌴🌴🌴🌴🌴
*/
package main

import "fmt"

func main() {
	xi := []int{42, 43, 44, 45, 46, 47}

	for i, v := range xi {
		fmt.Printf("For index %v: the value is %v\n", i, v)
	}
}
