// Using the code from the previous hands-on exercise:
// 	look in the $GOPATH/bin
// 		○ launch another terminal
// 		○ see the "go path" environment variable:
// 			■ go env GOPATH
// 		○ navigate to the $GOPATH/bin folder
// 		○ ls -la
// 	● "go install" your program (on the other terminal)
// 		○ look at the executable $GOPATH/bin
// 		○ ls -la
// 	● run the executable in the $GOPATH/bin
// 	● remove the executable in the $GOPATH/bin
// 		○ if you accidentally delete everything, you will need to reinstall your tooling in VS code
// 		○ if you messed it all up, reinstall go

package main

import "fmt"

func main() {
	fmt.Println("Will code for 🍕")
}
