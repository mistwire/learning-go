package main

import (
	"fmt"
)

func init() {
	fmt.Println("init func runs before main")
}

func main() {
	//SEQUENCE
	fmt.Println("this is the first func main statement to run")
	fmt.Println("this is the second statement to run")
	x := 40 // this is the third statement to run
	y := 5  // this is the fourth statement to run
	fmt.Printf(" x=%v \n y=%v\n", x, y)
}
