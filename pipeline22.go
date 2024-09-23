package piter

import (
	"context"
	"iter"
	"runtime"
)

func Iter22[I, O any](ctx context.Context, src iter.Seq2[I, error], fn func(I) (O, error)) iter.Seq2[O, error] {
	return func(yield func(O, error) bool) {
		for s, err := range src {
			var ok bool
			if err != nil {
				ok = yield(*new(O), err)
			} else {
				ok = yield(fn(s))
			}
			if !ok {
				return
			}
		}
	}
}

func Pipeline22[I, O any](ctx context.Context, src iter.Seq2[I, error], fn func(I) (O, error)) iter.Seq2[O, error] {
	inCh := newChan2(ctx, src)

	return func(yield func(O, error) bool) {
		for in := range inCh {
			var ok bool
			if in.e != nil {
				ok = yield(*new(O), in.e)
			} else {
				ok = yield(fn(in.d))
			}
			if !ok {
				return
			}
		}
	}
}

func FunOut22[I, O any](ctx context.Context, src iter.Seq2[I, error], fn func(I) (O, error)) iter.Seq2[O, error] {
	inCh := newChan2(ctx, src)

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
					if in.e != nil {
						outCh <- &chan2[O]{
							e: in.e,
						}
						continue
					}
					out, err := fn(in.d)
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
