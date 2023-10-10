package channels

import (
	"fmt"
	"time"
)

// TODO: learn more about channels and buffered channels

func RunChannels() {
	// channels
	println("Channels")

	resultch := make(chan string)
	go func() {
		resultch <- fhetchData(1)
	}()

	result := <-resultch
	fmt.Println(result)
}

func fhetchData(n int) string {
	time.Sleep(time.Second * 2)
	return fmt.Sprintf("Data %v", n)
}