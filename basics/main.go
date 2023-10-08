package main

import (
	"fmt"
	"learning/basics/api"
	"learning/basics/fuzz"
	"learning/basics/generics"
)

func main() {
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}
	reversed, err := fuzz.Reverse("Hello, world")
	sum := generics.SumNumbers(ints)
	fmt.Print(reversed, sum, err)

	api.StartSerevr()
	// gowiki.StartSerevr()
}