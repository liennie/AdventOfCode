package seq

import (
	"iter"
)

func Map[T1, T2 any](s iter.Seq[T1], f func(T1) T2) iter.Seq[T2] {
	return func(yield func(T2) bool) {
		for v := range s {
			if !yield(f(v)) {
				return
			}
		}
	}
}
