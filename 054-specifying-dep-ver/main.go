/*
● SPECIFYING DEPENDENCY VERSIONS — pinning to a specific module version
● Building on Lesson 052: using git tags to request a SPECIFIC version of a dependency
● "go get github.com/mistwire/puppy@v1.3.0" pins to v1.3.0
● Each version may add new exported functions (From11, From12, From13)
*/
package main

import (
	"fmt"

	"github.com/mistwire/puppy"
)

func main() {
	// these functions were added in different versions of the puppy module
	puppy.From11() // added in v1.1.0
	puppy.From12() // added in v1.2.0
	puppy.From13() // added in v1.3.0

	s1 := puppy.Bark()
	s2 := puppy.Barks()
	fmt.Println(s1)
	fmt.Println(s2)

	//also like
	fmt.Println(puppy.Bark())
	fmt.Println(puppy.Barks())

	// now puppy calls the dog module:
	s3 := puppy.BigBark()
	s4 := puppy.BigBarks()
	fmt.Println(s3)
	fmt.Println(s4)

}

/*
REMARKS:
- BUILDING ON LESSONS 049-052: now we're selecting WHICH version of a dependency to use
- Pin a version: "go get github.com/mistwire/puppy@v1.3.0"
  — updates go.mod to require that exact version
- Upgrade: "go get github.com/mistwire/puppy@latest" gets the newest tagged version
- Downgrade: "go get github.com/mistwire/puppy@v1.1.0" rolls back
- Go follows the MINIMUM VERSION SELECTION (MVS) algorithm:
  it always picks the MINIMUM version that satisfies all requirements
  (opposite of npm which picks the latest compatible version)
- If you use From13() but pin to v1.1.0, the code WON'T COMPILE — the function doesn't exist yet
  — this is the compile-time safety that Go provides
- "go mod tidy" after changing versions cleans up go.sum
- The go.mod "require" directive records the pinned version:
    require github.com/mistwire/puppy v1.3.0
*/
