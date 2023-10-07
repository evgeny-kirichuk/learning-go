package main

import (
	"fmt"
	"strings"
)

func arrays () {
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