// LESSON: Multidimensional Slices - slice of slices (2D data structures)
package main

import "fmt"

func main() {
	// strings with independent info:
	jb := []string{"James", "Bond", "Martini", "Chocolate"}
	jm := []string{"Jenny", "Moneypenny", "Guiness", "Strawberry"}
	fmt.Println(jb)
	fmt.Println(jm)

	//now we want to store them together in a slice of a slice:
	xxs := [][]string{jb, jm}
	fmt.Println(xxs)
}
