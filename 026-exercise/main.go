/*
Write a program that uses the following:
- var for zero value
- short declaration operator
- multiple initializations
- var when specificity is required
- blank identifier
*/
package main

import "fmt"

func main() {

	// var for zero value
	var a int
	fmt.Println(a)

	// short declaration operator
	b := 69 // nice
	fmt.Println(b)

	// multiple initializations
	c, d := 69, "this is a string :-)"
	fmt.Println(c, d)

	// var when specificity is required
	var e float32 = 69.69
	fmt.Printf("%v is of type %T \n", e, e)

	// blank identifier
	f, g, _ := 45, 46, 47
	fmt.Println(f, g)
}
