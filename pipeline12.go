package piter

import (
	"context"
	"iter"
	"runtime"
)

func Iter12[I, O any](ctx context.Context, src iter.Seq[I], fn func(I) (O, error)) iter.Seq2[O, error] {
	return func(yield func(O, error) bool) {
		for s := range src {
			if ok := yield(fn(s)); !ok {
				return
			}
		}
	}
}

func Pipeline12[I, O any](ctx context.Context, src iter.Seq[I], fn func(I) (O, error)) iter.Seq2[O, error] {
	inCh := newChan(ctx, src)

	return func(yield func(O, error) bool) {
		for in := range inCh {
			if ok := yield(fn(in)); !ok {
				return
			}
		}
	}
}

func FunOut12[I, O any](ctx context.Context, src iter.Seq[I], fn func(I) (O, error)) iter.Seq2[O, error] {
	inCh := newChan(ctx, src)

	c := runtime.GOMAXPROCS(0)
	outChs := make([]chan *chan2[O], 0, c)
	for i := range c {
		outCh := make(chan *chan2[O])
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
					out, err := fn(in)
					outCh <- &chan2[O]{
						d: out,
						e: err,
					}
				}
			}
		}()
		outChs[i] = outCh
	}

	multiplexedCh := FunIn(ctx, outChs...)
	return func(yield func(O, error) bool) {
		for out := range multiplexedCh {
			if ok := yield(out.d, out.e); !ok {
				return
			}
		}
	}
}
