package main

import "fmt"

func main() {
	x := 40 // this is the third statement to run
	y := 5  // this is the fourth statement to run
	fmt.Printf(" x=%v \n y=%v\n", x, y)

	if x < 42 && y < 42 {
		fmt.Println("both are less than the meaning of life")
	}

	if x > 30 || x < 42 {
		fmt.Println("x is getting close to the meaning of life")
	}

	if x != 42 && y != 10{
		fmt.Println("x is not 42 & y is not 10")
	}

	/*
		Logical operators
		Logical operators apply to boolean values
		and yield a result of the same type as the operands.
		The right operand is evaluated conditionally.

		&&    conditional AND    p && q  is  "if p then q else false"
		||    conditional OR     p || q  is  "if p then true else q"
		!     NOT                !p      is  "not p"
	*/
	// https://go.dev/ref/spec#Logical_operators

	/*
		The expression [evaluated in an if statement ]may be preceded by a simple statement,
		which executes before the expression is evaluated.
	*/
	// https://go.dev/ref/spec#If_statements
	/*
		The comma ok idiom is also like this
	*/
	// https://go.dev/play/p/OXGzjxVkag0
}
