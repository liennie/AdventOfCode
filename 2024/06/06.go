package main

import (
	"strings"

	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
	"github.com/liennie/AdventOfCode/pkg/set"
	"github.com/liennie/AdventOfCode/pkg/space"
)

func parse(filename string) (map[space.Point]string, space.Point) {
	res := map[space.Point]string{}
	start := space.Point{}
	y := 0
	for line := range load.File(filename) {
		for x, cell := range strings.Split(line, "") {
			if cell == "^" {
				cell = "."
				start = space.Point{x, y}
			}

			res[space.Point{x, y}] = cell
		}
		y++
	}
	return res, start
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	grid, start := parse(filename)

	// Part 1
	pos := start
	dir := space.Point{Y: -1}
	visited := set.New[space.Point]()
	for grid[pos] != "" {
		visited.Add(pos)

		switch grid[pos.Add(dir)] {
		case "#":
			dir = dir.Rot90(1)

		default:
			pos = pos.Add(dir)
		}
	}
	log.Part1(len(visited))
}
