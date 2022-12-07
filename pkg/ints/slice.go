package ints

import (
	"github.com/liennie/AdventOfCode/pkg/evil"
)

func Comb(n int) [][]int {
	if n < 0 {
		evil.Panic("Comb(%d)", n)
	}

	if n == 0 {
		return [][]int{{}}
	}

	res := Comb(n - 1)
	l := len(res)
	for i := 0; i < l; i++ {
		p := make([]int, len(res[i]))
		copy(p, res[i])
		p = append(p, n-1)
		res = append(res, p)
	}
	return res
}

func Uniq(ns []int) []int {
	res := []int{}
	s := map[int]bool{}

	for _, n := range ns {
		if !s[n] {
			s[n] = true
			res = append(res, n)
		}
	}

	return res
}
