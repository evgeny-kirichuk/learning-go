package mutexes

import (
	"sync"
)

type State struct {
	mu sync.Mutex
	count int32
}

func (s *State) SetState(i int) {
	// lock the mutex to prevent concurrent access to the state.count so there is no race condition
	s.mu.Lock()
	defer s.mu.Unlock()
	s.count = int32(i)

	// atomic operation to prevent race condition
	// atomic.AddInt32(&s.count, int32(i))
}

func RunExamples() {
	state := &State{}
	wg := sync.WaitGroup{}

	// race condition because of the concurrent access to the state.count
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			state.SetState(i)
			wg.Done()
		}(i)
	}

	wg.Wait()
	println(state.count)
}