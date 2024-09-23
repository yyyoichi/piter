package piter_test

import (
	"context"
	"fmt"
	"iter"
	"strconv"

	"github.com/yyyoichi/piter"
)

func ExampleIter11() {
	ctx := context.Background()
	src := toIter(1, 2, 3)
	dist := piter.Iter11(ctx, src, func(i int) string {
		return strconv.Itoa(i*10) + "kg"
	})
	for d := range dist {
		fmt.Println(d)
	}
	// Output:
	// 10kg
	// 20kg
	// 30kg
}

func ExamplePipeline11() {
	ctx := context.Background()
	src := toIter(1, 2, 3)
	dist := piter.Pipeline11(ctx, src, func(i int) string {
		return strconv.Itoa(i*10) + "kg"
	})
	for d := range dist {
		fmt.Println(d)
	}
	// Output:
	// 10kg
	// 20kg
	// 30kg
}

func ExampleFunOut11() {
	ctx := context.Background()
	src := toIter(1, 2, 3)
	dist := piter.Iter11(ctx, src, func(i int) string {
		return strconv.Itoa(i*10) + "kg"
	})
	for d := range dist {
		fmt.Println(d)
	}
	// Unordered output:
	// 10kg
	// 20kg
	// 30kg
}

func toIter[T any](src ...T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, s := range src {
			_ = yield(s)
		}
	}
}
