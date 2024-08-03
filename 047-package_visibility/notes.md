[Good article on package visibility](https://www.digitalocean.com/community/tutorials/understanding-package-visibility-in-go)

If a function is Capitalized, it is exported

```Go:
package logging

import (
	"fmt"
	"time"
)

var debug bool

func Debug(b bool) {
	debug = b
}

func Log(statement string) {
	if !debug {
		return
	}

	fmt.Printf("%s %s\n", time.Now().Format(time.RFC3339), statement)
}
