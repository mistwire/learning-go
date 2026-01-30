/*
● DEPENDENCY MANAGEMENT — transitive (nested) dependencies
● Building on Lesson 049: the puppy module now depends on ANOTHER module (dog)
● Go modules handle TRANSITIVE DEPENDENCIES automatically via go.sum
● Your code only imports puppy, but puppy internally imports dog
*/
package main

import (
	"fmt"

	"github.com/mistwire/puppy" // puppy now has its own dependency on a "dog" module
)

func main() {
	s1 := puppy.Bark()
	s2 := puppy.Barks()
	fmt.Println(s1)
	fmt.Println(s2)

	// or like this:
	fmt.Println(puppy.Bark())
	fmt.Println(puppy.Barks())

	// BigBark/BigBarks — these functions in puppy internally call the dog module
	// you don't import dog directly — puppy handles that dependency
	s3 := puppy.BigBark()
	s4 := puppy.BigBarks()
	fmt.Println(s3)
	fmt.Println(s4)

}

/*
REMARKS:
- BUILDING ON LESSON 049: same puppy package, but now it has its own dependencies
- TRANSITIVE DEPENDENCIES: puppy depends on dog → when you "go get" puppy, Go also fetches dog
- go.sum file records the CRYPTOGRAPHIC HASHES of all dependencies (direct + transitive)
  — ensures reproducible builds and detects tampering
- go.mod lists your DIRECT dependencies; go.sum lists ALL dependencies (direct + indirect)
- You never import transitive deps directly — that's the dependency's responsibility
- "go mod tidy" resolves the full dependency tree and updates go.mod + go.sum
- "go mod graph" shows the full dependency tree (who depends on whom)
- This is Go's answer to package managers like npm/pip — but built into the language toolchain
*/
