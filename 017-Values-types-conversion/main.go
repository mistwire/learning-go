/*
● TYPE CONVERSION (not "casting" — Go calls it conversion)
● Building on Lesson 013: Go is STATICALLY and STRONGLY typed — no implicit conversions
● You MUST explicitly convert between types, even between float32 and float64
● Syntax: targetType(value) — e.g., float64(m)
● Short declaration (:=) infers types: 42 → int, 42.0 → float64
*/
package main

import (
	"fmt"
)

func main() {
	// := infers the type from the literal value
	y := 42   // inferred as int
	z := 42.0 // inferred as float64 (Go defaults floating-point literals to float64)
	fmt.Printf("%v of type %T \n", y, y)
	fmt.Printf("%v of type %T \n", z, z)

	// var with explicit type — float32 instead of the default float64
	var m float32 = 43.742
	fmt.Printf("%v of type %T \n", m, m)

	/*
		// this does NOT work — Go won't implicitly convert float32 → float64
		// even though float64 can hold everything float32 can
		z = m
		fmt.Printf("%v of type %T \n", z, z)
	*/
	// EXPLICIT CONVERSION: float64(m) converts the float32 value to float64
	// https://go.dev/ref/spec#Conversions
	z = float64(m)
	fmt.Printf("%v of type %T \n", z, z)
}

/*
REMARKS:
- Go is STRONGLY TYPED — no implicit type conversions, ever
  (unlike C/JavaScript where int can silently become float)
- Conversion syntax: Type(value) — looks like a function call but it's a type conversion
- Even "compatible" types like float32 → float64 require explicit conversion
- int and float64 are different types — you can't mix them in arithmetic without converting
- This strictness catches bugs at compile time rather than runtime
- Common conversions: int() ↔ float64(), string() ↔ []byte(), int() ↔ int64()
- Note: converting float → int TRUNCATES (drops the decimal), it does NOT round
- %v prints the value, %T prints the type — useful for debugging type issues
- BUILDING ON LESSON 013: "var" with explicit type (var m float32) is needed when
  the default inferred type isn't what you want
*/
