/*
● FUNCTION EXPRESSIONS — assigning an anonymous function to a variable using :=
● FUNCTIONS ARE FIRST-CLASS VALUES in Go: they can be stored in variables, passed as arguments, and returned from functions
● BUILDING ON LESSON 145: Lesson 145 showed anonymous functions invoked immediately (IIFE style);
  here we store them in variables and call them later — the function is the value
*/
package main

import "fmt"

func main() {
	// Calling a regular named function declared at package scope
	foo()

	// FUNCTION EXPRESSION: assign an anonymous function literal to variable x
	// x has type func() — a function that takes no arguments and returns nothing
	x := func() {
		fmt.Println("Anonymous func ran")
	}

	// FUNCTION EXPRESSION with a parameter: y has type func(string)
	// The variable holds the function itself — it doesn't run until called with ()
	y := func(s string) {
		fmt.Println("Anon func showing my name:", s)
	}

	// Call x and y just like any named function — parentheses invoke the stored function
	x()
	y("Chris")
}

// foo is a regular named function declared at package scope
// Unlike function expressions, named functions are visible throughout the entire package
func foo() {
	fmt.Println("Foo ran")
}

/*
REMARKS:
- Function expressions: in Go, you assign a func literal to a variable with := — the variable's type IS the function's signature
- x has type func(); y has type func(string) — the compiler infers these types from the literal
- Named functions (foo) exist at package scope; function expressions live in the block where they're declared
- Deferred execution: unlike an IIFE (Lesson 145), storing the function in a variable delays execution until ()
- BUILDING ON LESSON 145: IIFE = define AND call immediately; function expression = define now, call later
- This is the foundation for callbacks (Lesson 148): once stored in a variable, a function can be passed around
- Go differs from Python: no lambda keyword — Go uses the same func syntax for both named and anonymous functions
*/
