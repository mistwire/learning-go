package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Hello Gophers again! 🥰")
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
} 

// go run main.go - runs the code
// go build - builds an executable of the code (uses the folder name to make the exe?)
// go env GOARCH GOOS - gives you the architecture & OS of the computer you are on
// 
