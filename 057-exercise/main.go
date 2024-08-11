// Using the code you wrote in the previous hands-on exercise ( -o to customize name of build):
// ● build your program for windows
// 	○ GOOS=windows GOARCH=amd64 go build
// ● build your program for mac (apple silicon)
// 	○ GOOS=darwin GOARCH=arm64 go build -o myapp-macos-arm64
// ● build your program for linux
// 	○ GOOS=linux GOARCH=amd64 go build -o myapp-linux

package main

import "fmt"

func main() {
	fmt.Println("Will code for 🍕")
}
