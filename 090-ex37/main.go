// LESSON: Exercise - map lookup with comma-ok idiom
/*
● use the code from the previous exercise
● add this code to print a single value stored in the map
age := m["James"]
fmt.Println(age)
● now use similar code to use the lookup of "Q" and print that value
● now use the "comma ok" idiom to test whether "Q" is actually stored in the map, then
print out a statement if it is not stored in the map
○ hint: check effective go for the "comma ok" idiom
*/

package main

import (
	"fmt"
)

func main() {
	// xi := []int{42, 43, 44, 45, 46, 47}

	// for i, v := range xi {
	// 	fmt.Printf("For index %v: the value is %v\n", i, v)
	// }
	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}
	for k, v := range m {
		fmt.Printf("For key %v: the value is %v\n", k, v)
	}
	james_age := m["James"]
	fmt.Println(james_age)

	if q_age, ok := m["Q"]; ok {
		fmt.Println(q_age)
	} else {
		fmt.Println("unknown record:", "Q")
	}
}
