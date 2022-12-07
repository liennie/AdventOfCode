package evil

import (
	"strconv"
	"strings"
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

func Split(s string, sep string) []int {
	return SliceAtoi(strings.Split(s, sep))
}

func SplitN(s string, sep string, n int) []int {
	return SliceAtoi(strings.SplitN(s, sep, n))
}
