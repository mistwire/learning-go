/*
● CUSTOM TYPES with IOTA — practical example using byte sizes
● Building on Lesson 025: combining iota with a CUSTOM TYPE and bit shifting
● "type ByteSize int" creates a NEW NAMED TYPE with int as the underlying type
● The expression 1 << (10 * iota) is implicitly repeated for each constant
● Each step shifts 10 more bits — doubling by 1024 each time (KB → MB → GB → ...)
*/
package main

import "fmt"

// "type" keyword creates a new named type — ByteSize is distinct from int
// even though its underlying type IS int (you'd need conversion to mix them)
type ByteSize int

const (
	_           = iota             // iota=0, discarded
	KB ByteSize = 1 << (10 * iota) // iota=1: 1 << 10 = 1024
	MB                              // iota=2: 1 << 20 = 1,048,576
	GB                              // iota=3: 1 << 30 = 1,073,741,824
	TB                              // iota=4: 1 << 40
	PB                              // iota=5: 1 << 50
	EB                              // iota=6: 1 << 60
	// ZB                           // would overflow int64
	// YB
)

func main() { // go formatting verbs: https://pkg.go.dev/fmt
	fmt.Printf("%d \t\t\t %b\n", KB, KB)
	fmt.Printf("%d \t\t %b\n", MB, MB)
	fmt.Printf("%d \t\t %b\n", GB, GB)
	fmt.Printf("%d \t\t %b\n", TB, TB)
	fmt.Printf("%d \t %b\n", PB, PB)
	fmt.Printf("%d \t %b\n", EB, EB)
}

/*
REMARKS:
- KEY DIFFERENCE FROM LESSON 025: the constants now have a NAMED TYPE (ByteSize) instead of being untyped
- "type ByteSize int" is the same "type" keyword used later for structs (Lessons 125-127)
  — here it creates a simple alias-like type, there it creates a composite struct type
- The expression 1 << (10 * iota) is set on the FIRST constant and implicitly repeats
  with iota incrementing: 1<<10 = 1024 (KB), 1<<20 = ~1M (MB), 1<<30 = ~1B (GB)
- ZB/YB are commented out because they would OVERFLOW int64 (max ~9.2 × 10^18)
  EB (1<<60 = ~1.15 × 10^18) is the largest that fits
- Named types provide TYPE SAFETY — you can't accidentally assign a plain int to a ByteSize
  without explicit conversion
- You can attach METHODS to named types (covered later with structs)
  — e.g., func (b ByteSize) String() string { ... } for pretty printing
- This pattern (custom type + iota) is the idiomatic way to create enumerations in Go
*/
