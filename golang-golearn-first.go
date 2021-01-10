package main

import (
	"fmt"
)

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func swap(x, y string) (string, string) {
	return y, x
}

func add(x, y int) int {
	return x + y
}

func main() {
	//8th
	fmt.Println(split(17))

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
