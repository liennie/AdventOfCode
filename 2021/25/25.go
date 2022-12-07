package main

import (
	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/ints"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
	"github.com/liennie/AdventOfCode/pkg/space"
)

func parse(filename string) (map[space.Point]byte, space.Point) {
	res := map[space.Point]byte{}

	y := 0
	mx := 0
	for line := range load.File(filename) {
		for x := 0; x < len(line); x++ {
			if line[x] == '>' || line[x] == 'v' {
				res[space.Point{X: x, Y: y}] = line[x]
			}
		}
		mx = ints.Max(mx, len(line))
		y++
	}

	return res, space.Point{X: mx, Y: y}
}

func step(cucumbers map[space.Point]byte, max space.Point) (map[space.Point]byte, bool) {
	next := map[space.Point]byte{}
	moved := false

	for p, c := range cucumbers {
		if c == '>' {
			to := space.Point{X: (p.X + 1) % max.X, Y: p.Y}
			if cucumbers[to] == 0 {
				moved = true
				next[to] = c
			} else {
				next[p] = c
			}
		} else if c != 0 {
			next[p] = c
		}
	}

	cucumbers = next
	next = map[space.Point]byte{}

	for p, c := range cucumbers {
		if c == 'v' {
			to := space.Point{X: p.X, Y: (p.Y + 1) % max.Y}
			if cucumbers[to] == 0 {
				moved = true
				next[to] = c
			} else {
				next[p] = c
			}
		} else if c != 0 {
			next[p] = c
		}
	}

	return next, moved
}

func simulate(cucumbers map[space.Point]byte, max space.Point) int {
	moves := 0
	moved := false
	for {
		moves++
		cucumbers, moved = step(cucumbers, max)
		if !moved {
			break
		}
	}
	return moves
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	cucumbers, max := parse(filename)

	// Part 1
	log.Part1(simulate(cucumbers, max))
}
