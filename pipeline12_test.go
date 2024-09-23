package piter_test

import (
	"context"
	"fmt"
	"strconv"

	"github.com/yyyoichi/piter"
)

func ExampleIter12() {
	ctx := context.Background()
	src := toIter("10", "20kg", "30")
	dist := piter.Iter12(ctx, src, func(i string) (int, error) {
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

func ExamplePipeline12() {
	ctx := context.Background()
	src := toIter("10", "20kg", "30")
	dist := piter.Pipeline12(ctx, src, func(i string) (int, error) {
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

func ExampleFunOut12() {
	ctx := context.Background()
	src := toIter("10", "20kg", "30")
	dist := piter.Pipeline12(ctx, src, func(i string) (int, error) {
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
