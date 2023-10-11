package context

import (
	"context"
	"time"
)

func RunExamples() {

}

func RunContext() {
	ctx, cancel := context.WithTimeout(time.Second * 2)
	defer cancel()

	type result struct	{
		user string
		err error
	}

	resultch := make(chan result, 1)

	go func() {
		user, err := thridPartyAPI()
		if err != nil {
			cancel()
			return
		}
		println(user)
	}()
}

func thridPartyAPI() (string, error) {
	time.Sleep(time.Second * 5)
	return "user data", nil
}