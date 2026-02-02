/*
● EXERCISE 56 — Anonymous structs with composite field types
● Building on Lesson 127: anonymous struct, now with MAP and SLICE fields
● Combines: anonymous structs (127), maps (116), slices (098), for-range (074)
● Demonstrates accessing nested composite types within an anonymous struct
● Shows that anonymous structs follow all the same field rules as named structs

Create and use an anonymous struct with these fields:
● first string
● friends map[string]int
● favDrinks []string
Print things.
*/
package main

import "fmt"

func main() {
	// anonymous struct — declared and initialized in one expression (Lesson 127)
	// unlike named structs, this type cannot be reused elsewhere
	person1 := struct {
		first     string
		friends   map[string]string // map field — composite type inside a composite type
		favDrinks []string          // slice field — same pattern as Lesson 128 but anonymous
	}{
		first: "Chris",
		friends: map[string]string{ // map literal nested inside struct literal
			"Kim":          "Schmoopy",
			"Mikey":        "Mycock",
			"Mr. Anderson": "Fishy",
		},
		favDrinks: []string{ // slice literal nested inside struct literal (like Lesson 128)
			"Asahi",
			"Green Tea",
			"Coffee",
		},
	}
	fmt.Println(person1.first)            // dot notation on anonymous struct — same as named
	fmt.Println(person1.favDrinks[2])     // struct field → slice index: "Coffee"
	fmt.Println(person1.friends["Mikey"]) // struct field → map key: "Mycock"
	fmt.Printf("%v thinks %v is delicsion and loves %v\n",
		person1.first, person1.favDrinks[0], person1.friends["Kim"])

	// range over a MAP field (Lesson 117 + 127)
	for k, v := range person1.friends {
		fmt.Printf("%v is friends with %v, also known as %v\n", person1.first, k, v)
	}

	// range over a SLICE field (Lesson 074 + 128)
	for _, v := range person1.favDrinks {
		fmt.Printf("%v loves %v\n", person1.first, v)
	}

}

/*
REMARKS:
- BUILDING ON LESSON 127: same anonymous struct pattern, now with COMPLEX field types
- BUILDING ON LESSON 128: same struct-with-slice pattern, but the struct itself is anonymous
- This exercise combines FOUR composite concepts in one expression:
    1. Anonymous struct (Lesson 127) — the struct type itself
    2. Map literal (Lesson 116) — friends map[string]string
    3. Slice literal (Lesson 098) — favDrinks []string
    4. For-range (Lesson 074) — iterating over both map and slice fields
- Access pattern chains work the same as named structs:
    ● person1.favDrinks[2] — struct → slice field → index
    ● person1.friends["Mikey"] — struct → map field → key lookup
- Anonymous structs are useful for ONE-OFF data structures:
    ● Test data, JSON unmarshaling, template data, quick groupings
    ● If you need to reuse the type, promote it to a named struct (Lesson 125)
- The map field iteration order is RANDOMIZED (Lesson 117):
    ● Friends will print in a different order on each run
- Note: the exercise prompt says map[string]int but the implementation uses map[string]string
    ● This is fine — the concept is the same regardless of value type
- KEY TAKEAWAY: anonymous structs follow ALL the same rules as named structs
    ● Dot notation, composite field access, range loops — everything works identically
    ● The only difference is that the type has no name and cannot be referenced elsewhere
*/