package main

import (
	"strings"

	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
	"github.com/liennie/AdventOfCode/pkg/space"
)

func parse(filename string) map[space.Point]string {
	res := map[space.Point]string{}
	y := 0
	for line := range load.File(filename) {
		for x, letter := range strings.Split(line, "") {
			res[space.Point{x, y}] = letter
		}
		y++
	}
	return res
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	grid := parse(filename)

	// Part 1
	dirs := []space.Point{
		{1, 0},
		{1, 1},
		{0, 1},
		{-1, 1},
		{-1, 0},
		{-1, -1},
		{0, -1},
		{1, -1},
	}
	word := strings.Split("XMAS", "")

	count := 0
	for start := range grid {
	dirs:
		for _, dir := range dirs {
			for n, letter := range word {
				if grid[start.Add(dir.Scale(n))] != letter {
					continue dirs
				}
			}

			count++
		}
	}
	log.Part1(count)

	// Part 2
	patterns := []map[space.Point]string{
		{{-1, 1}: "M", {-1, -1}: "M", {1, -1}: "S", {1, 1}: "S"},
		{{-1, 1}: "S", {-1, -1}: "M", {1, -1}: "M", {1, 1}: "S"},
		{{-1, 1}: "S", {-1, -1}: "S", {1, -1}: "M", {1, 1}: "M"},
		{{-1, 1}: "M", {-1, -1}: "S", {1, -1}: "S", {1, 1}: "M"},
	}

	count = 0
	for start := range grid {
		if grid[start] != "A" {
			continue
		}

	patterns:
		for _, pattern := range patterns {
			for offset, letter := range pattern {
				if grid[start.Add(offset)] != letter {
					continue patterns
				}
			}

			count++
		}
	}
	log.Part2(count)
}
