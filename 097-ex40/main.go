// LESSON: Exercise - array literal with compiler-determined length
/*
● Use an array literal to store these elements in an array and let the compiler determine
the length of the array, then also print out the array and the length of the array
"Almond Biscotti Café", "Banana Pudding ", "Balsamic Strawberry (GF)", "Bittersweet
Chocolate (GF)", "Blueberry Pancake (GF)", "Bourbon Turtle (GF)", "Browned Butter
Cookie Dough", "Chocolate Covered Black Cherry (GF)", "Chocolate Fudge Brownie",
"Chocolate Peanut Butter (GF)", "Coffee (GF)", "Cookies & Cream", "Fresh Basil (GF)",
"Garden Mint Chip (GF)", "Lavender Lemon Honey (GF)", "Lemon Bar", "Madagascar
Vanilla (GF)", "Matcha (GF)", "Midnight Chocolate Sorbet (GF, V)", "Non-Dairy Chocolate
Peanut Butter (GF, V)", "Non-Dairy Coconut Matcha (GF, V)", "Non-Dairy Sunbutter
Cinnamon (GF, V)", "Orange Cream (GF) ", "Peanut Butter Cookie Dough", "Raspberry
Sorbet (GF, V)", "Salty Caramel (GF)", "Slate Slate Different", "Strawberry Lemonade
Sorbet (GF, V)", "Vanilla Caramel Blondie", "Vietnamese Cinnamon (GF)", "Wolverine
Tracks (GF)"
https://go.dev/play/p/OEHPUGDoWD_t
curriculum item # 094-hands-on-exercise-40
*/

package main

import "fmt"

func main() {
	ice_creams := [...]string{"Almond Biscotti Café", "Banana Pudding ", 
	"Balsamic Strawberry (GF)", "Bittersweet Chocolate (GF)", "Blueberry Pancake (GF)", 
	"Bourbon Turtle (GF)", "Browned Butter Cookie Dough", "Chocolate Covered Black Cherry (GF)", 
	"Chocolate Fudge Brownie", "Chocolate Peanut Butter (GF)", "Coffee (GF)", "Cookies & Cream", 
	"Fresh Basil (GF)", "Garden Mint Chip (GF)", "Lavender Lemon Honey (GF)", "Lemon Bar", 
	"Madagascar Vanilla (GF)", "Matcha (GF)", "Midnight Chocolate Sorbet (GF, V)", 
	"Non-Dairy Chocolate Peanut Butter (GF, V)", "Non-Dairy Coconut Matcha (GF, V)", 
	"Non-Dairy Sunbutter Cinnamon (GF, V)", "Orange Cream (GF) ", "Peanut Butter Cookie Dough", 
	"Raspberry Sorbet (GF, V)", "Salty Caramel (GF)", "Slate Slate Different", 
	"Strawberry Lemonade Sorbet (GF, V)", "Vanilla Caramel Blondie", "Vietnamese Cinnamon (GF)", 
	"Wolverine Tracks (GF)"}

	fmt.Println(ice_creams)
	fmt.Println(len(ice_creams))
	fmt.Printf("%T\n", ice_creams)

}
