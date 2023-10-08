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
	var count uint = 10
	userData := map[string]User{}
	userData["John"] = User{"John", "Doe", count}
	userData["Adam"] = User{"Adam", "Doe", count}
	fmt.Println(userData, &count)

	for key, user := range userData {
		sendMessage("Hello, " + user.firstName + " " + user.lastName + " (" + key + ")");
	}
}

func sendMessage(text string) {
	message := fmt.Sprintf("New message: %v \n", text)
	fmt.Println("###########")
	fmt.Printf("sending message: %v",message)
	fmt.Println("###########")

	time.Sleep(5 * time.Second)
}