/*
● CONCURRENCY — goroutines, channels, and select
● Building on Lesson 070: select is the channel-based counterpart to switch
● CONCURRENCY vs PARALLELISM:
  ● Concurrency = DESIGN pattern (structuring code to handle multiple tasks)
  ● Parallelism = EXECUTION (actually running tasks simultaneously on multiple CPUs)
● Go's concurrency primitives: goroutines (lightweight threads) + channels (communication)
*/
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
● Go makes it easy to write concurrent code using goroutines and channels.
*/

func main() {
	// make(chan int) creates an UNBUFFERED channel — send blocks until a receiver is ready
	ch1 := make(chan int)
	ch2 := make(chan int)

	// TYPE CONVERSION (Lesson 017): int64 → time.Duration
	// time.Duration's underlying type is int64
	d1 := time.Duration(rand.Int63n(250))
	d2 := time.Duration(rand.Int63n(250))

	// GOROUTINE: "go" keyword launches a function concurrently
	// "go func() { ... }()" is an anonymous function launched as a goroutine
	// the () at the end immediately invokes it
	go func() {
		time.Sleep(d1 * time.Millisecond)
		ch1 <- 41 // SEND: pushes 41 into ch1
	}()

	go func() {
		time.Sleep(d2 * time.Millisecond)
		ch2 <- 42 // SEND: pushes 42 into ch2
	}()

	// SELECT: blocks until one channel operation can proceed
	// whichever goroutine finishes sleeping first "wins"
	// https://go.dev/ref/spec#Select_statements
	select {
	case v1 := <-ch1: // RECEIVE: pulls a value from ch1
		fmt.Println("value from channel 1", v1)
	case v2 := <-ch2: // RECEIVE: pulls a value from ch2
		fmt.Println("value from channel 2", v2)
	}
}

/*
REMARKS:
- GOROUTINES are NOT OS threads — they're multiplexed onto threads by the Go runtime
  ● Start with ~2KB stack (threads start with ~1MB)
  ● You can run thousands/millions of goroutines
  ● Launched with the "go" keyword: go myFunction()
- CHANNELS are typed conduits for goroutine communication
  ● ch := make(chan int) — unbuffered (send blocks until receive, and vice versa)
  ● ch := make(chan int, 10) — buffered (send only blocks when buffer is full)
  ● ch <- value — send into channel
  ● value := <-ch — receive from channel
- SELECT (from Lesson 070): multiplexes across multiple channels
  ● Blocks until at least one case can proceed
  ● If multiple are ready, picks one at random
  ● "default" case makes select non-blocking
- KEY DISTINCTION:
  ● Concurrency = dealing with many things at once (structure/design)
  ● Parallelism = doing many things at once (execution)
  ● "Concurrency is about structure, parallelism is about execution" — Rob Pike
- BUILDING ON LESSON 066: concurrency breaks the sequential execution model
  — goroutines may execute in any order
- Go's mantra: "Do not communicate by sharing memory; share memory by communicating"
  — channels are safer than shared variables with mutexes
*/
