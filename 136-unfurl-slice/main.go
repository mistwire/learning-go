/*
● UNFURLING A SLICE — passing a slice into a VARIADIC function using `...` after the argument
● The `...` operator has TWO roles in Go:
	○ In a PARAMETER list: `func foo(vals ...int)` — collects args into a slice (Lesson 135)
	○ In a FUNCTION CALL: `foo(mySlice...)` — UNFURLS a slice back into individual arguments
● BUILDING ON LESSON 135: Lesson 135 passed individual ints to a variadic func;
  here we start with a SLICE and unfurl it so the variadic func can accept it
*/
package main

import "fmt"

func main() {
	// start with an existing slice of ints
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// xi... UNFURLS the slice — it's equivalent to writing: sum(1, 2, 3, 4, 5, 6, 7, 8, 9)
	// without the `...` this would NOT compile: you can't pass a []int where ...int is expected
	x := sum(xi...)

	fmt.Println("The sum is", x)
}

// sum takes a VARIADIC parameter — same function from Lesson 135
// inside the function, ii is a []int regardless of whether the caller
// passed individual ints or unfurled a slice
func sum(ii ...int) int {
	fmt.Println(ii)        // prints the slice, e.g. [1 2 3 4 5 6 7 8 9]
	fmt.Printf("%T\n", ii) // prints []int — the variadic param is always a slice

	total := 0
	// range over the slice — see Lesson 074 for for-range
	for _, v := range ii {
		total += v
	}
	return total
}

/*
REMARKS:
- UNFURLING (also called "unpacking" or "spreading") lets you pass a SLICE to a VARIADIC function
- The syntax is: myFunc(mySlice...) — the `...` goes AFTER the argument, not before
- Without `...` on the call side, Go sees a type mismatch: []int vs ...int — it will not compile
- The unfurled slice and individual arguments produce IDENTICAL behavior inside the function;
  the variadic param is a []int either way
- KEY DISTINCTION — the `...` operator appears in two places:
    ■ DEFINITION side: `func sum(ii ...int)` — "accept zero or more ints, pack them into a slice"
    ■ CALL side: `sum(xi...)` — "take this slice and spread it into individual arguments"
- You can only unfurl a slice whose element type matches the variadic parameter type;
  e.g., []int can unfurl into ...int, but []string cannot
- BUILDING ON LESSON 135: Lesson 135 showed collecting individual args into a slice;
  this lesson shows the reverse — feeding a slice back into that same mechanism
- BUILDING ON LESSON 098: the underlying data is still a slice, so all slice concepts apply
- Real-world use: unfurling is common when building up arguments programmatically,
  e.g., appending to slices: `result = append(result, moreItems...)`
*/
