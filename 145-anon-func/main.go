/*
● ANONYMOUS FUNCTIONS — functions declared inline with no name, also called "function literals"
● IMMEDIATELY INVOKED — an anonymous function can be defined and called in the same expression
● PARAMETERS & ARGUMENTS — anon funcs support the full function signature: parameters and return values
● BUILDING ON LESSON 133-134: extends named function syntax; anon funcs follow the same rules but skip the identifier
*/
package main

import "fmt"

func main() {
	foo() // calling a named function defined below — standard function call

	// Anonymous function with no parameters, immediately invoked
	// Syntax: func(){ ... }()
	// The trailing () after the closing brace is the call; without it, the function is declared but never runs
	func() {
		fmt.Println("Anonymous func ran")
	}()

	// Anonymous function with a parameter, immediately invoked with an argument
	// The argument "Chris" is passed at the call site, inside the trailing ()
	// This is the same mechanism as a named function — just compressed into one expression
	func(s string) {
		fmt.Println("Anon func showing my name:", s)
	}("Chris")
}

// named func
// func (r receiver) identifier(p parameter(s)) (return(s)) { code }

// anon func — same structure, but NO identifier (name) between "func" and the parameter list
// func(parameters go here)(returns){code here}(arguments here)
func foo() {
	fmt.Println("Foo ran")
}

/*
REMARKS:
- FUNCTION LITERALS: anonymous functions are first-class values in Go — they can be assigned to variables, passed as arguments, or returned from functions
- IIFE PATTERN: immediately invoked function expressions are common in Go, especially with goroutines: go func(){ ... }()
- CLOSURES PREVIEW: anonymous functions can capture and use variables from their surrounding scope — this is called a CLOSURE (upcoming lesson)
- vs PYTHON: Python lambdas are limited to a single expression; Go anonymous functions are full functions with any number of statements
- BUILDING ON LESSON 133-134: the parameter/argument and return rules are identical to named functions — only the name is missing
- Common pitfall: forgetting the trailing () — a func literal without () is a value, not a call; the function won't execute
*/
