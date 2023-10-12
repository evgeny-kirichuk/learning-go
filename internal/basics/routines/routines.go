package routines

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type UserProfile struct {
	ID       int
	Comments []string
	Likes    int
	Friends  []int
}

type Response struct {
	data any
	err  error
}

func handleGetUserProfile(id int) (*UserProfile, error) {
	responsech := make(chan Response, 3)
	wg := &sync.WaitGroup{}
	go getComments(id, responsech, wg)
	go getLikes(id, responsech, wg)
	go getFriends(id, responsech, wg)
	wg.Add(3)
	wg.Wait()
	close(responsech)

	userProfile := &UserProfile{
		ID: id,
	}

	for res := range responsech {
		if res.err != nil {
			return nil, res.err
		}

		switch msg := res.data.(type) {
		case []string:
			userProfile.Comments = msg
		case int:
			userProfile.Likes = msg
		case []int:
			userProfile.Friends = msg
		}
	}

	return userProfile, nil
}

func getComments(id int, respch chan Response, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 200)
	comments := []string{
		"comment 1",
		"comment 2",
	}
	respch <- Response{comments, nil}
	wg.Done()

}

func getLikes(id int, respch chan Response, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 200)
	respch <- Response{23, nil}
	wg.Done()
}

func getFriends(id int, respch chan Response, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 100)
	respch <- Response{[]int{12, 34}, nil}
	wg.Done()
}

func RunExamples() {
	start := time.Now()
	userProfile, err := handleGetUserProfile(1)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(userProfile)
	fmt.Printf("%v elapsed\n", time.Since(start))
}
