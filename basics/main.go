package main

import (
	"fmt"
	"learning/basics/fuzz"
)

func main() {
	println("Hello, world")

	reversed, err := fuzz.Reverse("Hello, world")
	fmt.Print(reversed, err)
}