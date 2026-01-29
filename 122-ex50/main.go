/*
Using the code from the previous example, add a record to your map. Now print the map out
using the “range” loop
key 			value
`fleming_ian` 	`steaks`, `cigars`, `espionage`
*/

package main

import "fmt"

func main() {
	m := map[string][]string{
		`bond_james`:       {`shaken, not stirred`, `martinis`, `fast cars`},
		`moneypenny_jenny`: {`intelligence`, `literature`, `computer science`},
		`no_dr`:            {`cats`, `ice cream`, `sunsets`},
	}

	m[`fleming_ian`] = []string{`steaks`, `cigars`, `espionage`}

	for k, v := range m {
		fmt.Printf("Key: %v\n", k)
		for _, x := range v {
			fmt.Printf("Slice is: %v\n", x)
		}
	}

}
