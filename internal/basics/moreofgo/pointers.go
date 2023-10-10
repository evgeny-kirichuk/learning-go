package moreofgo

import "fmt"

type Vertex struct {
	X int
	Y int
}

func pointers() {
	// pointers and structs
	i, j := 42, 2701

	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)

	v1 := Vertex{1, 2}
	v2 := Vertex{X: 1}
	q := &v1
	q.X = 1e9
	b := &Vertex{1, 2}
	fmt.Println("vertex", v1, v2, q, b)
}
