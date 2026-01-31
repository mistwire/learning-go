/*
Take the code from the previous exercise, then store the VALUES of type person in a map with
the KEY of last name. Access each value in the map. Print out the values, ranging over the
slice.
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

	// fmt.Println(p1.first)
	// fmt.Println(p2.favIceCream[1])
	
	// for _, i := range p1.favIceCream{
	// 	fmt.Println(i)
	// }

	// for _, x := range p2.favIceCream{
	// 	fmt.Println(p2.first, "favorite ice cream is", x)
	// }

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}
	
	for _, v := range m{
		fmt.Println(v)
		for _, v2 := range v.favIceCream{
			fmt.Println(v.first, v2)
		}
	}
}