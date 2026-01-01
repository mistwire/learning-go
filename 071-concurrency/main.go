package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Concurrency and parallelism are related but distinct concepts in programming:
● Concurrency refers to code that is written in a concurrent design pattern. This means
that the code has the potential ability to execute multiple tasks simultaneously, where
each task may make progress independently of the others.
	○ In Go, concurrency is achieved using goroutines, lightweight threads of execution
	that are managed by the Go runtime.
	○ A Go program can create many goroutines that run concurrently, each performing
	a different task.
	○ The communication and synchronization of these goroutines is typically done
	using channels, which provide a way for goroutines to exchange data and
	coordinate their execution.
● Parallelism, on the other hand, refers to the ability of a program to execute multiple
tasks simultaneously by utilizing multiple CPUs or cores.
	○ Parallelism can often speed up the execution of a program by allowing multiple
	parts of the program to run in parallel on different processors. In Go, parallelism
	can be achieved by running multiple goroutines on different processors using the
	go keyword.
		■ serial / sequential execution
			● the opposite of parallel computing
			● The opposite of code running in parallel is code running serially or
			sequentially. In sequential execution, each instruction or task is
			executed one after the other in a predefined order, so that each
			instruction must wait for the previous one to finish before it can
			start. This differs from parallel execution, where multiple
			instructions or tasks can be executed simultaneously. Sequential
			execution is typically used when the instructions or tasks are
			dependent on each other, or when the resources required to
			execute the code are limited. Parallel execution, on the other
			hand, is used to speed up the execution of code by running
			Todd McLeod - Learn To Code Go - Page 48
			multiple instructions or tasks at the same time, often on multiple
			processors or cores.
● Go makes it easy to write concurrent code using goroutines and channels.
*/

func main() {
	// CONDITIONAL
	// if statements
	// switch statements

	// concurrency
	// select a channel

	ch1 := make(chan int)
	ch2 := make(chan int)

	// get an int64, convert it to type time.Duration, then use it in time.Sleep
	// Int63n returns an int64
	// type duration's underlying type is int64
	// time.Sleep takes type duration
	// time.Millisecond is a constant in the time package
	// https://pkg.go.dev/time#pkg-constants
	d1 := time.Duration(rand.Int63n(250))
	d2 := time.Duration(rand.Int63n(250))
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
}
