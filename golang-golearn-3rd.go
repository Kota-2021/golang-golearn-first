package main

import "fmt"

//No.2
/* type Vertex struct {
	X int
	Y int
} */

//No.3
/* type Vertex struct {
	X int
	Y int
} */

//No.4
/* type Vertex struct {
	X int
	Y int
} */

//No.5
/* type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
) */

func main() {
	//No.6
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	//No.5
	/* fmt.Println(v1, p, v2, v3) */

	//No.4
	/* v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v) */

	//No.3
	/* v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X) */

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
