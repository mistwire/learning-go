/*
● SWITCH & SELECT — multi-way branching and channel selection
● Building on Lesson 067-068: switch is a cleaner alternative to long if/else-if chains
● TWO forms: expression switch (switch x) and "variable-less" switch (switch {})
● Go switch does NOT fall through by default (opposite of C/Java)
● SELECT is like switch but for CHANNEL operations (concurrency preview)
● This is a comprehensive lesson — also previews for loops (Lesson 072) and
  for-range (Lesson 074) with slices and maps
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	x := 40
	y := 5
	fmt.Printf(" x=%v \n y=%v\n", x, y)

	// FORM 1: "variable-less" switch — cases are boolean expressions (like if/else-if)
	switch {
	case x < 42:
		fmt.Println("x is LESS THAN 42")
	case x == 42:
		fmt.Println("x is EQUAL TO 42")
	case x > 42:
		fmt.Println("x is GREATER THAN 42")
	default:
		fmt.Println("this is the default case for x")
	}

	// FORM 2: expression switch — compares x against each case value
	switch x {
	case 40:
		fmt.Println("x is 40")
	case 41:
		fmt.Println("x is 41")
	case 42:
		fmt.Println("x is 42")
	case 43:
		fmt.Println("x is 43")
	default:
		fmt.Println("this is the default case for x")
	}

	// FALLTHROUGH: Go does NOT fall through by default (unlike C/Java)
	// you must explicitly use the "fallthrough" keyword
	// fallthrough executes the NEXT case's body unconditionally (doesn't check its condition)
	switch x {
	case 40:
		fmt.Println("x is 40")
		fallthrough // falls into case 41's body — but STOPS there (no further fallthrough)
	case 41:
		fmt.Println("printed because of fallthrough: x is 41")
	case 42:
		fmt.Println("printed because of fallthrough: x is 42")
	case 43:
		fmt.Println("printed because of fallthrough: x is 43")
	default:
		fmt.Println("printed because of fallthrough: this is the default case for x")
	}

	// CHAINED FALLTHROUGH: every case with fallthrough cascades into the next
	switch x {
	case 40:
		fmt.Println("x is 40")
		fallthrough
	case 41:
		fmt.Println("printed because of ALL OF THE fallthrough statements: x is 41")
		fallthrough
	case 42:
		fmt.Println("printed because of ALL OF THE fallthrough statements: x is 42")
		fallthrough
	case 43:
		fmt.Println("printed because of ALL OF THE fallthrough statements: x is 43")
		fallthrough
	default:
		fmt.Println("printed because of ALL OF THE fallthrough statements: this is the default case for x")
	}

	// --- SELECT STATEMENT (concurrency preview) ---
	// select is like switch but for CHANNEL operations

	ch1 := make(chan int) // unbuffered channel of type int
	ch2 := make(chan int)

	// type conversion: int64 → time.Duration (Lesson 017 pattern)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	d1 := time.Duration(r.Int63n(250))
	d2 := time.Duration(r.Int63n(250))

	// GOROUTINES: "go" keyword launches a function in a new lightweight thread
	go func() {
		time.Sleep(d1 * time.Millisecond)
		ch1 <- 41 // send 41 into channel 1
	}()

	go func() {
		time.Sleep(d2 * time.Millisecond)
		ch2 <- 42 // send 42 into channel 2
	}()

	// SELECT: whichever channel receives a value FIRST wins
	// https://go.dev/ref/spec#Select_statements
	select {
	case v1 := <-ch1:
		fmt.Println("value from channel 1", v1)
	case v2 := <-ch2:
		fmt.Println("value from channel 2", v2)
	}

	// --- FOR LOOPS PREVIEW (detailed in Lessons 072-075) ---

	// C-style for loop: init; condition; post
	for i := 0; i < 5; i++ {
		fmt.Printf("counting to five: %v \t first  for loop\n", i)
	}

	// while-style: condition only
	for y < 10 {
		fmt.Printf("y is %v \t\t\t second for loop\n", y)
		y++
	}

	// infinite loop with break
	for {
		fmt.Printf("y is %v \t\t third  for loop\n", y)
		if y > 20 {
			break
		}
		y++
	}

	// continue — skip to next iteration
	for i := 0; i < 20; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println("counting even numbers: ", i)
	}

	// nested loops
	for i := 0; i < 5; i++ {
		fmt.Println("--")
		for j := 0; j < 5; j++ {
			fmt.Printf("outer loop %v \t inner loop %v\n", i, j)
		}
	}

	// for-range over a slice (preview of Lessons 096-108)
	xi := []int{42, 43, 44, 45, 46, 47}
	for i, v := range xi {
		fmt.Println("ranging over a slice", i, v)
	}

	// for-range over a map (preview of Lessons 116-119)
	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}
	for k, v := range m {
		fmt.Println("ranging over a map", k, v)
	}

}

/*
REMARKS:
- SWITCH vs IF/ELSE-IF: switch is cleaner when comparing one value against many cases
- Go switch does NOT fall through by default — this is the OPPOSITE of C/Java/JavaScript
  (you don't need "break" at the end of each case)
- "fallthrough" is rarely used in practice — it unconditionally executes the next case's body
  WITHOUT checking that case's condition
- You can have MULTIPLE values in one case: case 40, 41, 42:
- SELECT is switch for channels:
    ● Blocks until one case can proceed
    ● If multiple cases are ready, one is chosen at RANDOM
    ● Used for multiplexing goroutine communication
- BUILDING ON LESSON 067: switch replaces long if/else-if chains
- BUILDING ON LESSON 069: switch can also have an init statement: switch x := f(); { ... }
- This comprehensive lesson previews many concepts covered individually later:
    ● For loops (Lesson 072), nested loops (073), for-range (074)
    ● Slices (098), maps (116)
    ● Concurrency/goroutines (071)
*/
