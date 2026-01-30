/*
● Introduction to STRUCTS in Go
● A struct is a composite/aggregate data type that groups together
  zero or more fields with defined types under a single name
● Define a custom TYPE using the "type" keyword with a struct literal
● Create instances of the struct using a COMPOSITE LITERAL
● Access individual fields using DOT NOTATION
*/
package main

import "fmt"

// type declaration: "person" is a new type whose underlying type is a struct
// fields are listed with name followed by type
type person struct {
	first string
	last  string
	age   int
}

func main() {
	// creating a value of type "person" using a composite literal
	// fields are set using field:value syntax
	person1 := person{
		first: "James",
		last:  "Bond",
		age:   32, // trailing comma is REQUIRED on the last field when using multi-line syntax
	}

	person2 := person{
		first: "Jenny",
		last:  "Moneypenny",
		age:   32,
	}

	fmt.Println(person1)        // prints all fields: {James Bond 32}
	fmt.Println(person2)        // prints all fields: {Jenny Moneypenny 32}
	fmt.Printf("Struct type:%T\n%#v\n%v\n", person1, person2, person1)
	fmt.Println(person1.first)  // dot notation to access a single field
	fmt.Println(person2.last)   // dot notation to access a single field
	fmt.Println(person1.first, person1.last, person1.age)
}

/*
REMARKS:
- A STRUCT is Go's way of creating custom types that group related data together (like a class without methods)
- Defined at the PACKAGE LEVEL using: type <name> struct { fields... }
- The "type" keyword creates a new named type; the struct literal defines its structure
- Fields are accessed using DOT NOTATION: person1.first, person1.age, etc.
- Structs are VALUE types - assigning one struct to another COPIES all the fields
- You can also create a struct without field names: person{"James", "Bond", 32} (positional/ordered)
  but this is fragile and discouraged - always use field:value syntax for clarity
- Zero value of a struct has all fields set to their respective zero values ("" for string, 0 for int)
- Structs are comparable if all their fields are comparable (==, !=)
- Exported structs/fields start with an UPPERCASE letter; lowercase = unexported (package-private)
- Structs are foundational in Go - they are the basis for methods, interfaces, and composition
*/
