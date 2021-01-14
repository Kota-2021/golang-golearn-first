package main

import "fmt"

//No.2
/* type Vertex struct {
	X int
	Y int
} */

//No.3
type Vertex struct {
	X int
	Y int
}

func main() {
	//No.3
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

	//No.2
	/* fmt.Println(Vertex{1, 2})
	m := Vertex{5, 10}
	fmt.Println(m.X)
	fmt.Println(m.Y) */

	//No.1
	/* 	i, j := 42, 2701

	   	p := &i         // point to i
	   	fmt.Println(*p) // read i through the pointer
	   	*p = 21         // set i through the pointer
	   	fmt.Println(i)  // see the new value of i

	   	p = &j         // point to j
	   	*p = *p / 37   // divide j through the pointer
	   	fmt.Println(j) // see the new value of j */
}
