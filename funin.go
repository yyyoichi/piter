package piter

import (
	"context"
	"runtime"
	"sync"
)

func FunOut[I, O any](ctx context.Context, src <-chan I, fn func(I) O) <-chan O {
	c := runtime.GOMAXPROCS(0)
	outChs := make([]chan O, c)
	for i := range c {
		outCh := make(chan O)
		go func() {
			defer close(outCh)
			for {
				select {
				case <-ctx.Done():
					return
				case in, ok := <-src:
					if !ok {
						return
					}
					outCh <- fn(in)
				}
			}
		}()
		outChs[i] = outCh
	}
	return FunIn(ctx, outChs...)
}

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
