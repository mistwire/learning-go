/*
● VARIADIC PARAMETERS — functions that accept a variable number of arguments
● Use the `...` (ellipsis) operator before the type in the parameter list
● Inside the function, the variadic parameter becomes a SLICE of that type
● BUILDING ON LESSON 134: extends function syntax with variadic parameter form
● func (receiver) identifier(params, variadic ...Type) (returns) { code }
*/
package main

import "fmt"

func main() {
	// call sum with any number of int arguments — Go packs them into a []int
	x := sum(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("The sum is", x)
}

// func (r receiver) identifier(p parameters) (return(s)) {code}

// sum takes a VARIADIC parameter: `ii ...int` means "zero or more ints"
// inside the function, ii is a []int (a slice) — see Lesson 098 for slice basics
func sum(ii ...int) int {
	fmt.Println(ii)        // prints the slice, e.g. [1 2 3 4 5 6 7 8 9]
	fmt.Printf("%T\n", ii) // prints []int — confirms the variadic param is a slice

	total := 0
	// range over the slice just like any other slice (see Lesson 074 for for-range)
	for _, v := range ii {
		total += v
	}
	return total
}

/*
REMARKS:
- The `...` operator before the type makes a parameter VARIADIC — it accepts zero or more arguments
- Inside the function the variadic parameter is a regular SLICE ([]T), so all slice operations work on it
- The variadic parameter MUST be the LAST parameter in the function signature
- You can also UNFURL an existing slice into a variadic call using `...` after the argument:
    nums := []int{1, 2, 3}
    sum(nums...)   // passes slice elements as individual arguments
- fmt.Println is itself a variadic function: func Println(a ...any) — that's why it accepts any number of args
- BUILDING ON LESSON 134: this is the same func syntax, just with `...` added before the param type
- BUILDING ON LESSON 098: the variadic param becomes a slice, so len(), cap(), range all apply
*/
