/*
● ANONYMOUS STRUCTS
● Building on Lessons 125-126: instead of declaring a named type with "type",
  you can define a struct INLINE — with no reusable type name
● The struct definition and initialization happen in a SINGLE expression
● Compare the TYPE output of an anonymous struct vs. a named struct using %T
*/
package main

import "fmt"

// named struct type — same as lessons 125 & 126
type person struct {
	first string
	last  string
	age   int
}

func main() {
	// ANONYMOUS STRUCT: the type is defined and a value is created in one statement
	// struct { fields... }{ values... }
	//   ↑ type definition    ↑ composite literal
	person1 := struct {
		first string
		last  string
		age   int
	}{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	// named struct — for comparison
	person2 := person{
		first: "Jenny",
		last:  "Moneypenny",
		age:   32,
	}

	fmt.Printf("%T\n", person1) // struct { first string; last string; age int }  ← anonymous, no type name
	fmt.Printf("%T\n", person2) // main.person  ← named type, includes package prefix
}

/*
REMARKS:
- KEY DIFFERENCE FROM LESSONS 125-126: there is NO "type" declaration — the struct has no name
- An anonymous struct is defined and initialized in a single expression:
    variable := struct{ fields }{ values }
- The %T verb reveals the difference:
    named struct   → "main.person" (package-qualified type name)
    anonymous struct → "struct { first string; last string; age int }" (the full structure is the type)
- Anonymous structs CANNOT be reused — you can't create a second value of the same "type"
  without repeating the entire struct definition
- USE CASES for anonymous structs:
    ● One-off data grouping (e.g., quick test fixtures, template data)
    ● JSON unmarshalling when you only need the shape once
    ● Table-driven tests: []struct{ input int; expected string }{ ... }
- If you need the struct in MORE THAN ONE PLACE, always use a named type (lesson 125)
- Anonymous structs with identical fields ARE assignment-compatible with each other
  and with named structs that share the same field layout (structural typing)
- This is a common pattern in Go — you'll see it frequently in test files
*/
