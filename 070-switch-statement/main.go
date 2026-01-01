package main

import (
	"fmt"
	"math/rand"
	"time" 
)

/*
A switch statement is a control flow statement in programming that allows a program to evaluate
an expression or variable and then selectively execute different blocks of code based on the
value of the expression or variable. It provides a convenient way to write code that performs
different actions based on a single variable or expression without having to use multiple if/else
statements. The switch statement typically consists of a single expression or variable followed
by a series of case statements that specify the different values that the expression or variable
can take and the corresponding blocks of code to execute in each case. If the expression or
variable matches one of the specified case values, the corresponding block of code is executed,
and if none of the cases match, an optional default case can be used to execute a fallback block
of code.
https://go.dev/play/p/JuPB00o3YNC
*/




func main() {
	x := 40
	y := 5  // this is the fourth statement to run
	fmt.Printf(" x=%v \n y=%v\n", x, y)

	switch { // "variable-less" switch statement
	case x < 42:
		fmt.Println("x is LESS THAN 42")
	case x == 42:
		fmt.Println("x is EQUAL TO 42")
	case x > 42:
		fmt.Println("x is GREATER THAN 42")
	default:
		fmt.Println("this is the default case for x")
	}

	switch x { // you can put the variable here as well
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

	// no default fallthrough
	switch x {
	case 40:
		fmt.Println("x is 40")
		fallthrough
	case 41:
		fmt.Println("printed because of fallthrough: x is 41")
	case 42:
		fmt.Println("printed because of fallthrough: x is 42")
	case 43:
		fmt.Println("printed because of fallthrough: x is 43")
	default:
		fmt.Println("printed because of fallthrough: this is the default case for x")
	}

	// no default fallthrough
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

	// concurrency
	// switch on channel

	ch1 := make(chan int)
	ch2 := make(chan int)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// get an int64, convert it to type time.Duration, then use it in time.Sleep
	// Int63n returns an int64
	// type duration's underlying type is int64
	// time.Sleep takes type duration
	// time.Millisecond is a constant in the time package
	// https://pkg.go.dev/time#pkg-constants
	d1 := time.Duration(r.Int63n(250))
	d2 := time.Duration(r.Int63n(250))
	// fmt.Printf("%v \t %T\n", d1, d1)
	// time.Sleep(d1 * time.Millisecond)
	// fmt.Println("Done")

	go func() {
		time.Sleep(d1 * time.Millisecond)
		ch1 <- 41
	}()

	go func() {
		time.Sleep(d2 * time.Millisecond)
		ch2 <- 42
	}()

	// A "select" statement chooses which of a set of possible send or receive operations will proceed.
	// It looks similar to a "switch" statement but with the cases all referring to communication operations.
	// https://go.dev/ref/spec#Select_statements
	select {
	case v1 := <-ch1:
		fmt.Println("value from channel 1", v1)
	case v2 := <-ch2:
		fmt.Println("value from channel 2", v2)
	}

	// LOOPS / INTERATIVE
	// for statements

	/*
		The Go for loop is similar to—but not the same as—C's.
		It unifies for and while and there is no do-while.
		There are three forms, only one of which has semicolons.

		// Like a C for
		for init; condition; post { }

		// Like a C while
		for condition { }

		// Like a C for(;;)
		for { }

	*/
	// https://go.dev/doc/effective_go#for

	for i := 0; i < 5; i++ {
		fmt.Printf("counting to five: %v \t first  for loop\n", i)
	}

	for y < 10 {
		fmt.Printf("y is %v \t\t\t second for loop\n", y)
		y++
	}

	// break
	// takes you out of the loop
	for {
		fmt.Printf("y is %v \t\t third  for loop\n", y)
		if y > 20 {
			break
		}
		y++
	}

	// continue
	// takes to next iteration
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

	// for range loop
	// data structures - slice
	xi := []int{42, 43, 44, 45, 46, 47}
	for i, v := range xi {
		fmt.Println("ranging over a slice", i, v)
	}

	// for range loop
	// data structures - map
	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}
	for k, v := range m {
		fmt.Println("ranging over a map", k, v)
	}

}
