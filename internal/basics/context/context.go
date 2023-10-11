package context_examples

import (
	"context"
	"fmt"
	"log"
	"time"
)

func RunExamples() {
	start := time.Now()

	ctx := context.WithValue(context.Background(), "testKey", "test value")

	result, err := RunContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("result %s, %v elapsed\n", result, time.Since(start))
}

func RunContext(ctx context.Context) (string, error) {
	timeoutCtx, cancel := context.WithTimeout(ctx, time.Millisecond * 200)
	defer cancel()

	ctxValue := ctx.Value("testKey")
	fmt.Printf("ctx value: %s\n", ctxValue)

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

		resultch <- result{user, nil}
	}()

	select {
	case <-timeoutCtx.Done():
		return "", timeoutCtx.Err()
	case result := <-resultch:
		return result.user, result.err
	}
}

func thridPartyAPI() (string, error) {
	time.Sleep(time.Millisecond * 200)
	return "user data", nil
}