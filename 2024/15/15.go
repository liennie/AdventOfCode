package main

import (
	"github.com/liennie/AdventOfCode/pkg/channel"
	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
	"github.com/liennie/AdventOfCode/pkg/space"
)

func parse(filename string) (map[space.Point]rune, space.Point, []space.Point) {
	ch := load.Blocks(filename)
	defer channel.Drain(ch)

	var start space.Point
	grid := map[space.Point]rune{}
	y := 0
	for line := range <-ch {
		for x, r := range line {
			p := space.Point{X: x, Y: y}
			if r == '@' {
				start = p
				r = '.'
			}

			grid[p] = r
		}
		y++
	}

	var movement []space.Point
	for line := range <-ch {
		for _, r := range line {
			switch r {
			case '>':
				movement = append(movement, space.Point{X: 1})
			case '^':
				movement = append(movement, space.Point{Y: -1})
			case '<':
				movement = append(movement, space.Point{X: -1})
			case 'v':
				movement = append(movement, space.Point{Y: 1})
			default:
				evil.Panic("invalid movement %q", r)
			}
		}
	}

	return grid, start, movement
}

func move(grid map[space.Point]rune, pos, dir space.Point) bool {
	switch grid[pos] {
	case '.':
		return true

	case '#':
		return false

	case 'O':
		if move(grid, pos.Add(dir), dir) {
			grid[pos], grid[pos.Add(dir)] = grid[pos.Add(dir)], grid[pos]
			return true
		}
		return false

	default:
		evil.Panic("trying to move invalid spot %v %q", pos, grid[pos])
	}
	return false
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	grid, robot, movement := parse(filename)

	// Part 1
	for _, m := range movement {
		if move(grid, robot.Add(m), m) {
			robot = robot.Add(m)
		}
	}
	sum := 0
	for p, r := range grid {
		if r != 'O' {
			continue
		}

		sum += p.X + 100*p.Y
	}
	log.Part1(sum)
}
