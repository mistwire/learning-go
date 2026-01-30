/*
● MODULUS — the % operator returns the REMAINDER of integer division
● Building on Lesson 072: commonly used inside for loops for filtering (even/odd, every Nth)
● Integer division (/) truncates: 83/40 = 2 (not 2.075)
● Modulus (%): 83%40 = 3 (the remainder after dividing 83 by 40)
● Uses "for i := range 100" — Go 1.22+ range-over-integer syntax
*/
package main

import "fmt"

func main() {
	x := 83 / 40 // integer division — truncates to 2 (drops the .075)
	y := 83 % 40 // modulus — the remainder is 3 (83 = 40*2 + 3)
	// %% is how you print a literal % in Printf (escape the format verb)
	fmt.Printf("83 / 40 = %v\n83 %% 40 = %v\n", x, y)

	// range over an integer (Go 1.22+): iterates from 0 to 99
	// i%2 == 0 checks if i is even (no remainder when divided by 2)
	for i := range 100 {
		if i%2 == 0 {
			fmt.Printf("i is even: %v\n", i)
		}
	}
}

/*
REMARKS:
- MODULUS (%) returns the REMAINDER after integer division
- Classic use cases:
    ● Even/odd check: n%2 == 0 (even), n%2 != 0 (odd) — used in Lesson 072 with continue
    ● Wrap-around: index%length keeps values within bounds (circular buffers)
    ● Every Nth item: i%5 == 0 (every 5th iteration)
    ● Clock arithmetic: minutes%60, hours%24
- INTEGER DIVISION (/) in Go TRUNCATES — it drops the fractional part
  — 83/40 = 2 (not 2.075) because both operands are int
  — for precise division, convert to float64 first: float64(83)/float64(40) (Lesson 017)
- To print a literal % in fmt.Printf, use %% (otherwise Go thinks it's a format verb)
- "for i := range 100" is a Go 1.22+ feature — range over an integer
  — equivalent to: for i := 0; i < 100; i++
  — cleaner syntax for simple counted loops
- BUILDING ON LESSON 072: modulus is frequently combined with for loops and continue
  for filtering iterations (skip odd, process every 3rd, etc.)
*/
