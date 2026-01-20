/*
● create a loop that runs 5 times
● nest a loop within the first loop that also prints 5 times
● print something in each loop to illustrate what is occuring
*/
package main

import "fmt"

func main() {
	for i := range 5 {
		for j := range 5 {
			fmt.Printf("Outer loop %v\t Inner loop %v\n", i, j)
		}
		fmt.Println("---------")
	}
}
