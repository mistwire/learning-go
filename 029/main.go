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
	fmt.Printf("%d \t\t %b \t\t %#x\n", x, x, x)
	fmt.Printf("%d \t\t %b \t\t %#x\n", y, y, y)
	fmt.Printf("%d \t\t %b \t %#x\n", z, z, z)

}
