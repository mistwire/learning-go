/*
● DEFER — a statement that delays a function call until the SURROUNDING FUNCTION returns
● Deferred calls are placed on a STACK (LIFO — Last In, First Out)
	○ If you defer multiple calls, the LAST deferred runs FIRST
● A deferred function runs when the surrounding function:
	○ executes a RETURN statement
	○ reaches the END of its function body
	○ or the goroutine is PANICKING
● Similar to `finally` in Java/Python or `ensure` in Ruby
● BUILDING ON LESSON 134: uses the same func syntax — defer just controls WHEN it runs
*/
package main

import "fmt"

func main() {
	// defer pushes foo() onto a deferred-call stack
	// foo() will NOT run here — it waits until main() is about to return
	defer foo()

	// bar() runs immediately — it is NOT deferred
	bar()

	// execution order: bar() prints first, then foo() prints as main() exits
	// output:
	//   bar
	//   foo
}

// func (r receiver) identifier(p parameter(s)) (return(s)) { code }

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}

/*
REMARKS:
- DEFER delays a function call — it does NOT skip it; the call always runs (unless os.Exit is called)
- Deferred calls execute in LIFO order: last defer in = first to run
- Deferred arguments are evaluated IMMEDIATELY, only the call itself is delayed:
    x := 1
    defer fmt.Println(x) // captures x=1 NOW, prints 1 later even if x changes
    x = 2
- Common real-world uses:
    ■ Closing files: `defer f.Close()` right after opening — keeps open/close logic together
    ■ Unlocking mutexes: `defer mu.Unlock()` right after Lock()
    ■ Logging/timing: `defer logElapsed(time.Now())` at function start
- BUILDING ON LESSON 134: defer works with any function call — named functions, methods,
  closures, even function literals; the func syntax is unchanged
- BUILDING ON LESSON 071: defer is especially useful with goroutines and cleanup in
  concurrent code
- KEY PATTERN: open a resource, immediately defer its cleanup — this is idiomatic Go
  and prevents resource leaks even when errors cause early returns
*/
