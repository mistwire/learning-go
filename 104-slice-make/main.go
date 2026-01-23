package main

import "fmt"

func main() {
	// si := []string{"a", "b", "c"}
	// fmt.Println(si)

	// make creates a slice, but doesn't put any values into it
	// 		   type, len, cap
	xi := make([]int, 0, 10) //0 length, 10 capacity
	fmt.Println(xi)
	fmt.Printf("length: %v\n",len(xi))
	fmt.Printf("capacity: %v\n",cap(xi))
	fmt.Println("------------")
	xi = append(xi, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(xi)
	fmt.Printf("length: %v\n",len(xi))
	fmt.Printf("capacity: %v\n",cap(xi))
	fmt.Println("------------")
	xi = append(xi, 10, 11, 12, 13)
	fmt.Println(xi)
	fmt.Printf("length: %v\n",len(xi))
	fmt.Printf("capacity: %v\n",cap(xi))
}

