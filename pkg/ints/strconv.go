package ints

import (
	"strconv"

	"github.com/liennie/AdventOfCode/pkg/evil"
)

func Atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		evil.Panic("Atoi(%s): %w", s, err)
	}
	return i
}

func SliceAtoi(s []string) []int {
	res := make([]int, len(s))
	for i, strNum := range s {
		res[i] = Atoi(strNum)
	}
	return res
}
