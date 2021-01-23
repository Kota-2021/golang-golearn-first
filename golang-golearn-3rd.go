package main

import "fmt"

//import "golang.org/x/tour/pic"

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

//No.11
/* func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
} */

//No.13
/* func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
} */

//No.15
/* func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
} */

//No.16
/* var pow = []int{1, 2, 4, 8, 16, 32, 64, 128} */

//No.18
/* func Pic(dx, dy int) [][]uint8 {
	ss := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		s := make([]uint8, dx)
		for x := 0; x < dx; x++ {
			s[x] = uint8((x + y) / 2)
		}
		ss[y] = s
	}
	return ss
} */

//No.19
/* type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex */

//No.20
/* type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
} */

//No.21
/* type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
} */

func main() {
	//No.22
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

	//No.20 No.21
	/* 	fmt.Println(m) */

	//No.19
	/* m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"]) */

	//No.18
	/* pic.Show(Pic) */

	//No.17
	/* pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	} */

	//No.16
	/* for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	} */

	//No.15
	/* 	var s []int
	   	printSlice(s)

	   	// append works on nil slices.
	   	s = append(s, 0)
	   	printSlice(s)

	   	// The slice grows as needed.
	   	s = append(s, 1)
	   	printSlice(s)

	   	// We can add more than one element at a time.
	   	s = append(s, 2, 3, 4)
	   	printSlice(s) */

	//No.14
	// Create a tic-tac-toe board.
	/* board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	} */

	//No.13
	/* a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d) */

	//No.12
	/* 	var s []int
	   	fmt.Println(s, len(s), cap(s))
	   	if s == nil {
	   		fmt.Println("nil!")
	   	} */

	//No.11
	/* 	s := []int{2, 3, 5, 7, 11, 13}
	   	printSlice(s)

	   	// Slice the slice to give it zero length.
	   	s = s[:0]
	   	printSlice(s)

	   	// Extend its length.
	   	s = s[:4]
	   	printSlice(s)

	   	// Drop its first two values.
	   	s = s[2:]
	   	printSlice(s) */

	//No.10
	/* 	s := []int{2, 3, 5, 7, 11, 13}

	   	s = s[1:4]
	   	fmt.Println(s)

	   	s = s[:2]
	   	fmt.Println(s)

	   	s = s[1:]
	   	fmt.Println(s) */

	//No.9
	/* q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s) */

	//No.8
	/* names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names) */

	//No.7
	/* primes := []int{2, 3, 5, 7, 11, 13}

	var s []int = primes[0:]
	fmt.Println(s) */

	//No.6
	/* var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes) */

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
