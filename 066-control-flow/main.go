/*
● CONTROL FLOW INTRO — sequence of execution and the init() function
● Go programs execute in a predictable order: init() → main()
● Within a function, statements execute SEQUENTIALLY (top to bottom)
● Control flow categories: SEQUENCE, CONDITIONAL (if/switch), ITERATIVE (for loops)
*/
package main

import (
	"fmt"
)

// init() runs AUTOMATICALLY before main() — no need to call it
// used for package-level setup (initializing variables, validating config, etc.)
// a package can have MULTIPLE init() functions (even in the same file)
func init() {
	fmt.Println("init func runs before main")
}

func main() {
	// SEQUENCE — statements execute top-to-bottom, one after another
	fmt.Println("this is the first func main statement to run")
	fmt.Println("this is the second statement to run")
	x := 40 // this is the third statement to run
	y := 5  // this is the fourth statement to run
	fmt.Printf(" x=%v \n y=%v\n", x, y)
}

/*
REMARKS:
- Go has THREE categories of control flow:
    1. SEQUENCE — default, top-to-bottom execution (this lesson)
    2. CONDITIONAL — if/else, switch (Lessons 067-070)
    3. ITERATIVE — for loops (Lessons 072-075)
- init() is a special function:
    ● Runs before main() automatically — never called directly
    ● A single file can have MULTIPLE init() functions
    ● Execution order: package-level var init → init() functions → main()
    ● Common uses: database connections, config validation, registering drivers
- Unlike most languages, Go does NOT have while, do-while, or foreach keywords
  — "for" is the only loop construct (Lesson 072)
- There is no ternary operator (condition ? a : b) in Go — use if/else instead
- BUILDING ON LESSON 033 (scope): init() and main() are at package scope
  — variables declared inside them are block-scoped
*/
