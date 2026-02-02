/*
● EXERCISE 53 — Structs with slice fields
● Building on Lesson 125: creating a struct that contains a SLICE as a field
● Combines structs (125) with slices (098) and for-range (074)
● Access slice fields with DOT NOTATION, then range over them
*/
/*
Create your own type "person" which will have an underlying type of "struct" so that it can store
the following data:
● first name
● last name
● favorite ice cream flavors
Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice
which stores the favorite flavors.
*/

package main

import "fmt"

// struct with a SLICE field — []string can hold any number of flavors
// this is a common pattern: structs group related data, slices handle variable-length lists
type person struct {
	first       string
	last        string
	favIceCream []string // a slice inside a struct — combines two composite types
}

func main() {
	// composite literal with a nested slice literal
	p1 := person{
		first:       "Chris",
		last:        "Williams",
		favIceCream: []string{"vanilla", "maple crunch", "walnut"},
	}

	p2 := person{
		first:       "Kim",
		last:        "Fabre",
		favIceCream: []string{"chocolate", "maple walnut", "almond joy"},
	}

	// dot notation accesses struct fields (Lesson 125)
	fmt.Println(p1.first)
	// dot notation + index access: struct field → slice element
	fmt.Println(p2.favIceCream[1]) // "maple walnut"

	// range over a slice FIELD of a struct (Lesson 074 + 125)
	for _, i := range p1.favIceCream {
		fmt.Println(i)
	}

	for _, x := range p2.favIceCream {
		fmt.Println(p2.first, "favorite ice cream is", x)
	}

}

/*
REMARKS:
- BUILDING ON LESSON 125: struct fields can be ANY type — including slices, maps, other structs
- This is the first time we see a COMPOSITE TYPE inside another composite type:
    ● person is a struct (composite)
    ● favIceCream is a []string (composite) inside the struct
- Access pattern chains: p1.favIceCream[1]
    ● p1 → the struct value
    ● .favIceCream → the slice field (dot notation, Lesson 125)
    ● [1] → the second element of that slice (index access, Lesson 100)
- for-range works on struct fields just like standalone slices:
    for _, v := range p1.favIceCream — iterates the slice inside the struct
- Struct fields that are slices are still REFERENCE types (Lesson 106):
    ● If you assign p1 to another variable, the favIceCream slice header is copied
    ● But both still point to the same underlying array
- This pattern (struct + slice field) is extremely common in real Go code:
    ● User with []Role, Order with []Item, Config with []string, etc.
- BUILDING ON LESSON 098: the slice literal []string{...} works the same inside a struct
  as it does standalone — same rules for append, len, range
*/
