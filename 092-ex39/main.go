/*
What do these print:
● fmt.Println(true && true)
● fmt.Println(true && false)
● fmt.Println(true || true)
● fmt.Println(true || false)
● fmt.Println(!true)
*/
package main

import "fmt"

func main() {
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
}
