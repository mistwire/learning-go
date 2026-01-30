/*
● MAP DELETE — removing entries and safely accessing missing keys
● Building on Lessons 116-117: delete() is a BUILTIN for removing map entries
● Deleting a non-existent key does NOT panic — it's a no-op
● Accessing a non-existent key returns the ZERO VALUE — also no panic
*/
package main

import "fmt"

func main() {
	am := map[string]int{
		"Todd":   42,
		"Henry":  16,
		"Padget": 14,
	}

	fmt.Println("The age of Henry was", am["Henry"])

	fmt.Println(am)
	fmt.Printf("%#v\n", am)

	an := make(map[string]int)
	an["Lucas"] = 28
	an["Steph"] = 37
	an["George"] = 78
	fmt.Println(an)
	fmt.Printf("%#v\n", an)
	fmt.Println(len(an)) // 3

	// DELETE: delete(map, key) — removes the key-value pair
	delete(an, "George")

	fmt.Println("--- accessing keys that don't exist")
	delete(an, "Georage")      // key doesn't exist — no panic, no error, just a no-op
	fmt.Println(an["Georgey"]) // key doesn't exist — returns 0 (zero value for int), no panic
	fmt.Println("------------------------")
	fmt.Println(an) // only Lucas and Steph remain
	fmt.Println("------------------------")

}

/*
REMARKS:
- delete(map, key) is a BUILTIN function — only works with maps (not slices)
- SAFE OPERATIONS on maps — these will NEVER panic:
    ● delete(m, missingKey) — silently does nothing
    ● m[missingKey] — returns the zero value for the value type (0, "", false, nil)
- UNSAFE OPERATION — this WILL panic:
    ● Writing to a nil map: var m map[string]int; m["key"] = 42 — PANIC
    ● Always initialize with make() or a literal first
- CONTRAST WITH LESSON 103 (slice delete):
    ● Maps: delete(m, key) — simple, O(1), built-in
    ● Slices: xs = append(xs[:i], xs[i+1:]...) — manual, O(n), requires index not key
- The zero-value-on-missing-key behavior is WHY the comma-ok idiom (Lesson 119) exists
    ● m["key"] returns 0 — but is the key missing, or is the value actually 0?
    ● v, ok := m["key"] — ok tells you definitively if the key exists
- BUILDING ON LESSON 116: after delete, len(m) decreases by 1
*/
