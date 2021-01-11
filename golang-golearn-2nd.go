package main

import (
	"fmt"
)

//No.4
/*
func sqrt(x float64) string {
	if x < 0 {
		// return sqrt(x) <- but code
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}
*/

//No.6
/*
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}
*/

//No.7
/*
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim

}
*/

//No.8
/* func mySqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
} */

func main() {
	//No.13
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")

	//No.12
	/* hello := []string{"Hello", "(o^｡^o)", ".", "World", "this", "like", "I", ".", "world"}
	for i := 1; i < 9; i++ {
		defer fmt.Println(hello[i])
	}
	fmt.Println(hello[0]) */

	//defer fmt.Println("world")
	//fmt.Println("hello")

	//No.11
	/* t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	} */

	//No.10
	/* fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	} */

	//No.9
	/* fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	} */

	//No.8
	/* a := 2.0
	fmt.Println(mySqrt(a))
	fmt.Println(math.Sqrt(a)) */

	//No.7
	/* fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	) */

	//No.6
	/*
		fmt.Println(
			pow(3, 2, 10),
			pow(3, 3, 20),
		)
	*/

	//No.5
	/*
		fmt.Println(sqrt(2), sqrt(-2))
	*/

	//No.4
	/*
		for {
		}
	*/

	//No.2 & No.3
	/*
		sum := 1
		for sum < 1000 {
			sum += sum
		}
		fmt.Println(sum)
	*/

	//No.1
	/*
		sum := 0
		for i := 0; i < 10; i++ {
			sum += i
		}
		fmt.Println("sum : ", sum)
	*/
}
