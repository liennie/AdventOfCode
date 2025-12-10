package comb

import (
	"iter"

	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/ints"
)

func Comb[T any](s []T) iter.Seq[[]T] {
	return func(yield func([]T) bool) {
		if len(s) == 0 {
			yield([]T{})
			return
		}

		Comb(s[1:])(func(c []T) bool {
			if !yield(c) {
				return false
			}
			return yield(append([]T{s[0]}, c...))
		})
	}
}

func Uniq[T comparable](s iter.Seq[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		d := map[T]struct{}{}
		for v := range s {
			if _, ok := d[v]; !ok {
				d[v] = struct{}{}
				if !yield(v) {
					return
				}
			}
		}
	}
}

func Choose(n, k int) int {
	evil.Assert(0 <= k && k <= n && n >= 0, "choose: invalid args n=%d k=%d", n, k)

	d := 1
	for m := 1; m <= k; m++ {
		d *= m
	}

	ch := 1
	for m := n; m > n-k; m-- {
		gcd := ints.GCD(m, d)
		d /= gcd

		pch := ch
		ch *= (m / gcd)

		evil.Assert(ch >= pch, "choose: overflow detected v %d m %d", pch, m)
	}
	evil.Assert(d == 1, "choose: remaining division %d", d)
	return ch
}
