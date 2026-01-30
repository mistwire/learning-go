/*
● EMBEDDED STRUCTS (composition / "has-a" relationship)
● Building on Lesson 125: instead of duplicating fields, EMBED one struct inside another
● Go does NOT have inheritance — it uses COMPOSITION via embedding
● An embedded struct's fields are PROMOTED to the outer struct
● If the outer struct declares a field with the SAME NAME as a promoted field,
  the outer field SHADOWS (overrides) the inner one
*/
package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

// secretAgent EMBEDS person (no field name — just the type)
// this is different from a named field like "p person"
// embedding promotes person's fields (first, last, age) into secretAgent
type secretAgent struct {
	person       // embedded struct — note: type name only, no field name
	l2k   bool   // new field specific to secretAgent
	first string // SHADOWS the promoted "first" from person
}

func main() {
	// when initializing, the embedded struct must be set using its TYPE NAME as the key
	person1 := secretAgent{
		person: person{ // "person:" is required — you can't inline the embedded fields
			first: "James",
			last:  "Bond",
			age:   32,
		},
		first: "Ethan Hawk", // this sets secretAgent.first (the OUTER field)
		l2k:   true,
	}

	// a plain person value — same as lesson 125
	person2 := person{
		first: "Jenny",
		last:  "Moneypenny",
		age:   32,
	}

	fmt.Println(person1.person, person2) // person1.person accesses the entire embedded struct
	fmt.Println(person1.l2k)             // access secretAgent's own field directly
	fmt.Println(person1.first, person1.person.first)
	// person1.first        → "Ethan Hawk" (outer field SHADOWS the promoted one)
	// person1.person.first → "James"      (explicit path to the INNER/embedded field)
	// person1.last         → "Bond"       (promoted — no shadowing, so direct access works)
}

/*
REMARKS:
- KEY DIFFERENCE FROM LESSON 125: instead of copying fields into a new struct, you EMBED the struct
- Embedding uses just the type name (person) rather than a named field (p person)
- PROMOTED FIELDS: the embedded struct's fields become accessible as if they belong to the outer struct
  e.g., person1.last works even though "last" is defined on person, not secretAgent
- SHADOWING: if the outer struct defines a field with the same name as a promoted field,
  the outer field wins on direct access (person1.first → "Ethan Hawk")
  but the inner field is still reachable via the explicit path (person1.person.first → "James")
- In the composite literal, you MUST use the type name as the key: person: person{...}
  you cannot flatten the embedded fields into the outer literal
- Embedding is NOT inheritance — there is no polymorphism or "is-a" relationship
  secretAgent HAS a person, it is not a subtype of person
- You can embed multiple structs; if promoted field names collide, you must use the explicit path
- Embedding also promotes METHODS — if person had methods, secretAgent would gain them too
  (this becomes important when working with interfaces)
- This pattern is Go's idiomatic alternative to class inheritance in other languages
*/
