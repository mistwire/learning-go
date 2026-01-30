/*
● EXERCISE — Print verbs %v and %T
● Building on Lesson 017: using format verbs to inspect VALUES and TYPES
● Short declaration (:=) infers: "Yogurt" → string, 123 → int, 123.123 → float64
*/
//	write a program that uses the following:
//
// ● for a VARIABLE storing a VALUE of TYPE
// ○ string
// ○ int
// ○ float64
// ● use print verbs to show
// ○ value
// ○ type
package main

import "fmt"

func main() {
	// multiple assignment — Go infers each type from the literal
	x, y, z := "Yogurt", 123, 123.123
	fmt.Printf("%v is of type %T\n", x, x) // Yogurt — string
	fmt.Printf("%v is of type %T\n", y, y) // 123 — int
	fmt.Printf("%v is of type %T\n", z, z) // 123.123 — float64
}

/*
REMARKS:
- %v is the DEFAULT format verb — works for any type (value)
- %T prints the TYPE of the value — extremely useful for debugging
- Go's default type inference for literals:
    ● "text" → string
    ● 123 → int (platform-sized: 64-bit on modern systems)
    ● 123.123 → float64 (NOT float32)
    ● true → bool
- Other useful format verbs (building toward Lesson 029):
    %d = decimal integer, %b = binary, %#x = hexadecimal with 0x prefix
    %s = string, %f = floating point, %e = scientific notation
- fmt.Printf does NOT add a newline — you must include \n explicitly
  (unlike fmt.Println which adds one automatically)
*/
