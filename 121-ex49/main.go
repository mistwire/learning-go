/*
Create a map with a key of TYPE string which is a person’s “last_first” name, and a value of
TYPE []string which stores their favorite things. Store three records in your map. Print out all of
the values, along with their index position in the slice.
key 				value
`bond_james` 		`shaken, not stirred`, `martinis`, `fast cars`
`moneypenny_jenny` 	`intelligence`, `literature`, `computer science`
`no_dr` 			`cats`, `ice cream`, `sunsets`
*/

package main

import "fmt"

func main() {
	m := map[string][]string{
		`bond_james`:       {`shaken, not stirred`, `martinis`, `fast cars`},
		`moneypenny_jenny`: {`intelligence`, `literature`, `computer science`},
		`no_dr`:            {`cats`, `ice cream`, `sunsets`},
	}

	for k, v := range m {
		fmt.Printf("Key: %v\n", k)
		for _, x := range v {
			fmt.Printf("Slice is: %v\n", x)
		}
	}

}
