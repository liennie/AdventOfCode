package main

import (
	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
	"github.com/liennie/AdventOfCode/pkg/set"
	"github.com/liennie/AdventOfCode/pkg/space"
)

func parse(filename string) (map[space.Point]rune, space.Point) {
	res := map[space.Point]rune{}
	start := space.Point{}
	load.Grid(filename, func(p space.Point, r rune) {
		if r == '^' {
			r = '.'
			start = p
		}

		res[p] = r
	})
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
	for grid[pos] != 0 {
		visited.Add(pos)

		switch grid[pos.Add(dir)] {
		case '#':
			dir = dir.Rot90(1)

		default:
			pos = pos.Add(dir)
		}
	}
	log.Part1(len(visited))

	// Part 2
	obstacles := set.New[space.Point]()
	for obsPos := range visited {
		if grid[obsPos] == '#' {
			continue
		}
		if obsPos == start {
			continue
		}

		grid[obsPos] = '#'

		pos = start
		dir = space.Point{Y: -1}
		visDir := map[space.Point]set.Set[space.Point]{}
		for grid[pos] != 0 && !visDir[pos].Contains(dir) {
			if visDir[pos] == nil {
				visDir[pos] = set.New[space.Point]()
			}
			visDir[pos].Add(dir)

			switch grid[pos.Add(dir)] {
			case '#':
				dir = dir.Rot90(1)

			default:
				pos = pos.Add(dir)
			}
		}
		if visDir[pos].Contains(dir) {
			obstacles.Add(obsPos)
		}

		grid[obsPos] = '.'
	}
	log.Part2(len(obstacles))
}
