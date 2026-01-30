/*
● IOTA — auto-incrementing constant generator for enumerations
● iota starts at 0 and increments by 1 for each constant in a const block
● Use _ (blank identifier from Lesson 013) to discard the 0 value if you don't want it
● BIT SHIFTING with << creates powers of 2 — useful for flags and bitmasks
*/
package main

import "fmt"

const (
	_ = iota // iota = 0, discarded with _   https://go.dev/wiki/Iota
	a        // a = 1 (iota auto-increments)
	b        // b = 2
	c        // c = 3
	d        // d = 4
	e        // e = 5
	f        // f = 6
)

func main() { // go formatting verbs: https://pkg.go.dev/fmt
	// << is the LEFT SHIFT operator: 1<<n shifts the binary 1 left by n positions
	// 1<<1 = 2, 1<<2 = 4, 1<<3 = 8, ... (powers of 2)
	// %d = decimal, %b = binary representation
	fmt.Printf("%d \t %b\n", 1, 1)       // 1       1
	fmt.Printf("%d \t %b\n", 1<<a, 1<<a) // 2      10
	fmt.Printf("%d \t %b\n", 1<<b, 1<<b) // 4     100
	fmt.Printf("%d \t %b\n", 1<<c, 1<<c) // 8    1000
	fmt.Printf("%d \t %b\n", 1<<d, 1<<d) // 16  10000
	fmt.Printf("%d \t %b\n", 1<<e, 1<<e) // 32 100000
	fmt.Printf("%d \t %b\n", 1<<f, 1<<f) // 64 1000000
}

/*
REMARKS:
- iota is a PREDECLARED IDENTIFIER that represents successive untyped integer constants
- It resets to 0 at each new const block and increments by 1 for each const spec
- The expression applied to the first constant is IMPLICITLY REPEATED for subsequent ones
  e.g., if you write KB = 1 << (10 * iota), MB, GB all get the same expression with increasing iota
- BIT SHIFTING (<<) combined with iota is a classic pattern for:
    ● Permission flags (read=1, write=2, execute=4)
    ● Byte sizes (KB, MB, GB — see Lesson 026)
    ● Feature toggles that can be combined with bitwise OR (|)
- Constants in Go are UNTYPED by default — they have higher precision than any concrete type
  and adapt to the context where they're used
- Constants are set at COMPILE TIME, not runtime — they cannot be changed
- iota is Go's elegant alternative to manually numbering enums like in C/Java
*/
