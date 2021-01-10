package main

import (
	"fmt"
	"math"
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

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	//No.6
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

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
