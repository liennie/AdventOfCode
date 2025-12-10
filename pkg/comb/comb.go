package comb

import (
	"iter"
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
