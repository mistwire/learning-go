/*
● MODULAR CODE — importing and using external packages
● Building on Lesson 033: EXPORTED functions (uppercase) are accessible from other packages
● External packages are imported by their MODULE PATH (e.g., github.com/mistwire/puppy)
● Access exported functions with package.FunctionName() syntax
*/
package main

import (
	"fmt"

	"github.com/mistwire/puppy" // external module — fetched via "go get" or "go mod tidy"
)

func main() {
	// calling EXPORTED functions from the puppy package
	// Bark() and Barks() start with uppercase — they are exported
	s1 := puppy.Bark()
	s2 := puppy.Barks()
	fmt.Println(s1)
	fmt.Println(s2)

	// or call directly in Println — no intermediate variable needed
	fmt.Println(puppy.Bark())
	fmt.Println(puppy.Barks())
}

/*
REMARKS:
- Go uses MODULES for dependency management — defined by go.mod at the project root
- External packages are imported by their full module path (github.com/user/repo)
- "go get github.com/mistwire/puppy" downloads the module and adds it to go.mod
- "go mod tidy" cleans up go.mod and go.sum — adds missing deps, removes unused ones
- BUILDING ON LESSON 033 (scope): only EXPORTED identifiers (uppercase) are accessible
  from outside the package — puppy.Bark() works because Bark starts with "B"
- The import path groups: standard library imports first, then a blank line, then external imports
  (this is enforced by goimports/gofmt)
- You can alias imports: import p "github.com/mistwire/puppy" → p.Bark()
- Unused imports cause a COMPILE ERROR (just like unused variables from Lesson 013)
*/
