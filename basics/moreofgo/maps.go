package moreofgo

import (
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName string
	age uint
}

func maps() {
	var username string
	for {
		fmt.Println("Enter your name: ")
		fmt.Scan(&username)
		go sendMessage("Hello, " + username);
	}
}

func sendMessage(text string) {
	time.Sleep(5 * time.Second)
	message := fmt.Sprintf("New message: %v \n", text)
	fmt.Println("###########")
	fmt.Printf("sending message: %v",message)
	fmt.Println("###########")

}