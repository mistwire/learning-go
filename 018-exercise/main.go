// the Go tour - https://go.dev/tour

package main

import (
	"fmt"
	"math/cmplx"
	"math/rand"
)

// The var statement declares a list of variables; as in function argument lists, the type is last.
var c, python, java bool

const d int = 42

// A const statement can appear anywhere a var statement can.
// They cannot be declared using the := syntax.
const Pi = 3.14

// A var declaration can include initializers, one per variable.
var i, j int = 1, 2

func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

// return values may be named
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // A return statement without arguments returns the named return values
}

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	x := rand.Intn(1000)
	y := rand.Intn(1000)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(add(x, y))
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(split(100))
	// when you don't assign a value to a variable, it gets a zero/false value by default:
	var i int
	fmt.Println(i, c, python, java, d)
	var c, python, java = true, false, "no!"
	fmt.Printf("%T %v\t %T %v\t %T %v\t %T %v\t %T %v\t \n", i, i, j, j, c, c, python, python, java, java)
	// the := syntax is shorthand for declaring and initializing a variable
	k := 3
	fmt.Println(k)
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	// Variables declared without an explicit initial value are given their zero value.
	var m int
	var n float64
	var o bool
	var p string
	fmt.Printf("%v %v %v %q\n", m, n, o, p)
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
