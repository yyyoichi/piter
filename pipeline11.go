package piter

import (
	"context"
	"iter"
	"runtime"
)

func Iter11[I, O any](ctx context.Context, src iter.Seq[I], fn func(I) O) iter.Seq[O] {
	return func(yield func(O) bool) {
		for s := range src {
			if ok := yield(fn(s)); !ok {
				return
			}
		}
	}
}

func Pipeline11[I, O any](ctx context.Context, src iter.Seq[I], fn func(I) O) iter.Seq[O] {
	inCh := newChan(ctx, src)

	return func(yield func(O) bool) {
		for in := range inCh {
			if ok := yield(fn(in)); !ok {
				return
			}
		}
	}
}

func FunOut11[I, O any](ctx context.Context, src iter.Seq[I], fn func(I) O) iter.Seq[O] {
	inCh := newChan(ctx, src)

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
				case in, ok := <-inCh:
					if !ok {
						return
					}
					outCh <- fn(in)
				}
			}
		}()
		outChs[i] = outCh
	}

	multiplexedCh := FunIn(ctx, outChs...)
	return func(yield func(O) bool) {
		for out := range multiplexedCh {
			if ok := yield(out); !ok {
				return
			}
		}
	}
}
