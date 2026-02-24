/*
● CALLBACKS — passing a function as an argument to another function
● FUNCTION TYPE AS PARAMETER — func(int, int) int is a valid, concrete parameter type
● HIGHER-ORDER FUNCTIONS — doMath accepts a function and calls it; the caller decides the behavior
● BUILDING ON LESSONS 146-147: 146 stored functions in variables; 147 returned them; here we pass them as arguments
  — all three are expressions of functions being first-class values in Go
*/

// a callback is passing a func into a function as an argument
package main

import "fmt"

func main() {
	// Print the types so we can see exactly what we're working with:
	// doMath type: func(int, int, func(int, int) int) int
	// add type:    func(int, int) int
	fmt.Printf("doMath type: %T\n", doMath)
	fmt.Printf("add type: %T\n", add)

	// Pass add as a CALLBACK — doMath will delegate the operation to add(40, 2)
	x := doMath(40, 2, add)
	fmt.Println(x) // 42

	// Pass subtract as a different CALLBACK — same doMath function, entirely different behavior
	// This is POLYMORPHISM via first-class functions, no interfaces or structs needed
	y := doMath(47, 5, subtract)
	fmt.Printf("subtract type: %T\n", subtract)
	fmt.Println(y) // 42
}

// doMath is a HIGHER-ORDER FUNCTION — its third parameter f has type func(int, int) int
// Any function matching that exact signature can be passed in as the callback
// doMath itself doesn't know or care whether f adds, subtracts, multiplies, etc.
func doMath(a int, b int, f func(int, int) int) int {
	return f(a, b) // delegate the actual computation to the callback
}

// add satisfies the func(int, int) int type — eligible to be passed as a callback
func add(a int, b int) int {
	return a + b
}

// subtract also satisfies func(int, int) int — polymorphic behavior via callbacks
func subtract(a int, b int) int {
	return a - b
}

/*
REMARKS:
- Callbacks: a callback is simply a function passed as an argument — Go supports this natively via first-class functions
- Function type as parameter: func(int, int) int is a fully declared type — any matching function can fill that slot
- POLYMORPHISM without interfaces: swapping the callback changes doMath's behavior entirely — no subclassing needed
- BUILDING ON LESSONS 146-147: storing (146), returning (147), and passing (148) functions are all the same idea at its core
- Real-world use: Go's sort.Slice uses exactly this pattern: sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
- Signature must match EXACTLY: func(int, int) int will NOT accept func(int, int) float64 — Go is strictly typed
- Go philosophy: callbacks enable flexible, composable code without the overhead of class hierarchies or inheritance
- Common pitfall: forget the () and you pass the function itself — include () to call it and pass the result instead
*/
