package ints

import (
	"iter"
	"math"
	"slices"

	"github.com/liennie/AdventOfCode/pkg/evil"
)

func Mod(a, b int) int {
	m := a % b
	if m < 0 {
		return m + b
	}
	return m
}

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func Pow(n, exp int) int {
	evil.Assert(exp >= 0, "negative exponent %d", exp)

	res := 1
	for range exp {
		res *= n
	}
	return res
}

func Clamp(a, min, max int) int {
	if a < min {
		return min
	}
	if a > max {
		return max
	}
	return a
}

func Wrap(a, min, max int) int {
	return Mod((a-min), (max-min+1)) + min
}

func Min(ns ...int) int {
	return MinSeq(slices.Values(ns))
}

func MinSeq(ns iter.Seq[int]) int {
	min := math.MaxInt
	for n := range ns {
		if n < min {
			min = n
		}
	}

	return min

}

func MinFunc[T any](f func(T) int, ts ...T) int {
	return MinSeqFunc(f, slices.Values(ts))
}

func MinSeqFunc[T any](f func(T) int, ts iter.Seq[T]) int {
	min := math.MaxInt
	for t := range ts {
		if n := f(t); n < min {
			min = n
		}
	}

	return min

}

func Max(ns ...int) int {
	return MaxSeq(slices.Values(ns))
}

func MaxSeq(ns iter.Seq[int]) int {
	max := math.MinInt
	for n := range ns {
		if n > max {
			max = n
		}
	}

	return max

}

func MaxFunc[T any](f func(T) int, ts ...T) int {
	return MaxSeqFunc(f, slices.Values(ts))
}

func MaxSeqFunc[T any](f func(T) int, ts iter.Seq[T]) int {
	max := math.MinInt
	for t := range ts {
		if n := f(t); n > max {
			max = n
		}
	}

	return max

}

func Sum(s ...int) int {
	return SumSeq(slices.Values(s))
}

func SumSeq(s iter.Seq[int]) int {
	sum := 0
	for n := range s {
		sum += n
	}
	return sum
}

func SumFunc[T any](f func(T) int, ts ...T) int {
	return SumSeqFunc(f, slices.Values(ts))
}

func SumSeqFunc[T any](f func(T) int, ts iter.Seq[T]) int {
	sum := 0
	for t := range ts {
		sum += f(t)
	}
	return sum
}

func Product(s ...int) int {
	return ProductSeq(slices.Values(s))
}

func ProductSeq(s iter.Seq[int]) int {
	prod := 1
	for n := range s {
		prod *= n
	}
	return prod
}

func ProductFunc[T any](f func(T) int, ts ...T) int {
	return ProductSeqFunc(f, slices.Values(ts))
}

func ProductSeqFunc[T any](f func(T) int, ts iter.Seq[T]) int {
	prod := 1
	for t := range ts {
		prod *= f(t)
	}
	return prod
}

func GCD(a, b int) int {
	a = Abs(a)
	b = Abs(b)
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func LCM(a, b int) int {
	return Abs(a*b) / GCD(a, b)
}
