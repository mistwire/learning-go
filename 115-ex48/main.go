/*
Create a slice of a slice of string ([][]string). Store the following data in the multi-dimensional
slice:
"James", "Bond", "Shaken, not stirred"
"Miss", "Moneypenny", "I'm 008."
Range over the records, then range over the data in each record.
*/
package main

import (
	"fmt"
)

func main() {

	jb := []string{"James", "Bond", "Shaken, not stirred"}
	mm := []string{"Miss", "Moneypenny", "I'm 008"}

	xxs := [][]string{jb, mm}
	fmt.Println(xxs)

	for i, v := range xxs {
		fmt.Println(i, v)
		for _, b := range v {
			fmt.Println(b)
		}
	}

}

/*
REMARKS:
- [][]string is a slice of slices - a 2D/multi-dimensional data structure
- Each inner slice can have a DIFFERENT length (jagged/ragged array)
- Useful for representing tables, matrices, or grouped data
- Nested range loops: outer loop iterates rows, inner loop iterates columns
- Using _ (blank identifier) discards the index when you only need values
- Each "row" (jb, mm) is an independent slice that can be manipulated separately
- Common pattern for CSV data, database rows, or any tabular information
*/
