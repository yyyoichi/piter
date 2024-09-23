package piter

import (
	"context"
	"iter"
)

func newChan[T any](ctx context.Context, src iter.Seq[T]) <-chan T {
	inCh := make(chan T)
	go func() {
		defer close(inCh)
		for in := range src {
			select {
			case <-ctx.Done():
				return
			case inCh <- in:
			}
		}
	}()
	return inCh
}

type chan2[T any] struct {
	e error
	d T
}

func newChan2[T any](ctx context.Context, src iter.Seq2[T, error]) <-chan *chan2[T] {
	inCh := make(chan *chan2[T])
	go func() {
		defer close(inCh)
		for s, err := range src {
			select {
			case <-ctx.Done():
				return
			case inCh <- &chan2[T]{
				d: s,
				e: err,
			}:
			}
		}
	}()
	return inCh
}
