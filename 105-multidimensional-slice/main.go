/*
● MULTIDIMENSIONAL SLICES — slice of slices (2D data structures)
● Building on Lesson 098: [][]string is a slice where each element is itself a []string
● Each inner slice can have a DIFFERENT length (jagged/ragged structure)
● Useful for tables, grids, grouped records, CSV data
*/
package main

import "fmt"

func main() {
	// two independent slices — each is a "row" of data
	jb := []string{"James", "Bond", "Martini", "Chocolate"}
	jm := []string{"Jenny", "Moneypenny", "Guiness", "Strawberry"}
	fmt.Println(jb)
	fmt.Println(jm)

	// [][]string — a slice of slices: stores both "rows" in a single structure
	xxs := [][]string{jb, jm}
	fmt.Println(xxs)
	// access: xxs[0] = jb, xxs[1] = jm
	// xxs[0][1] = "Bond" (row 0, column 1)
}

/*
REMARKS:
- [][]string is read as "a slice of (slices of string)" — two levels of []
- Each inner slice is INDEPENDENT — they can have different lengths
  (unlike fixed 2D arrays in C where all rows must be the same size)
- Access pattern: xxs[row][col] — first index selects the inner slice, second selects the element
- To iterate, use NESTED for-range (Lesson 073):
    for i, row := range xxs {
        for j, val := range row { ... }
    }
- Each inner slice is still a slice header — modifying jb after creating xxs
  would affect xxs[0] too (they share the same underlying array — Lesson 106)
- BUILDING ON LESSON 098: this is slices all the way down — same rules apply
  (append, slicing, len, cap all work on each level)
- Common uses:
    ● CSV/tabular data: [][]string
    ● Matrices: [][]float64
    ● Grouped results: [][]int
- This connects to embedded structs (Lesson 126) — structs are often a better choice
  when the "columns" have different types or named fields
*/
