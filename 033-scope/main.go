/*
● SCOPE — where variables, types, and functions are visible
● THREE levels of scope in Go: universe, package, block
● VARIABLE SHADOWING: a local variable can shadow an outer one with the same name
● First look at STRUCTS and METHODS (expanded in Lessons 125-127)
● EXPORTED identifiers: uppercase names are visible outside the package
*/
/*
THIS IS ONLY SOME OF THE CODE SHOWN IN THE LECTURE
ALL OF THE CODE SHOWN IN THE LECTURE IS AVAILABLE ON
https://github.com/GoesToEleven/learn-to-code-go-version-03
look in the "031-scope" directory
*/

package main

// imported packages are in "file block" scope — visible throughout this file
import (
	"fmt"
	"learning-go/033-scope/furtherexploredgo"
)

// x is in "package block" scope — visible to ALL functions in this package
var x = 42

func main() {
	// x can be accessed here — it's in package scope
	fmt.Println(x)

	// printMe() also accesses the same package-level x
	printMe()

	// y does not exist here yet — it hasn't been declared
	// fmt.Println(y) // would not compile

	// y is in "block scope" — only visible inside this function
	y := 43
	fmt.Println(y)

	// struct and method preview (detailed in Lessons 125-127)
	p1 := person{
		"James",
	}
	p1.sayHello()

	// VARIABLE SHADOWING: this x is a NEW variable that shadows the package-level x
	// the := creates a new local x; the package-level x (42) is unchanged
	x := 32
	fmt.Println("Variable Shadowed x:", x) // prints 32 (local)

	// package-level x is still 42 — shadowing doesn't modify the outer variable
	printMe() // still prints 42

	// Fasincating is EXPORTED (uppercase F) — visible outside its package
	furtherexploredgo.Fasincating()

	// statement-before-condition: z is scoped to the if/else block only
	if z := 82; z > 50 {
		fmt.Println(z)
	}
	// z is NOT accessible here — its scope ended with the closing }
	// fmt.Println(z) // would not compile
	/*
	take a look at the "comma ok idiom"
	https://go.dev/doc/effective_go#maps
	*/
}

func printMe() {
	// accesses the PACKAGE-LEVEL x, not the shadowed local x in main()
	fmt.Printf("Package block scope x: %d\n", x)
}

// type person is in "package block" scope — visible to all functions in this file
type person struct {
	first string
}

// METHOD: a function with a RECEIVER — attached to values of type person
// (p person) is the receiver — "p" is the variable name, "person" is the type
func (p person) sayHello() {
	fmt.Println("Hi, my name is", p.first)
}

/*
REMARKS:
- Go has THREE scope levels:
    1. UNIVERSE BLOCK — built-in types (int, string, bool) and functions (len, cap, make, append)
    2. PACKAGE BLOCK — var, type, func declared at the top level of a file (outside any function)
    3. BLOCK SCOPE — variables declared inside { } (functions, if, for, switch blocks)
- SHADOWING: a variable in an inner scope with the same name HIDES the outer one
  — the outer variable still exists and is unchanged
  — this is a common source of bugs; linters can warn about it
- EXPORTED vs UNEXPORTED (ties into Lesson 049):
    ● Uppercase first letter = exported (visible outside the package): Fasincating()
    ● Lowercase first letter = unexported (package-private): person, printMe
- Statement-before-condition (if z := 82; z > 50) limits z's scope to the if/else block
  — this pattern is used heavily with the comma-ok idiom (Lesson 069)
- METHODS preview: func (p person) sayHello() attaches a function to a type
  — this is how Go achieves object-oriented behavior without classes
  — covered in depth with structs (Lessons 125-127)
- imports are FILE-scoped — each .go file must import what it uses
*/