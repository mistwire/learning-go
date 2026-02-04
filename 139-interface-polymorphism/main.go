/*
● INTERFACES & POLYMORPHISM — an INTERFACE defines a set of METHOD SIGNATURES;
  any type that implements those methods IMPLICITLY satisfies the interface
● POLYMORPHISM — a single function can accept DIFFERENT TYPES as long as they
  satisfy the same interface; the behavior changes based on the concrete type
● Go interfaces are satisfied IMPLICITLY — there is no `implements` keyword
	○ If a type has the right methods, it IS that interface type — automatically
● BUILDING ON LESSON 138: methods attach behavior to types; interfaces GROUP
  types by shared behavior
● BUILDING ON LESSON 125-126: structs define data, embedded structs enable
  composition; interfaces complete the picture by enabling polymorphism
*/
package main

import "fmt"

// person is a basic struct — see Lesson 125
type person struct {
	first string
}

// secretAgent EMBEDS person — field promotion gives it person's fields (Lesson 126)
// it also adds its own field: ltk (license to kill)
type secretAgent struct {
	person     // EMBEDDED STRUCT — sa.first is promoted from person
	ltk  bool
}

// func (r receiver) identifier(p parameter(s)) (return(s)) { code }

// speak is a METHOD on person — see Lesson 138 for method basics
// person satisfies the human interface because it has a speak() method
func (p person) speak() {
	fmt.Println("I am", p.first)
}

// speak is a METHOD on secretAgent — this is a DIFFERENT method than person.speak()
// secretAgent satisfies the human interface independently with its OWN speak()
// this SHADOWS the promoted person.speak() method (see Lesson 126 for shadowing)
func (sa secretAgent) speak() {
	fmt.Println("I am a secret agent", sa.first)
}

// human is an INTERFACE — it declares a METHOD SET (just one method here: speak)
// ANY type that has a speak() method automatically satisfies this interface
// there is NO explicit "implements" declaration — this is IMPLICIT satisfaction
type human interface {
	speak()
}

// saySomething accepts ANY value that satisfies the human interface
// this is POLYMORPHISM — the same function works with person AND secretAgent
// the actual behavior depends on the CONCRETE TYPE passed in at runtime
func saySomething(h human) {
	h.speak() // calls the speak() method of whatever concrete type h holds
}

func main() {
	// create a secretAgent — note the embedded person struct syntax (Lesson 126)
	sa1 := secretAgent{
		person: person{
			first: "James",
		},
		ltk: true,
	}

	// create a person
	p2 := person{
		first: "Jenny",
	}

	// POLYMORPHISM in action:
	// both sa1 and p2 satisfy the human interface, so both can be passed
	// to saySomething — but each calls its OWN speak() method
	saySomething(sa1) // output: I am a secret agent James
	saySomething(p2)  // output: I am Jenny
}

/*
REMARKS:
- An INTERFACE is a type that defines a METHOD SET — a list of method signatures
- Any type that has ALL the methods in the interface IMPLICITLY satisfies it
  (no `implements` keyword — this is a key difference from Java/C#)
- POLYMORPHISM means one function (saySomething) works with multiple types (person,
  secretAgent) — the concrete type determines which speak() method runs
- The EMPTY INTERFACE `interface{}` (or `any` in Go 1.18+) has zero methods, so
  EVERY type satisfies it — that's why fmt.Println accepts `any` arguments
- Interface values hold two things internally: a concrete TYPE and a concrete VALUE;
  this is sometimes called a "fat pointer"
- BUILDING ON LESSON 138: methods make a type satisfy an interface; without methods,
  interfaces wouldn't work
- BUILDING ON LESSON 126: secretAgent embeds person, which promotes person's fields
  AND methods — but secretAgent defines its OWN speak(), which SHADOWS the promoted one
- BUILDING ON LESSON 125: structs = data, methods = behavior, interfaces = polymorphism;
  together these three concepts are Go's alternative to class-based OOP
- Go philosophy: "The bigger the interface, the weaker the abstraction" (Rob Pike)
  — prefer SMALL interfaces with one or two methods (e.g., io.Reader, io.Writer)
- KEY PATTERN: accept interfaces, return structs — this makes your functions flexible
  about what they receive but specific about what they return
*/
