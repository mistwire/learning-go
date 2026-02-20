/*
● io.Writer AS A METHOD PARAMETER — writeOut accepts any io.Writer, decoupling it from any specific destination
● POLYMORPHISM — the same method can write to a file, a buffer, stdout, or any other io.Writer
● BUILDING ON LESSON 142-143: those lessons used io.Writer in standalone functions; here it's a method on a struct
● BUILDING ON LESSON 139: the interface is satisfied implicitly — no "implements" keyword needed
*/
package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

type person struct {
	first string
}

// writeOut accepts any io.Writer — this is "program to the interface, not the implementation"
// The method doesn't care whether w is a file, a buffer, or stdout; it just calls Write
// func (r receiver) identifier(p parameter(s)) (return(s)) { code }
func (p person) writeOut(w io.Writer) {
	// Write expects []byte, so we convert the string with []byte()
	// See Lesson 017 for type conversions
	w.Write([]byte(p.first))
}

func main() {
	p := person{
		first: "Kim",
	}

	// os.Create returns (*os.File, error) — *os.File implements io.Writer
	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatalf("error %s", err)
	}
	defer f.Close() // defer ensures cleanup even if the function panics; see Lesson 137

	// bytes.Buffer is an in-memory buffer that also implements io.Writer
	// Declared as a value type — zero value is an empty, ready-to-use buffer
	var b bytes.Buffer

	p.writeOut(f)  // writes "Kim" to output.txt on disk
	p.writeOut(&b) // writes "Kim" into the in-memory buffer; & needed because Write has a pointer receiver
	fmt.Println(b.String()) // retrieve buffer contents as a string and print
}

/*
REMARKS:
- io.Writer POLYMORPHISM: writeOut works with *os.File, bytes.Buffer, os.Stdout, or any future io.Writer — zero code changes needed
- WHY &b: bytes.Buffer's Write method has a pointer receiver, so we must pass &b, not b, to satisfy io.Writer
- BUILDING ON LESSON 141: "accept interfaces, return structs" — writeOut accepts the io.Writer interface, keeping it decoupled
- BUILDING ON LESSON 137: defer f.Close() is the idiomatic cleanup pattern — place it immediately after the resource is created
- KEY INSIGHT: the caller decides the destination; the method only knows about the io.Writer contract
- Go philosophy: small, composable interfaces (io.Writer has exactly ONE method) enable wide reuse
*/
