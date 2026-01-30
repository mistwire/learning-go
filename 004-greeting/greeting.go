/*
● USER INPUT with fmt.Scanln
● Declaring variables with "var" to hold user-provided values
● Using POINTERS (&) to pass variable addresses so Scanln can write into them
● Basic ERROR HANDLING with the if/else pattern
● Multiple return values: Scanln returns both a count and an error
*/
package main

import (
	"fmt"
)

func main() {
	// "var" declares variables with their ZERO VALUES ("" for string, 0 for int)
	var name string
	var age int

	fmt.Print("Enter your name and age: ")
	// fmt.Scanln reads space-separated input from stdin
	// &name, &age pass the ADDRESSES (pointers) so Scanln can modify the variables
	// returns: n (number of items scanned), err (any error encountered)
	n, err := fmt.Scanln(&name, &age)

	// basic error handling pattern: check if err is not nil
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Scanned %d items. Name: %s, Age: %d\n", n, name, age)
	}
}

/*
REMARKS:
- fmt.Scanln reads from standard input, stopping at a newline
- The & operator takes the ADDRESS of a variable — this is how Go passes "by reference"
  (Go is always pass-by-value; & gives you a pointer value to pass)
- Multiple return values (n, err) are idiomatic Go — most I/O functions return (result, error)
- The "if err != nil" pattern is THE standard way to handle errors in Go
  (there is no try/catch/exception mechanism)
- %s is the format verb for strings, %d for decimal integers
- var is used here because we need zero-value variables BEFORE we have values to assign
  (short declaration := wouldn't make sense since we don't have initial values)
*/
