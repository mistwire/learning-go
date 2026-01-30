/*
● MAP COMMA-OK IDIOM — checking if a key exists in a map
● Building on Lessons 069 & 118: the comma-ok idiom from if-statements applied to maps
● v, ok := m[key] — ok is true if key exists, false if not
● Solves the ambiguity: is the value 0 because the key is missing, or because it's actually 0?
*/
package main

import "fmt"

func main() {
	an := make(map[string]int)
	an["Lucas"] = 28
	an["Steph"] = 37
	an["George"] = 78
	fmt.Println(an)

	delete(an, "George") // remove George (Lesson 118)

	// COMMA-OK — two-value form of map access
	// v = value (or zero value if missing), ok = boolean (key exists?)
	v, ok := an["Georgey"] // Georgey doesn't exist → v=0, ok=false
	if ok {
		fmt.Println("The value prints: ", v)
	} else {
		fmt.Println("Key doesn't exit")
	}

	// IDIOMATIC FORM: combine init statement + condition (Lesson 069)
	// x and ok are scoped to the if/else block
	if x, ok := an["Lucas"]; !ok {
		fmt.Println("Key doesn't exit")
	} else {
		fmt.Println("The value prints: ", x) // 28
	}

}

/*
REMARKS:
- THIS TIES TOGETHER two earlier patterns:
    1. Comma-ok / init-statement in if (Lesson 069)
    2. Map zero-value ambiguity (Lesson 118)
- WITHOUT comma-ok: m["missing"] returns 0 — you can't tell if the key exists
- WITH comma-ok: v, ok := m["missing"] — ok is false, v is the zero value
    ● ok == true: key exists, v is the real value
    ● ok == false: key does NOT exist, v is the zero value (unreliable)
- IDIOMATIC PATTERN (most common in real Go code):
    if v, ok := m[key]; ok {
        // key exists, use v
    }
- This pattern appears EVERYWHERE in Go:
    ● Map lookups: v, ok := m[key]
    ● Type assertions: v, ok := i.(string)
    ● Channel receives: v, ok := <-ch (ok is false when channel is closed)
- BUILDING ON LESSON 069: the semicolon separates the init statement from the condition
  — v and ok are scoped to the if/else block (Lesson 033)
- BUILDING ON LESSON 118: this is the proper way to handle missing keys
  instead of relying on zero values
- This is the LAST map lesson before exercises — the complete map toolkit is:
    Create (116) → Iterate (117) → Delete (118) → Check existence (119)
*/
