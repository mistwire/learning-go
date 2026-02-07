/*
● WRAPPER FUNCTION with INTERFACE PARAMETER — write your own functions that accept
  interface types, enabling POLYMORPHISM in your code
● BUILDING ON LESSON 140: fmt.Println accepts fmt.Stringer — now YOU write a function
  that does the same thing: accept ANY type that satisfies an interface
● This is the power of interfaces: ONE FUNCTION works with MANY TYPES
  ○ logInfo doesn't know about book or count
  ○ logInfo only knows about fmt.Stringer
  ○ Any type with String() string can be passed to logInfo — including types
    that don't exist yet!
● DECOUPLING: the wrapper function depends on BEHAVIOR (the interface), not
  on concrete types — this is a core Go design pattern
*/
package main

import (
	"fmt"
	"log"
	"strconv"
)

// book is a struct type — see Lesson 125
type book struct {
	title string
}

// func (r receiver) identifier(p parameter(s)) (return(s)) { code }

// String() makes book satisfy fmt.Stringer (Lesson 140)
func (b book) String() string {
	return fmt.Sprint("The title of the book is ", b.title)
}

// count is a named type based on int (Lesson 140)
type count int

// String() makes count satisfy fmt.Stringer
func (c count) String() string {
	return fmt.Sprint("This is the number ", strconv.Itoa(int(c)))
}

// logInfo is a WRAPPER FUNCTION — it accepts fmt.Stringer, not a concrete type
// This means ANY type that implements String() string can be passed here
// The function is "decoupled" from book and count — it doesn't import them,
// doesn't know they exist — it only knows the INTERFACE
func logInfo(s fmt.Stringer) {
	// s.String() works on ANY type passed in — this is POLYMORPHISM
	// Go dispatches to the correct String() method at runtime
	log.Println("Log from 141", s.String())
}

func main() {
	b := book{
		title: "West With The Night",
	}

	var n count = 42

	// Both book and count can be passed to logInfo
	// because they BOTH satisfy the fmt.Stringer interface
	// The function doesn't care about the concrete type — only that
	// the value has a String() method
	logInfo(b) // calls book.String() → "The title of the book is West With The Night"
	logInfo(n) // calls count.String() → "This is the number 42"
}

/*
REMARKS:
- A WRAPPER FUNCTION accepts an interface type, letting it work with any value
  that satisfies that interface — this is Go's primary polymorphism pattern
- The interface parameter (fmt.Stringer) acts as a CONTRACT: "I don't care what
  type you are, as long as you can give me a String()"
- DECOUPLING: logInfo depends on behavior (String() method), not concrete types
  ○ You could add a new type `type user struct{...}` with a String() method
  ○ It would work with logInfo immediately — NO changes to logInfo required
  ○ This is the OPEN/CLOSED PRINCIPLE: open for extension, closed for modification
- BUILDING ON LESSON 140: we used the Stringer interface from stdlib; now we
  WRITE OUR OWN FUNCTION that accepts it
- BUILDING ON LESSON 139: interfaces are satisfied IMPLICITLY — book and count
  don't declare "I implement Stringer"; Go checks at compile time
- BUILDING ON LESSON 134: functions take parameters — here the parameter type
  is an INTERFACE instead of a concrete type like int or string
- WHY USE WRAPPER FUNCTIONS?
  ○ Add consistent behavior (logging, metrics, error handling) across types
  ○ Create adapters between different parts of your system
  ○ Enable testing by accepting interfaces instead of concrete dependencies
- GO PHILOSOPHY: "Accept interfaces, return structs" — functions should accept
  the narrowest interface they need, but return concrete types when possible
- NEXT STEP: you'll see this pattern everywhere in Go — http.Handler, io.Reader,
  sort.Interface — the stdlib is built on small interfaces and wrapper functions
*/
