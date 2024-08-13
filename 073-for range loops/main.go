package main

import "fmt"

func main() {
	// for range loop https://go.dev/doc/effective_go#for
	// data structures - slice
	xi := []int{42, 43, 44, 45, 46, 47}
	for i, v := range xi {
		fmt.Printf("ranging over a slice: Index:%v Value:%v\n", i, v)
	}

	// for range loop
	// data structures - map
	// string ⤵️ is the key, int is the value
	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}
	for k, v := range m {
		fmt.Printf("ranging over map: %v is the key, %v is the value\n", k, v)
	}
	for pos, char := range "日本\x80語" { // \x80 is an illegal UTF-8 encoding
		fmt.Printf("character %#U starts at byte position %d\n", char, pos)
	}
}

// note: maps are unordered!
