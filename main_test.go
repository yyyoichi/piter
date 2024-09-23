package piter_test

import (
	"context"
	"fmt"
	"iter"
	"strconv"

	"github.com/yyyoichi/piter"
)

func Example() {
	ctx := context.Background()

	twice := func(i int) int {
		return i * 2
	}
	var src iter.Seq[int] = func(yield func(int) bool) {
		_ = yield(2)
		_ = yield(3)
		_ = yield(4)
	}
	d1 := piter.Iter11(ctx, src, twice)
	d2 := piter.Pipeline11(ctx, d1, twice)
	d3 := piter.FunOut11(ctx, d2, twice)

	for d := range d3 {
		fmt.Println(d)
	}

	var strSrc iter.Seq[string] = func(yield func(string) bool) {
		_ = yield("a")
		_ = yield("2")
		_ = yield("4")
	}
	var ns = make([]int, 0, 3)
	for n, err := range piter.Pipeline12(ctx, strSrc, strconv.Atoi) {
		if err != nil {
			break
		}
		ns = append(ns, n)
	}
	for _, n := range ns {
		// Expect no output
		fmt.Println(n)
	}

	// Unordered output:
	// 16
	// 24
	// 32
}
