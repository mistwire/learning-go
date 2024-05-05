// Comments go in like this
/*
Or like this
*/

package main

import "fmt" // https://pkg.go.dev/fmt@go1.22.2

func main(){
	const name = "Kim"
	const age = 22

	// const name, age = "Kim", 22 also works

	fmt.Println("Hello, Gophers🥰") // UTF-8 means emojis work 
	fmt.Printf("%s is %d years old.\n", name, age)
	fmt.Printf("%s is %d years old:\n\tand the type is %T and %T", name, age, name, age)
}

