package main

import "fmt"

func main() {
	xs := []string{"Almond Biscotti Café", "Banana Pudding ", "Balsamic Strawberry (GF)"}

	fmt.Println(xs)
	fmt.Printf("%T\n", xs)

	// blank identifier _ to not return a value
	for _, v := range xs {
		fmt.Printf("%v\n", v)
	}

	fmt.Println("---------------------")
	fmt.Println(xs[0])
	fmt.Println(xs[1])
	fmt.Println(xs[2])
	
	// doesn't work:
	//fmt.Println(xs[3])
	fmt.Println("---get len of a slice------")
	fmt.Println(len(xs))
	
	//the old way:
	fmt.Println("---the old way------")
	for i := 0; i < len(xs); i++ {
		fmt.Println(i)
	}
}
