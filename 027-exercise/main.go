package main

import (
	"fmt"
)

func main() {
	// ● var for zero value
	var x int
	fmt.Println(x)

	// ● short declaration operator
	y := 42
	fmt.Println(y)

	// ● multiple initializations
	a, b := 43, "Your Mom"
	fmt.Println(a, b)

	// ● var when specificity is required
	var c uint32 = 1231342343
	fmt.Printf("%v is of type %T\n", c, c)

	// ● blank identifier
	e, f, _ := 45, 46, 47
	fmt.Println(e, f)

}
