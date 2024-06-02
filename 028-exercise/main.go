/*
Write a program that uses print verbs (https://pkg.go.dev/fmt) to show the following numbers:
- 747
- 911
- 90210
as
- decimal
- binary
- hexadecimal
*/

package main

import "fmt"

func main(){
	x, y, z := 747, 911, 90210
	fmt.Printf("%v\t as decimal: %d\t as binary: %b\t\t as hex: %X\n",x ,x, x, x)
	fmt.Printf("%v\t as decimal: %d\t as binary: %b\t\t as hex: %X\n",y ,y, y, y)
	fmt.Printf("%v\t as decimal: %d\t as binary: %b\t as hex: %X\n",z ,z, z, z)

}