package moreofgo

import (
	"fmt"
	"sync"
	"time"
)

type User struct {
	firstName string
	lastName string
	age uint
}

var wg = sync.WaitGroup{}

func maps() {
	var username string

		fmt.Println("Enter your name: ")
		fmt.Scan(&username)
		wg.Add(1)
		go sendMessage("Hello, " + username);
		wg.Wait()
}

func sendMessage(text string) {
	defer wg.Done()
	time.Sleep(5 * time.Second)
	message := fmt.Sprintf("New message: %v \n", text)
	fmt.Println("###########")
	fmt.Printf("sending message: %v",message)
	fmt.Println("###########")

}