package main

import (
	"fmt"
	"math/cmplx"
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

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	//12th
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

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
