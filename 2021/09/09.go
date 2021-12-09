package main

import (
	"strings"

	"github.com/liennie/AdventOfCode/common/load"
	"github.com/liennie/AdventOfCode/common/log"
	"github.com/liennie/AdventOfCode/common/util"
)

func parse(filename string) [][]int {
	res := [][]int{}

	for line := range load.File(filename) {
		res = append(res, util.SliceAtoi(strings.Split(line, "")))
	}

	return res
}

func main() {
	defer util.Recover(log.Err)

	const filename = "input.txt"

	heightmap := parse(filename)

	// Part 1
	low := []int{}
	for i := range heightmap {
		for j := range heightmap[i] {
			c := heightmap[i][j]

			if (i == 0 || heightmap[i-1][j] > c) &&
				(j == 0 || heightmap[i][j-1] > c) &&
				(i == len(heightmap)-1 || heightmap[i+1][j] > c) &&
				(j == len(heightmap[i])-1 || heightmap[i][j+1] > c) {
				low = append(low, c)
			}
		}
	}
	log.Part1(util.Sum(low) + len(low))
}
