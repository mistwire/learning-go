/*
Create your own type “person” which will have an underlying type of “struct” so that it can store
the following data:
● first name
● last name
● favorite ice cream flavors
Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice
which stores the favorite flavors.
*/

package main

import "fmt"

type person struct {
	first string
	last string
	favIceCream []string
}

func main (){
	p1 := person{
		first: "Chris",
		last: "Williams",
		favIceCream: []string{"vanilla", "maple crunch", "walnut"},
	}

	p2 := person{
		first: "Kim",
		last: "Fabre",
		favIceCream: []string{"chocolate", "maple walnut", "almond joy"},
	}

	fmt.Println(p1.first)
	fmt.Println(p2.favIceCream[1])
	
	for _, i := range p1.favIceCream{
		fmt.Println(i)
	}

	for _, x := range p2.favIceCream{
		fmt.Println(p2.first, "favorite ice cream is", x)
	}

}