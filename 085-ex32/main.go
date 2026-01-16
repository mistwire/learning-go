/*
● create a for loop that uses break statement
● increment or decrement the variable being checked as a condition in the loop
*/
package main

import "fmt"

func main(){
	i := 0
	for {
		fmt.Printf("i is %v\n", i)
		if i == 20{
			break
		}
		i++
	}
}