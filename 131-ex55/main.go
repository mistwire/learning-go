/*
● EXERCISE 55 — Embedded structs and field promotion
● Building on Lesson 126: practicing the embedding pattern with a real-world example
● Embeds an engine struct inside a vehicle struct
● Demonstrates FIELD PROMOTION: accessing engine's fields directly on vehicle
● Combines: struct definition (125), embedding (126), composite literals (125)
*/
/*
Create a type engine struct, and include this field:
● electric bool
Create a type vehicle struct, and include these fields:
	■ engine (embedded)
	■ make
	■ model
	■ doors
	■ color
Create two VALUES of TYPE vehicle using composite literals.
Print out each of these values.
Print out a single field from each of these values.
*/

package main

import "fmt"

// engine is a small struct — embedding it into vehicle avoids duplicating fields
type engine struct {
	electric bool
}

// vehicle EMBEDS engine (type name only, no field name)
// this promotes engine's fields into vehicle (Lesson 126)
type vehicle struct {
	engine        // embedded struct — electric is promoted to vehicle
	make   string
	model  string
	doors  int
	color  string
}

func main() {
	// composite literal with embedded struct — must use type name as key
	car1 := vehicle{
		engine: engine{
			electric: true, // engine: engine{...} — same pattern as Lesson 126
		},
		make:  "Volkswagen",
		model: "ID.4",
		doors: 4,
		color: "Red",
	}
	car2 := vehicle{
		engine: engine{
			electric: false,
		},
		make:  "Jeep",
		model: "Wrangler",
		doors: 4,
		color: "Grey",
	}
	fmt.Println(car1) // prints entire struct including embedded engine
	fmt.Println(car2)
	fmt.Println(car1.electric) // FIELD PROMOTION: car1.electric instead of car1.engine.electric
	fmt.Println(car2.make)     // regular field access (not promoted — belongs to vehicle directly)

}

/*
REMARKS:
- BUILDING ON LESSON 126: this exercise practices the same embedding pattern with a practical example
- vehicle embeds engine — the "electric" field is PROMOTED to vehicle:
    ● car1.electric works (promoted access — shorter, cleaner)
    ● car1.engine.electric also works (explicit path — always available)
- In the composite literal, you MUST use the type name: engine: engine{...}
    ● You cannot write electric: true directly in the vehicle literal
    ● The embedded struct must be initialized as a nested value
- No shadowing here — vehicle doesn't define its own "electric" field,
  so the promoted field is accessed directly without ambiguity (contrast with Lesson 126)
- fmt.Println on a struct prints all fields including embedded ones:
    ● car1 prints as {{true} Volkswagen ID.4 4 Red}
    ● The {true} is the embedded engine struct
- REAL-WORLD PATTERN: embedding is common for composition in Go
    ● A Vehicle HAS an Engine (composition), not "Vehicle IS an Engine" (inheritance)
    ● Other examples: HTTPServer embedding Logger, Request embedding Context
- BUILDING ON LESSON 125: same composite literal syntax, just with an embedded struct added
*/
