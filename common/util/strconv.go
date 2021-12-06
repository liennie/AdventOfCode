package util

import (
	"strconv"
)

func Atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		Panic("Atoi(%s): %w", s, err)
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
