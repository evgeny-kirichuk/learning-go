package channels

import (
	"fmt"
	"time"
)

// TODO: learn more about channels and buffered channels

func RunAsyncChannels() {
	// channels
	println("Channels")

	resultch := make(chan string)
	go func() {
		resultch <- fetchData(1)
	}()

	result := <-resultch
	fmt.Println(result)
}

func RunChannels() {
	msgch := make(chan string, 128)
	msgch <- "Hello"
	msgch <- "World"
	msgch <- "!"
	close(msgch)

	// for with brake way
	for {
		msg, ok := <-msgch
		if !ok {
			break
		}
		fmt.Println(msg)
	}

	// range over channel way
	for msg := range msgch {
		fmt.Println("inside range")
		fmt.Println(msg)
	}

	fmt.Println("Done")
}

func fetchData(n int) string {
	time.Sleep(time.Second * 2)
	return fmt.Sprintf("Data %v", n)
}
