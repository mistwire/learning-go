package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// concurrency
	// switch on channel

	ch1 := make(chan int)
	ch2 := make(chan int)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// https://pkg.go.dev/time#pkg-constants
	d1 := time.Duration(r.Int63n(250))
	d2 := time.Duration(r.Int63n(250))
	fmt.Printf("%v \t %T\n", d1, d1)
	fmt.Printf("%v \t %T\n", d2, d2)
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
}
