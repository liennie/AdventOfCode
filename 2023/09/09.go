package main

import (
	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
)

func parse(filename string) [][]int {
	res := [][]int{}
	for line := range load.File(filename) {
		res = append(res, evil.Fields(line))
	}
	return res
}

func differences(values []int) []int {
	res := make([]int, len(values)-1)
	for i := 0; i < len(values)-1; i++ {
		res[i] = values[i+1] - values[i]
	}
	return res
}

func allZeros(values []int) bool {
	for _, v := range values {
		if v != 0 {
			return false
		}
	}
	return true
}

func last[T any, S []T](s S) T {
	return s[len(s)-1]
}

func next(values []int) int {
	vv := [][]int{values}
	for !allZeros(last(vv)) && len(last(vv)) > 1 {
		vv = append(vv, differences(last(vv)))
	}

	for i := len(vv) - 1; i > 0; i-- {
		vv[i-1] = append(vv[i-1], last(vv[i-1])+last(vv[i]))
	}

	return last(vv[0])
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	values := parse(filename)

	// Part 1
	sum := 0
	for _, v := range values {
		sum += next(v)
	}
	log.Part1(sum)
}
