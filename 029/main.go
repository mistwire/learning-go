/*
● EXERCISE — Number formatting: decimal, binary, hexadecimal
● Building on Lesson 028: expanding the format verb toolkit
● %d = decimal, %b = binary, %#x = hexadecimal with 0x prefix
● \t inserts a tab for alignment
*/
// write a program that uses print verbs to show the following numbers
// 	● 747
// 	● 911
// 	● 90210
// 	as
// 	● decimal
// 	● binary
// 	● hexadecimal

package main

import "fmt"

func main() {
	x, y, z := 747, 911, 90210
	// each number printed in three representations: decimal, binary, hex
	fmt.Printf("%d \t\t %b \t\t %#x\n", x, x, x)
	fmt.Printf("%d \t\t %b \t\t %#x\n", y, y, y)
	fmt.Printf("%d \t\t %b \t %#x\n", z, z, z)

}

/*
REMARKS:
- BUILDING ON LESSON 028: from basic %v/%T to numeric-specific format verbs
- KEY FORMAT VERBS for integers:
    %d — decimal (base 10)
    %b — binary (base 2) — connects to Lesson 025's bit shifting examples
    %o — octal (base 8)
    %x — hexadecimal (base 16, lowercase)
    %X — hexadecimal (base 16, uppercase)
    %#x — hexadecimal WITH "0x" prefix (the # flag adds the prefix)
- The # flag works with other bases too: %#o adds a leading 0, %#b adds 0b
- \t is a tab character — useful for aligning columnar output
- Same value, same variable — just different REPRESENTATIONS of the same number
  (the underlying bits don't change, only how they're displayed)
- This ties back to Lesson 025 where we saw binary representations of bit-shifted values
*/
