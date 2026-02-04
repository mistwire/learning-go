/*
● METHODS — a method is a FUNC with a RECEIVER; it attaches behavior to a TYPE
● The RECEIVER binds the function to a specific type, so you call it with dot notation: `value.Method()`
● Method syntax follows the same pattern as regular funcs with an added receiver:
	○ func (r ReceiverType) identifier(params) (returns) { code }
● Go does NOT have classes — methods on types are how you associate behavior with data
● BUILDING ON LESSON 125: STRUCTS define the data; METHODS define the behavior on that data
● BUILDING ON LESSON 134: same func signature pattern, with the RECEIVER added before the name
*/
package main

import "fmt"

// person is a STRUCT type — see Lesson 125 for struct basics
// methods let us attach behavior (like speak) to this type
type person struct {
	first string
}

// func (r receiver) identifier(p parameter(s)) (return(s)) { code }

// speak is a METHOD on the person type
// (p person) is the VALUE RECEIVER — p is a COPY of the person value (pass by value!)
// any person value can call this method: p1.speak()
func (p person) speak() {
	// p.first accesses the struct field — same dot notation as Lesson 125
	fmt.Println("Hello, my name is", p.first)
}

func main() {
	// create two person values — see Lesson 125 for struct composite literals
	p1 := person{
		first: "James",
	}
	p2 := person{
		first: "Jenny",
	}

	// call the speak METHOD using dot notation on each person value
	// this is like calling speak(p1) but with the receiver syntax instead
	p1.speak() // output: Hello, my name is James
	p2.speak() // output: Hello, my name is Jenny
}

/*
REMARKS:
- A METHOD is just a function with a RECEIVER — the receiver ties it to a type
- VALUE RECEIVER `(p person)` gets a COPY of the value — changes to p inside the
  method do NOT affect the original (everything in Go is PASS BY VALUE — Lesson 134)
- POINTER RECEIVER `(p *person)` gets a pointer to the original value — use this when
  you need to MODIFY the receiver or avoid copying large structs (covered in a later lesson)
- You can define methods on ANY NAMED TYPE, not just structs:
    type myInt int
    func (m myInt) double() myInt { return m * 2 }
- Go does NOT have classes or inheritance — instead it uses TYPES + METHODS + COMPOSITION
  (embedded structs from Lesson 126) to achieve similar goals
- BUILDING ON LESSON 125: structs define fields (data), methods define behavior — together
  they give you something similar to objects in other languages
- BUILDING ON LESSON 134: the func signature is identical except for the receiver;
  compare: `func speak()` vs `func (p person) speak()`
- BUILDING ON LESSON 126: methods are PROMOTED through embedded structs — if type A embeds
  type B, then A gets all of B's methods automatically (composition over inheritance)
- KEY IDIOM: use short receiver names (one or two letters), typically the first letter(s)
  of the type name: `p` for person, `s` for server — this is idiomatic Go
*/
