package piter_test

import (
	"context"
	"fmt"
	"iter"
	"strconv"

	"github.com/yyyoichi/piter"
)

func ExampleIter22() {
	ctx := context.Background()
	src := toIter2("10", "20kg", "30")
	dist := piter.Iter22(ctx, src, func(i string) (int, error) {
		return strconv.Atoi(i)
	})
	for d, err := range dist {
		if err != nil {
			fmt.Println("err")
		} else {
			fmt.Println(d)
		}
	}
	// Output:
	// 10
	// err
	// 30
}

func ExamplePipeline22() {
	ctx := context.Background()
	src := toIter2("10", "20kg", "30")
	dist := piter.Iter22(ctx, src, func(i string) (int, error) {
		return strconv.Atoi(i)
	})
	for d, err := range dist {
		if err != nil {
			fmt.Println("err")
		} else {
			fmt.Println(d)
		}
	}
	// Output:
	// 10
	// err
	// 30
}

func ExampleFunOut22() {
	ctx := context.Background()
	src := toIter2("10", "20kg", "30")
	dist := piter.Iter22(ctx, src, func(i string) (int, error) {
		return strconv.Atoi(i)
	})
	for d, err := range dist {
		if err != nil {
			fmt.Println("err")
		} else {
			fmt.Println(d)
		}
	}
	// Unordered output:
	// 10
	// err
	// 30
}

func toIter2[T any](src ...T) iter.Seq2[T, error] {
	return func(yield func(T, error) bool) {
		for _, s := range src {
			_ = yield(s, nil)
		}
	}
}
