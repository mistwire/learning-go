package main

import "fmt"

func main() {
	an := make(map[string]int)
	an["Lucas"] = 28
	an["Steph"] = 37
	an["George"] = 78
	fmt.Println(an)

	// delete from a map
	delete(an, "George")

	//comma ok idiom

	v, ok := an["Georgey"] // Georgey doesn't exist in the map
	if ok {
		fmt.Println("The value prints: ", v)
	} else {
		fmt.Println("Key doesn't exit")
	}

	// how you normally see it - comma ok is combined with statement statement idiom:
	// 
	if x, ok := an["Lucas"]; !ok {
		fmt.Println("Key doesn't exit")
	} else {
		fmt.Println("The value prints: ", x)
	}

}
