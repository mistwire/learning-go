/*
● RETURNING FUNCTIONS — a function can return another function as its value
● FUNCTION TYPES — a function's signature (e.g. func() int) is a complete type in Go
● HIGHER-ORDER FUNCTIONS — bar() is a higher-order function because it returns a function
● BUILDING ON LESSON 146: just as we stored functions in variables, a function can produce and hand back a new function
*/
package main

import "fmt"

func main() {
	// foo() returns an int — x holds the plain integer value 42
	x := foo()
	fmt.Println(x) // 42

	// bar() returns a func() int — y holds the *function itself*, NOT a computed result
	// bar() runs and hands us a brand new function; that function is stored in y
	y := bar()
	fmt.Println(y()) // call y() to invoke the returned function — prints: 43

	// %T reveals the type of each identifier:
	// foo  → func() int             (a function that returns int)
	// bar  → func() func() int      (a function that returns a function that returns int)
	// y    → func() int             (the function bar returned, now stored in y)
	fmt.Printf("%T\n", foo)
	fmt.Printf("%T\n", bar)
	fmt.Printf("%T\n", y)
}

// foo returns a plain int — a standard single-return function (see Lesson 133)
func foo() int {
	return 42
}

// bar's return type is func() int — the entire function signature is the return type
// Reading the signature: "bar is a function that returns a function that returns int"
func bar() func() int {
	// Return an anonymous function literal — the caller receives this function as a value
	return func() int {
		return 43
	}
}

/*
REMARKS:
- Returning functions: bar()'s return type is func() int — declare the full signature between func and the opening brace
- Key distinction: x := foo() stores 42 (a value); y := bar() stores a function (not 43); y() then gives 43
- HIGHER-ORDER FUNCTION: any function that takes OR returns another function earns this label
- Reading nested types: func() func() int — "a function with no params that returns (a function with no params that returns int)"
- Use fmt.Printf("%T") liberally when learning first-class functions — seeing the types makes the concept concrete
- BUILDING ON LESSON 146: 146 stored an external function expression; here the function is *created inside* bar and returned
- Returned functions are the foundation for CLOSURES: the returned func can close over variables in bar's local scope
- Go vs Python: Go requires the return type to be spelled out explicitly — func() int — whereas Python infers it
*/
