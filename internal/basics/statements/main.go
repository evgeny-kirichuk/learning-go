package statements

import (
	"fmt"
	"runtime"
)

func main() {
	defer fmt.Println("Deferred statement")
	count := 0

	// while-like loop
	for count < 5 {
		fmt.Println(count)
		count = count + 1
	}

	// for loop
	for i := count; i <= 10; i++ {
		// defer statements are executed in LIFO order
		defer fmt.Println(i)
	}

	// infinite loop
	// for {
	// 	fmt.Println("Hello")
	// }

	// if statement
	if count == 5 {
		fmt.Println("Count is 5")
	}

	// if with a short statement
	if x := count; x == 5 {
		fmt.Println("x is 5")
	}

	// switch statement
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.", os)
	}

}