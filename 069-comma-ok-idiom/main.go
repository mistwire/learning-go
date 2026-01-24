// LESSON: Comma-Ok Idiom - statement before condition in if, scoped variables
package main

import (
	"fmt"
	"math/rand"
)

/*
	The expression [evaluated in an if statement ]may be preceded by a simple statement,
	which executes before the expression is evaluated.
*/
// https://go.dev/ref/spec#If_statements
/*
	The comma ok idiom is also like this
*/
// https://go.dev/play/p/OXGzjxVkag0

func main() {
	//SEQUENCE
	x := 40

	if z := 2 * rand.Intn(x); z >= x {
		fmt.Printf("z is %v and that is GREATER THAN OR EQUAL x which is %v\n", z, x)
	} else {
		fmt.Printf("z is %v and that is LESS THAN x which is %v\n", z, x)
	} // the scope of z ends here
}
