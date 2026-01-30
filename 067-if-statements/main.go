/*
● IF STATEMENTS — conditional branching
● Building on Lesson 066: moving from SEQUENCE to CONDITIONAL control flow
● Three forms: if, if/else, if/else-if/else
● The condition must be a BOOLEAN expression — no truthy/falsy like JavaScript
*/
package main

import (
	"fmt"
)

func main() {
	x := 42
	y := 5
	fmt.Printf(" x=%v \n y=%v\n", x, y)

	// FORM 1: if only — executes block only when condition is true
	if x < 42 {
		fmt.Println("Less than the meaning of life")
	}

	// FORM 2: if/else — two branches, one MUST execute
	if x < 42 {
		fmt.Println("Less than the meaning of life")
	} else {
		fmt.Println("equal to, or greater than, the meaning of life")
	}

	// FORM 3: if/else-if/else — multiple conditions checked in order
	// first true condition wins; else is the fallback
	if x < 42 {
		fmt.Println("Less than the meaning of life")
	} else if x == 42 {
		fmt.Println("equal to the meaning of life")
	} else {
		fmt.Println("greater than the meaning of life")
	}

	/*
		"If" statements specify the conditional execution of two branches
		according to the value of a boolean expression. If the expression evaluates
		to true, the "if" branch is executed, otherwise, if present, the "else" branch is executed.
	*/
	// https://go.dev/ref/spec#If_statements

}

/*
REMARKS:
- BUILDING ON LESSON 066: if/else is the first CONDITIONAL control flow mechanism
- Go conditions MUST evaluate to a boolean — no implicit truthiness
  (unlike Python/JS where 0, "", nil are "falsy")
- No parentheses around the condition: if x < 42 { ... } (NOT if (x < 42) { ... })
- Opening brace { MUST be on the same line as the if/else — Go enforces this
- else/else-if MUST be on the same line as the closing } of the previous block
- Go has NO ternary operator — use if/else instead:
    // Not valid Go: result := x > 0 ? "positive" : "negative"
    // Instead: if x > 0 { result = "positive" } else { result = "negative" }
- Conditions can also include a statement before the expression (Lesson 069: comma-ok idiom)
  e.g., if v, ok := m["key"]; ok { ... }
*/
