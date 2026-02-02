/*
Syntax of functions in Go
● func (receiver) identifier(parameters) (returns) { code }
	○ parameters & arguments
		■ we define our func with parameters (if any)
		■ we call our func and pass in arguments (in any)
● Everything in Go is PASS BY VALUE
● sample funcs
	○ no params, no returns
	○ 1 param, no returns
	○ 1 param, 1 return
	○ 2 params, 2 returns
*/

// func (r receiver) identifier(p parameter(s)) (return(s)) { <your code here> }
package main

import "fmt"

func main() {
	foo()

	bar("Chris")

	s := aloha("Williams")
	fmt.Println(s)

	s1, n := dogYears("Chris", 50)
	fmt.Println(s1, n)
}

func foo() {
	fmt.Println("I am from foo")
}

func bar(s string) {
	fmt.Println("My name is", s)
}

func aloha(s string) string {
	return fmt.Sprint("They call me Mr. ", s)
}

func dogYears(name string, age int) (string, int) {
	age *= 7
	return fmt.Sprint(name, " is this old in dog years:"), age
}