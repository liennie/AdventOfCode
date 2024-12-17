package evil

import (
	"strconv"
	"strings"
)

func Atoi(s string) int {
	i, err := strconv.Atoi(strings.TrimSpace(s))
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

func SliceItoa(s []int) []string {
	res := make([]string, len(s))
	for i, n := range s {
		res[i] = strconv.Itoa(n)
	}
	return res
}

func Split(s string, sep string) []int {
	return SliceAtoi(strings.Split(s, sep))
}

func SplitN(s string, sep string, n int) []int {
	return SliceAtoi(strings.SplitN(s, sep, n))
}

func Fields(s string) []int {
	return SliceAtoi(strings.Fields(s))
}

func Join(elems []int, sep string) string {
	return strings.Join(SliceItoa(elems), sep)
}
