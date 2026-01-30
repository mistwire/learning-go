/*
● GIT TAGS — versioning Go modules with semantic version tags
● Building on Lessons 049-050: modules need VERSION NUMBERS for dependency management
● Go uses SEMANTIC VERSIONING (semver): vMAJOR.MINOR.PATCH (e.g., v1.2.3)
● Git tags (e.g., "git tag v1.0.0") mark a commit as a release version
*/
package main

import "fmt"

func main() {
	fmt.Println("Hello, we're working on versions")
}

/*
REMARKS:
- Go modules use SEMANTIC VERSIONING: vMAJOR.MINOR.PATCH
    ● MAJOR: breaking changes (v1 → v2 changes the import path!)
    ● MINOR: new features, backwards compatible (v1.1.0 → v1.2.0)
    ● PATCH: bug fixes, backwards compatible (v1.2.0 → v1.2.1)
- To publish a version: "git tag v1.0.0" then "git push origin v1.0.0"
- Go uses the git tag to resolve "go get github.com/user/repo@v1.0.0"
- Without a tag, Go uses a PSEUDO-VERSION based on the commit hash
  (e.g., v0.0.0-20240101120000-abcdef123456)
- MAJOR VERSION 2+: the import path MUST include the major version
  e.g., "github.com/user/repo/v2" — this is Go's way of handling breaking changes
- BUILDING ON LESSON 050: go.mod records the exact version of each dependency
  git tags are what make those version numbers meaningful
- "go list -m -versions github.com/user/repo" shows all available tagged versions
*/
