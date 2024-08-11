package main

import "fmt"

var mystring string = "Hello Gophers!"

const Pi = 3.14159 

func main () {
	myint := 42
	fmt.Println(mystring)
	fmt.Println(Pi)
	fmt.Printf("The value of mystring is %v and the type is %T\n", mystring, mystring)
	fmt.Printf("The value of const Pi is %v and the type is %T\n", Pi, Pi)
	fmt.Printf("The value of myint is %v and the type is %T", myint, myint)
}