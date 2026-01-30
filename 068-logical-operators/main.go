/*
● LOGICAL OPERATORS — combining conditions with && (AND), || (OR), ! (NOT)
● Building on Lesson 067: using COMPARISON OPERATORS to produce boolean values,
  then combining them with LOGICAL OPERATORS
● SHORT-CIRCUIT evaluation: Go stops evaluating as soon as the result is determined
*/
package main

import (
	"fmt"
)

func init() {
	fmt.Println("init func runs before main")
}

func main() {
	fmt.Println("this is the first func main statement to run")
	fmt.Println("this is the second statement to run")
	x := 40
	y := 5
	fmt.Printf(" x=%v \n y=%v\n", x, y)

	/*
		Comparison operators — produce boolean values (true/false)
		==    equal
		!=    not equal
		<     less
		<=    less or equal
		>     greater
		>=    greater or equal
	*/
	// https://go.dev/ref/spec#Comparison_operators

	// && (AND): BOTH conditions must be true
	if x < 42 && y < 42 {
		fmt.Println("both are less than the meaning of life")
	}

	// || (OR): AT LEAST ONE condition must be true
	if x > 30 || x < 42 {
		fmt.Println("x is getting close to the meaning of life")
	}

	// ! (NOT) combined with != (not equal)
	if x != 42 {
		fmt.Println("x is not 42")
	}

	/*
		Logical operators — combine boolean expressions
		&&    conditional AND    p && q  is  "if p then q else false"
		||    conditional OR     p || q  is  "if p then true else q"
		!     NOT                !p      is  "not p"
	*/
	// https://go.dev/ref/spec#Logical_operators

}

/*
REMARKS:
- BUILDING ON LESSON 067: logical operators let you combine multiple if conditions
- TWO categories:
    ● COMPARISON operators: ==, !=, <, <=, >, >= (produce a bool)
    ● LOGICAL operators: &&, ||, ! (combine bools)
- SHORT-CIRCUIT EVALUATION:
    ● && stops at the first false (if p is false, q is never evaluated)
    ● || stops at the first true (if p is true, q is never evaluated)
    — this matters when the right side has side effects (e.g., function calls)
- Operator precedence (highest to lowest): !, comparison ops, &&, ||
    — use parentheses to make complex conditions clear
- Go does NOT have bitwise logical operators for bools
  (& and | work on integers for bitwise ops, see Lesson 025)
- These operators are used everywhere: if statements, for loop conditions, switch cases
*/
