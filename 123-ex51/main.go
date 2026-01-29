/*
Using the code from the previous example, delete a record from your map. Now print the map
out using the “range” loop
*/

package main

import "fmt"

func main() {
	m := map[string][]string{
		`bond_james`:       {`shaken, not stirred`, `martinis`, `fast cars`},
		`moneypenny_jenny`: {`intelligence`, `literature`, `computer science`},
		`no_dr`:            {`cats`, `ice cream`, `sunsets`},
	}
	// add to map
	m[`fleming_ian`] = []string{`steaks`, `cigars`, `espionage`}

	// delete from map
	delete(m, "no_dr")

	// fmt.Println(m)
	for k, v := range m {
		fmt.Printf("Key: %v\n", k)
		for _, x := range v {
			fmt.Printf("Slice is: %v\n", x)
		}
	}

}