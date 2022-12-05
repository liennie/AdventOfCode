package util

import (
	"math"
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

func Clamp(a, min, max int) int {
	if a < min {
		return min
	}
	if a > max {
		return max
	}
	return a
}

func Min(ns ...int) int {
	min := math.MaxInt
	for _, n := range ns {
		if n < min {
			min = n
		}
	}

	return min
}

func MinFunc[T any](f func(T) int, ts ...T) int {
	min := math.MaxInt
	for _, t := range ts {
		if n := f(t); n < min {
			min = n
		}
	}

	return min
}

func Max(ns ...int) int {
	max := math.MinInt
	for _, n := range ns {
		if n > max {
			max = n
		}
	}

	return max
}

func MaxFunc[T any](f func(T) int, ts ...T) int {
	max := math.MinInt
	for _, t := range ts {
		if n := f(t); n > max {
			max = n
		}
	}

	return max
}

func Sum(s ...int) int {
	sum := 0
	for _, n := range s {
		sum += n
	}
	return sum
}

func SumFunc[T any](f func(T) int, ts ...T) int {
	sum := 0
	for _, t := range ts {
		sum += f(t)
	}
	return sum
}

func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func LCM(a, b int) int {
	return Abs(a*b) / GCD(a, b)
}
