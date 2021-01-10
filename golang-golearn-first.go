package main

import (
	"fmt"
)

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// var c, python, java bool

//var i, j int = 1, 2

/*
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)
*/

//const Pi = 3.14

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	//17th
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))

	//16th
	/*
		const World = "世界"
		fmt.Println("Hello", World)
		fmt.Println("Happy", Pi, "Day")

		const Truth = true
		fmt.Println("Go rules?", Truth)
	*/

	//15th
	/*
		v := 5 + 7i // change me!
		fmt.Printf("v is of type %T\n", v)
	*/

	//14th
	/*
		var x, y int = 3, 4
		var f float64 = math.Sqrt(float64(x*x + y*y))
		var z uint = uint(f)
		fmt.Println(x, y, z)
	*/

	//13th
	/*
		var i int
		var f float64
		var b bool
		var s string
		fmt.Printf("%v %v %v %q\n", i, f, b, s)
	*/

	//12th
	/*
		fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
		fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
		fmt.Printf("Type: %T Value: %v\n", z, z)
	*/

	//11th
	/*
		var i, j int = 1, 2
		k := 3
		c, python, java := true, false, "no!"

		fmt.Println(i, j, k, c, python, java)
	*/

	//10th
	/*
		var c, python, java = true, false, "no!"
		fmt.Println(i, j, c, python, java)
	*/

	//9th
	/*
		var i int
		fmt.Println(i, c, python, java)
	*/

	//8th
	/*
		fmt.Println(split(17))
	*/

	//7th
	/*
		a, b := swap("hello", "world")
		fmt.Println(a, b)
	*/

	//5th 6th
	/*
		fmt.Println(add(42, 13))
	*/

	//4th
	//fmt.Println(math.pi) ← but code
	/*
		fmt.Println(math.Pi)
	*/

	// 3rd
	//Sqrt returns the square root of x.
	/*
		fmt.Printf("Now you have %e problems.\n", math.Sqrt(7))
		fmt.Printf("Now you have %E problems.\n", math.Sqrt(7))
		fmt.Printf("Now you have %f problems.\n", math.Sqrt(7))
		fmt.Printf("Now you have %F problems.\n", math.Sqrt(7))
		fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
		fmt.Printf("Now you have %G problems.\n", math.Sqrt(7))
	*/

	// 2nd
	/*
		rand.Seed(time.Now().UnixNano())
		fmt.Println("My favorite number is", rand.Intn(100))
	*/

	// 1st
	/*	fmt.Println("Hello, World!")
	 */
}
