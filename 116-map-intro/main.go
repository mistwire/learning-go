/*
● MAP INTRO — key-value data structure (like dictionaries/hash maps in other languages)
● Syntax: map[keyType]valueType — e.g., map[string]int
● Two ways to create: composite literal or make()
● Access values by key: m["key"]
● len() returns number of key-value pairs
*/
package main

import "fmt"

func main() {
	// MAP LITERAL: map[key type]value type{ key: value, ... }
	// trailing comma required on last entry (multi-line)
	am := map[string]int{
		"Todd":   42,
		"Henry":  16,
		"Padget": 14,
	}

	// access by key — returns the value associated with that key
	fmt.Println("The age of Henry was", am["Henry"])

	fmt.Println("------Println------")
	fmt.Println(am)
	fmt.Println("------Printf------")
	fmt.Printf("%%v:\t%v\n", am)   // default format
	fmt.Printf("%%#v:\t%#v\n", am) // Go-syntax representation

	// MAKE: creates an empty map — then add entries with m[key] = value
	an := make(map[string]int)
	fmt.Println(an) // map[] (empty)
	an["Lucas"] = 28
	an["Steph"] = 37
	fmt.Println(an)
	fmt.Printf("%%v: %#v\n", an)

	// len() returns the number of key-value pairs
	fmt.Println(len(an)) // 2
}

/*
REMARKS:
- Maps are Go's HASH TABLE / dictionary — unordered key-value storage
- TWO ways to create a map:
    1. Literal: m := map[string]int{"key": 42}
    2. Make: m := make(map[string]int) then m["key"] = 42
- KEY TYPES must be COMPARABLE (support ==): string, int, bool, structs (if all fields are comparable)
    ● Slices, maps, and functions CANNOT be keys
- VALUE TYPES can be anything — including slices, maps, structs, functions
- Accessing a MISSING KEY returns the ZERO VALUE for the value type (0 for int, "" for string)
    ● This does NOT panic — but you can't tell if the key exists or the value is actually zero
    ● Use the comma-ok idiom to distinguish (Lesson 119)
- Maps are REFERENCE types (like slices) — passing a map to a function lets it modify the original
- Map iteration order is NOT guaranteed (intentionally randomized — Lesson 074)
- Zero value of a map is nil — you MUST initialize with make() or a literal before writing to it
    ● Reading from a nil map returns zero values; WRITING to a nil map PANICS
- BUILDING ON LESSON 098: maps are another composite type alongside slices
  — both use make(), both are reference types, both work with for-range
*/