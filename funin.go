package piter

import (
	"context"
	"sync"
)

func FunIn[T any](cxt context.Context, channels ...chan T) <-chan T {
	var wg sync.WaitGroup
	multiplexedCh := make(chan T)
	multiplex := func(c <-chan T) {
		defer wg.Done()
		for i := range c {
			select {
			case <-cxt.Done():
				return
			case multiplexedCh <- i:
			}
		}
	}

	wg.Add(len(channels))
	for _, c := range channels {
		go multiplex(c)
	}

	go func() {
		wg.Wait()
		close(multiplexedCh)
	}()

	return multiplexedCh
}
