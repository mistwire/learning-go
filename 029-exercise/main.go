/* https://go.dev/ref/spec#Numeric_types 
write a program that declares 2 variables:
- one var to store a value of type int8
	- assign it to the largest possible number then print it
- one var to store the value of type uint8
	- assign it to the largest possible number then print it
*/

package main

import "fmt"

func main(){
	var x int8 = 127
	var y uint8 = 255
	fmt.Printf("%v is of type %T\n",x, x)
	fmt.Printf("%v is of type %T\n",y, y)
}