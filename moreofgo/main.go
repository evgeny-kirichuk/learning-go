package main

import (
	"fmt"
	"strings"
)

type Vertex struct {
	X int
	Y int
}

func main() {
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
	fmt.Println("vertex",v1, v2, q, b)


	// arrays and slices
	var a [2]string // array of 2 strings
	a[0] = "Hello"
	a[1] = "World"
	a[0] = "Bye"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	nums := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = nums[1:4]
	halfS := s[0:]
	s = append(s, 3000)
	fmt.Println(s, halfS, cap(s))

	// make slice
	madenSlice := make([]int, 5)
	fmt.Println(madenSlice)

	// slice of slice
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}
	board[0][0] = "X"
	board[2][2] = "O"
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}


}