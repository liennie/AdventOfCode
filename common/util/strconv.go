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
