package main

import (
	"strconv"
	"strings"

	"github.com/liennie/AdventOfCode/common/load"
	"github.com/liennie/AdventOfCode/common/log"
	"github.com/liennie/AdventOfCode/common/util"
)

func parse(filename string) []map[rune]int {
	counts := []map[rune]int{}

	for line := range load.File(filename) {
		for i, r := range line {
			if i >= len(counts) {
				counts = append(counts, make([]map[rune]int, i-len(counts)+1)...)
			}
			if counts[i] == nil {
				counts[i] = map[rune]int{}
			}

			counts[i][r]++
		}
	}

	return counts
}

func parseBin(s string) int {
	i, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		panic(err)
	}
	return int(i)
}

func rates(counts []map[rune]int) (int, int) {
	gamma := &strings.Builder{}
	epsilon := &strings.Builder{}

	for _, m := range counts {
		max := 0
		min := int(^uint(0) >> 1)
		var maxR rune
		var minR rune
		for r, c := range m {
			if c > max {
				max = c
				maxR = r
			}
			if c < min {
				min = c
				minR = r
			}
		}
		gamma.WriteRune(maxR)
		epsilon.WriteRune(minR)
	}

	return parseBin(gamma.String()), parseBin(epsilon.String())
}

func main() {
	defer util.Recover(log.Err)

	const filename = "input.txt"

	counts := parse(filename)
	gamma, epsilon := rates(counts)

	// Part 1
	log.Part1(gamma * epsilon)
}
