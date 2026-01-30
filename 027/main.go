/*
● EXERCISE — Variable declarations review
● Practicing all the declaration forms from Lesson 013:
  var zero value, short declaration (:=), multiple init, explicit type, blank identifier (_)
*/
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

	// var for zero value — int zero value is 0
	var a int
	fmt.Println(a)

	// short declaration operator — type inferred as int
	b := 69 // nice
	fmt.Println(b)

	// multiple initializations — mixed types in one statement
	c, d := 69, "this is a string :-)"
	fmt.Println(c, d)

	// var when specificity is required — need float32, not the default float64
	var e float32 = 69.69
	fmt.Printf("%v is of type %T \n", e, e)

	// blank identifier — discard the third value
	f, g, _ := 45, 46, 47
	fmt.Println(f, g)
}

/*
REMARKS:
- This exercise reinforces the THREE declaration styles from Lesson 013
- WHEN TO USE EACH:
    ● := — when you have an initial value and the default type is fine (most common)
    ● var x Type — when you need a zero-value variable before you have a value
    ● var x Type = value — when you need a SPECIFIC type that differs from the default
      (e.g., float32 instead of float64, int8 instead of int)
- Multiple assignment works with := and with var (var a, b int = 1, 2)
- The blank identifier _ prevents "declared and not used" compile errors
- fmt.Printf with %v (value) and %T (type) is the go-to for inspecting variables
*/
