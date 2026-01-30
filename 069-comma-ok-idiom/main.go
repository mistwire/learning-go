/*
● COMMA-OK IDIOM — statement-before-condition in if statements
● Building on Lesson 067: an if statement can have an INIT STATEMENT before the condition
● Syntax: if <init statement>; <condition> { ... }
● The variable declared in the init statement is SCOPED to the if/else block
● This pattern is used heavily with maps (Lesson 119) and type assertions
*/
package main

import (
	"fmt"
	"math/rand"
)

/*
	The expression [evaluated in an if statement] may be preceded by a simple statement,
	which executes before the expression is evaluated.
*/
// https://go.dev/ref/spec#If_statements
/*
	The comma ok idiom is also like this
*/
// https://go.dev/play/p/OXGzjxVkag0

func main() {
	x := 40

	// INIT STATEMENT + CONDITION: z is declared, then the condition is checked
	// z exists ONLY inside this if/else block (block scope from Lesson 033)
	if z := 2 * rand.Intn(x); z >= x {
		fmt.Printf("z is %v and that is GREATER THAN OR EQUAL x which is %v\n", z, x)
	} else {
		fmt.Printf("z is %v and that is LESS THAN x which is %v\n", z, x)
	} // the scope of z ends here — z is NOT accessible after this line
}

/*
REMARKS:
- The COMMA-OK IDIOM combines two Go features:
    1. Init statement in if: if <statement>; <condition> { ... }
    2. Multiple return values: v, ok := someFunction()
- The classic form (seen in Lessons 119 with maps): if v, ok := m["key"]; ok { ... }
    ● v gets the value, ok gets a boolean indicating if the key exists
    ● ok limits the scope of v and ok to the if/else block
- BUILDING ON LESSON 033 (scope): the init statement creates block-scoped variables
  — they don't leak into the surrounding function scope
- BUILDING ON LESSON 067: this is an extension of if statements, not a new construct
- This pattern appears with:
    ● Map lookups: if v, ok := m[key]; ok { ... } (Lesson 119)
    ● Type assertions: if v, ok := i.(string); ok { ... }
    ● Channel receives: if v, ok := <-ch; ok { ... }
- The semicolon separates the init statement from the condition
  — similar to the first clause of a for loop (Lesson 072)
*/
