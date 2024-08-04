package main

import (
	"fmt"
)

func main() {
	var name string
	var age int

	fmt.Print("Enter your name and age: ")
	n, err := fmt.Scanln(&name, &age)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Scanned %d items. Name: %s, Age: %d\n", n, name, age)
	}
}
