// LESSON: Basics - comments, packages, imports, constants, fmt.Printf verbs, raw strings
// Comments go in like this
/*
Or like this
*/

package main

import ( // uses parens for imports, no comma after each package
	"fmt" // https://pkg.go.dev/fmt@go1.22.2
)

func main() {
	const name = "Kim"
	const age = 22

	// const name, age = "Kim", 22 also works

	fmt.Println("Hello, Gophers🥰😎😘") // UTF-8 means emojis work 🥳
	fmt.Printf("%s is %d years old.\n", name, age)
	fmt.Printf("%s is %d years old:\n\tand the type is %T and %T", name, age, name, age)

	// raw string literal (uses backticks NOT single quotes!!!😅)
	fmt.Println(`
	UTF-8 saves space. 
In UTF-8, common characters like “C” take 8 bits, 
	while rare characters like “💕” take 32 bits. 
		Other characters take 16 or 24 bits. 
	A blog post like this one takes about 
	four times less space in UTF-8 
	than it would in UTF-32. 
	So it loads four times faster.
	`)
	// https://go.dev/ref/spec#String_literals
}
