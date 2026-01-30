/*
● MAP FOR-RANGE — comprehensive map operations: iterate, comma-ok, delete, len
● Building on Lessons 074 & 116: applying for-range to maps
● THREE range forms: keys only, keys+values, values only (with _)
● Also previews: comma-ok idiom (Lesson 119) and delete (Lesson 118)
*/
package main

import "fmt"

func main() {

	m := map[string]int{
		"todd":  42,
		"henry": 16,
	}

	fmt.Println("Henry's age is ", m["henry"])

	// RANGE FORM 1: keys only — one variable
	for k := range m {
		fmt.Printf("just the keys: %s\n", k)
	}

	// RANGE FORM 2: keys AND values — two variables
	for k, v := range m {
		fmt.Printf("%s is %d years old\n", k, v)
	}

	// RANGE FORM 3: values only — discard key with _ (Lesson 013)
	for _, v := range m {
		fmt.Printf("just the values: %d\n", v)
	}

	// COMMA-OK IDIOM with maps (detailed in Lesson 119)
	// v = value (or zero value if missing), ok = bool (key exists?)
	if v, ok := m["Padget"]; ok {
		fmt.Printf("%s is %d years old\n", "Padget", v)
	} else {
		fmt.Printf("%s not found\n", "Padget")
	}

	// ADD a new entry
	m["Shakespeare"] = 459
	fmt.Printf("%#v\n", m)

	// DELETE a key — builtin function (detailed in Lesson 118)
	delete(m, "Shakespeare")
	fmt.Printf("%#v\n", m)
	delete(m, "Shakespeare") // deleting a non-existent key does NOT panic

	// len() returns number of key-value pairs
	fmt.Println("len: ", len(m))

}

/*
REMARKS:
- BUILDING ON LESSON 074: for-range on maps gives (key, value) instead of (index, value)
- THREE forms of map range (same pattern as slice range):
    ● for k := range m — keys only
    ● for k, v := range m — keys and values
    ● for _, v := range m — values only (discard key)
- Map iteration order is RANDOMIZED — you'll get different order each run
  (Go intentionally randomizes to prevent code from depending on order)
- This lesson is a comprehensive overview — the next lessons break out:
    ● Lesson 118: delete() in detail
    ● Lesson 119: comma-ok idiom in detail
- COMMA-OK (preview): m["key"] returns TWO values when used in the v, ok := form
    ● v = the value (or zero value if key is missing)
    ● ok = true if key exists, false if not
- delete(m, key) is a BUILTIN — removes a key-value pair
    ● Deleting a non-existent key is safe — no panic
    ● Unlike slices, maps have a dedicated delete function
- CONTRAST WITH SLICES: maps have delete(); slices don't (use append trick — Lesson 103)
*/
