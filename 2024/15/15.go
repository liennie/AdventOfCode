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

func canMove(grid map[space.Point]rune, pos, dir space.Point) bool {
	switch grid[pos] {
	case '.':
		return true

	case '#':
		return false

	case 'O':
		return canMove(grid, pos.Add(dir), dir)

	case '[':
		if dir.X == 0 {
			return canMove(grid, pos.Add(dir), dir) && canMove(grid, pos.Add(space.Point{X: 1}).Add(dir), dir)
		}
		return canMove(grid, pos.Add(dir), dir)

	case ']':
		if dir.X == 0 {
			return canMove(grid, pos.Add(dir), dir) && canMove(grid, pos.Add(space.Point{X: -1}).Add(dir), dir)
		}
		return canMove(grid, pos.Add(dir), dir)

	default:
		evil.Panic("trying to move invalid spot %v %q", pos, grid[pos])
	}

	return false
}

func swap(grid map[space.Point]rune, pos, dir space.Point) {
	grid[pos], grid[pos.Add(dir)] = grid[pos.Add(dir)], grid[pos]
}

func move(grid map[space.Point]rune, pos, dir space.Point) bool {
	if !canMove(grid, pos, dir) {
		return false
	}

	switch grid[pos] {
	case '.':
		return true

	case '#':
		return false

	case 'O':
		move(grid, pos.Add(dir), dir)
		swap(grid, pos, dir)

	case '[':
		sec := pos.Add(space.Point{X: 1})

		switch {
		case dir.X == 0:
			move(grid, pos.Add(dir), dir)
			move(grid, sec.Add(dir), dir)
			swap(grid, pos, dir)
			swap(grid, sec, dir)

		case dir.X > 0:
			move(grid, sec.Add(dir), dir)
			swap(grid, sec, dir)
			swap(grid, pos, dir)

		case dir.X < 0:
			move(grid, pos.Add(dir), dir)
			swap(grid, pos, dir)
			swap(grid, sec, dir)
		}

	case ']':
		sec := pos.Add(space.Point{X: -1})

		switch {
		case dir.X == 0:
			move(grid, pos.Add(dir), dir)
			move(grid, sec.Add(dir), dir)
			swap(grid, pos, dir)
			swap(grid, sec, dir)

		case dir.X > 0:
			move(grid, pos.Add(dir), dir)
			swap(grid, pos, dir)
			swap(grid, sec, dir)

		case dir.X < 0:
			move(grid, sec.Add(dir), dir)
			swap(grid, sec, dir)
			swap(grid, pos, dir)
		}

	default:
		evil.Panic("trying to move invalid spot %v %q", pos, grid[pos])
	}

	return true
}

func enlarge(grid map[space.Point]rune) map[space.Point]rune {
	res := map[space.Point]rune{}
	for p, r := range grid {
		next := space.Point{X: p.X * 2, Y: p.Y}
		switch r {
		case '.', '#':
			res[next] = r
			res[next.Add(space.Point{X: 1})] = r

		case 'O':
			res[next] = '['
			res[next.Add(space.Point{X: 1})] = ']'

		default:
			evil.Panic("invalid char at pos %v %q", p, r)
		}
	}
	return res
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	grid, start, movement := parse(filename)

	large := enlarge(grid)

	// Part 1
	robot := start
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

	// Part 2
	robot = space.Point{X: start.X * 2, Y: start.Y}
	for _, m := range movement {
		if move(large, robot.Add(m), m) {
			robot = robot.Add(m)
		}
	}
	sum = 0
	for p, r := range large {
		if r != '[' {
			continue
		}

		sum += p.X + 100*p.Y
	}
	log.Part2(sum)
}
