/*
● VARIABLES in Go
● Three ways to declare variables: short declaration (:=), var with zero value, var with type+value
● MULTIPLE ASSIGNMENT: assign several variables in one statement
● BLANK IDENTIFIER (_): discard values you don't need
● All declared variables MUST be used — Go enforces this at compile time
*/
package main

import "fmt"

func main() {
	// SHORT DECLARATION OPERATOR (:=)
	// declares AND assigns in one step — Go infers the type
	// can ONLY be used inside a function (not at package level)
	a := 42
	fmt.Println(a)

	// MULTIPLE ASSIGNMENT with mixed types
	// _ (blank identifier) discards the value at that position
	b, c, d, _, f := 0, 1, 2, 3, "happiness"
	fmt.Println(b, c, d, f)

	// this would not work — "e" would be declared but never used (compile error)
	/*
		b, c, d, e := 0, 1, 2, 3
		fmt.Println(b, c, d)
	*/

	// VAR with zero value — declares the variable, Go assigns the zero value for the type
	// zero values: 0 for int, "" for string, false for bool, nil for pointers/slices/maps
	var g int
	fmt.Println(g) // prints 0 (zero value)
	g = 43         // assignment operator (=) — variable already exists
	fmt.Println(g)

	// VAR with explicit type AND value — use when you need a specific type
	var h int = 44
	fmt.Println(h)

}

/*
REMARKS:
- Go has THREE ways to declare variables:
    1. := (short declaration) — most common inside functions, type is inferred
    2. var x int — declares with zero value, useful when you don't have an initial value yet
    3. var x int = 42 — explicit type, useful when you need specificity (e.g., var x float32 = 3.14)
- := can ONLY be used inside functions; at package level you must use "var"
- Every declared variable MUST be used — unused variables cause a COMPILE ERROR
  (this is unique to Go and prevents dead code)
- The blank identifier _ lets you discard values (e.g., from multi-return functions)
- = is assignment (variable must already exist); := is declaration + assignment (creates new variable)
- Zero values ensure variables are always initialized — no "undefined" or garbage values in Go
*/
