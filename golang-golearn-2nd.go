package main

import (
	"fmt"
)

func main() {
	//No.4
	for {
	}

	//No.2 & No.3
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	//No.1
	/*
		sum := 0
		for i := 0; i < 10; i++ {
			sum += i
		}
		fmt.Println("sum : ", sum)
	*/
}
