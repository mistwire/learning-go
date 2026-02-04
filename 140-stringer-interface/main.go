/*
● STRINGER INTERFACE — the `fmt.Stringer` interface requires a single method:
	○ String() string
● Any type that implements String() controls how it appears when printed with fmt.Println,
  fmt.Sprintf, etc. — this is Go's equivalent of toString() in Java or __str__ in Python
● You can implement Stringer on ANY NAMED TYPE — structs, custom ints, etc.
● BUILDING ON LESSON 139: Stringer is a real-world INTERFACE from the standard library;
  fmt.Println uses it polymorphically — if a value satisfies fmt.Stringer, its String()
  method is called automatically
● BUILDING ON LESSON 138: String() is a METHOD with a receiver, just like speak() was
*/
package main

import (
	"fmt"
	"strconv" // strconv.Itoa converts int to ASCII string — "Integer to ASCII"
)

// book is a struct type — see Lesson 125
type book struct {
	title string
}

// func (r receiver) identifier(p parameter(s)) (return(s)) { code }

// String() makes book satisfy the fmt.Stringer interface
// when fmt.Println(b) is called, Go checks if b has a String() method
// and uses its return value instead of the default struct formatting {West With The Night}
func (b book) String() string {
	return fmt.Sprint("The title of the book is ", b.title)
}

// count is a NAMED TYPE based on int — you can define methods on any named type,
// not just structs (see Lesson 138 REMARKS)
type count int

// String() makes count satisfy the fmt.Stringer interface too
// strconv.Itoa(int(c)) converts: count → int → string
// the int(c) conversion is needed because Itoa expects an int, not a count
func (c count) String() string {
	return fmt.Sprint("This is the number ", strconv.Itoa(int(c)))
}

func main() {
	b := book{
		title: "West With The Night",
	}

	// count is a named type based on int — we can use int literals directly
	var n count = 42

	// fmt.Println checks if each argument satisfies fmt.Stringer
	// since both book and count have String(), their custom output is used
	fmt.Println(b) // output: The title of the book is West With The Night
	fmt.Println(n) // output: This is the number 42

	// without the String() methods, output would be:
	//   {West With The Night}    (default struct format)
	//   42                       (default int format)
}

/*
REMARKS:
- The fmt.Stringer INTERFACE is defined in the fmt package as:
    type Stringer interface {
        String() string
    }
- Any type with a `String() string` method IMPLICITLY satisfies it (Lesson 139)
- fmt.Println, fmt.Sprintf, fmt.Sprint all check for Stringer — it controls
  how your type appears in formatted output throughout the standard library
- This is POLYMORPHISM from the standard library: fmt.Println accepts `any`,
  but dispatches to String() when the value satisfies fmt.Stringer
- NAMED TYPES like `type count int` can have methods — the underlying type (int)
  cannot, but once you give it a name, it becomes its own type with its own method set
- strconv.Itoa is "Integer to ASCII" — converts an int to its string representation;
  the type conversion int(c) is needed because count and int are DIFFERENT TYPES in Go
  even though count's underlying type is int (Go has strong typing — Lesson 017)
- BUILDING ON LESSON 139: Stringer is a one-method interface — this follows Go's
  philosophy of small, focused interfaces ("the bigger the interface, the weaker
  the abstraction")
- BUILDING ON LESSON 138: the String() method uses a value receiver, just like
  the speak() method did
- BUILDING ON LESSON 017: type conversion int(c) is required because Go does NOT
  do implicit type conversion — even between a named type and its underlying type
- COMMON PATTERN: implement fmt.Stringer on your types for better debugging output,
  logging, and user-facing messages
*/
