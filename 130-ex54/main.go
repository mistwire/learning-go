/*
● EXERCISE 54 — Storing structs in a map
● Building on Lesson 128: takes the person struct and stores values in a MAP
● Map key = last name (string), map value = person (struct)
● Combines: structs (125), maps (116), for-range (074), and slice fields (128)
*/
/*
Take the code from the previous exercise, then store the VALUES of type person in a map with
the KEY of last name. Access each value in the map. Print out the values, ranging over the
slice.
*/

package main

import "fmt"

type person struct {
	first       string
	last        string
	favIceCream []string
}

func main() {
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

	// map[string]person — KEY is the last name, VALUE is the entire person struct
	// p1.last evaluates to "Williams" — used as the map key
	m := map[string]person{
		p1.last: p1, // "Williams": person{...}
		p2.last: p2, // "Fabre": person{...}
	}

	// NESTED RANGE: outer loop ranges over the map, inner loop over the slice field
	for _, v := range m {
		fmt.Println(v.first, v.last) // prints the entire person struct
		for _, v2 := range v.favIceCream {
			fmt.Println(v.first, v2) // accesses struct fields from the map value
		}
	}
	for _, p := range m {
    fmt.Printf("Name: %s %s, Favorite Ice Cream: %s", p.first, p.last, p.favIceCream) }
}

/*
REMARKS:
- BUILDING ON LESSON 128: same struct, now stored in a map instead of standalone variables
- map[string]person — the VALUE type is a struct (any type can be a map value)
- Using a STRUCT FIELD as the map key: p1.last evaluates to a string at runtime
    ● The key is the string "Williams", not the expression p1.last
    ● This is a common pattern: index records by a unique field (ID, name, email)
- NESTED RANGING combines two patterns:
    1. for _, v := range m — iterate map values (Lesson 117), v is a person struct
    2. for _, v2 := range v.favIceCream — iterate the slice field of that struct (Lesson 128)
- v.first and v.favIceCream use DOT NOTATION on the map value directly
    ● No need to extract the struct first — range gives you the value, then you access fields
- Map iteration order is RANDOMIZED (Lesson 117):
    ● "Williams" might print before or after "Fabre" on different runs
- BUILDING ON LESSON 116: maps can hold ANY value type — int, string, []string, structs, etc.
- This exercise connects three major topics together:
    ● Structs for grouping data (Lesson 125)
    ● Maps for key-based lookup (Lesson 116)
    ● Slices for variable-length lists within structs (Lesson 098)
- In real Go code, map[string]SomeStruct is one of the most common patterns
  (e.g., users by ID, config by name, cache by key)
*/