/*
● EXERCISE — Numeric type limits: int8 vs uint8
● Building on Lesson 017: using "var" with explicit types to control integer size
● SIGNED integers (int8) use one bit for the sign → range: -128 to 127
● UNSIGNED integers (uint8) use all bits for value → range: 0 to 255
*/
// write a program that declares two variables
// 	● one variable to store a VALUE of TYPE int8
// 		○ assign to it the largest number possible, then print it
// 	● one variable to store a VALUE of TYPE uint8
// 		○ assign to it the largest number possible, then print it

package main

import "fmt"

func main() {
	var x int8 = 127   // max int8: 2^7 - 1 = 127 (1 bit for sign, 7 for value)
	var y uint8 = 255  // max uint8: 2^8 - 1 = 255 (all 8 bits for value)
	fmt.Printf("%v is of type %T\n", x, x)
	fmt.Printf("%v is of type %T\n", y, y)

}

/*
REMARKS:
- Go has SIZED integer types: int8, int16, int32, int64 (and unsigned: uint8, uint16, uint32, uint64)
- "int" without a size is platform-dependent (64-bit on modern systems) — used by default with :=
- SIGNED (int8): one bit stores the sign (+ or -), leaving 7 bits for the value
    range: -128 to 127 (2^7 values each direction)
- UNSIGNED (uint8): no sign bit, all 8 bits store the value
    range: 0 to 255 (2^8 - 1)
- uint8 is aliased as "byte" in Go — they are the same type
- int32 is aliased as "rune" — represents a Unicode code point
- OVERFLOW: if you exceed the max, Go wraps around (128 in int8 becomes -128)
  but the compiler catches constant overflows at compile time
- This connects to Lesson 017: "var" with explicit type is REQUIRED here because
  := would infer "int" (too large), and we specifically need int8/uint8
- These sized types matter for:
    ● Binary protocols and file formats (exact byte layouts)
    ● Memory-sensitive applications (millions of values)
    ● Interoperability with C code via cgo
*/
