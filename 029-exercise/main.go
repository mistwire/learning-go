// write a program that uses print verbs to show the following numbers
// ● 747
// ● 911
// ● 90210
// as
// ● decimal
// ● binary
// ● hexadecimal

package main

import "fmt"

func main() {
	a := 747
	b := 911
	c := 90210

	fmt.Printf("decimal: %v\t binary: %#b\t\t hex: %#x\n",a, a, a)
	fmt.Printf("decimal: %v\t binary: %#b\t\t hex: %#x\n",b, b, b)
	fmt.Printf("decimal: %v\t binary: %#b\t hex: %#x\n",c, c, c)

}
