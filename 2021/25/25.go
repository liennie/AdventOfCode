package main

import (
	"github.com/liennie/AdventOfCode/common/load"
	"github.com/liennie/AdventOfCode/common/log"
	"github.com/liennie/AdventOfCode/common/util"
)

func parse(filename string) (map[util.Point]byte, util.Point) {
	res := map[util.Point]byte{}

	y := 0
	mx := 0
	for line := range load.File(filename) {
		for x := 0; x < len(line); x++ {
			if line[x] == '>' || line[x] == 'v' {
				res[util.Point{X: x, Y: y}] = line[x]
			}
		}
		mx = util.Max(mx, len(line))
		y++
	}

	return res, util.Point{X: mx, Y: y}
}

func step(cucumbers map[util.Point]byte, max util.Point) (map[util.Point]byte, bool) {
	next := map[util.Point]byte{}
	moved := false

	for p, c := range cucumbers {
		if c == '>' {
			to := util.Point{X: (p.X + 1) % max.X, Y: p.Y}
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
	next = map[util.Point]byte{}

	for p, c := range cucumbers {
		if c == 'v' {
			to := util.Point{X: p.X, Y: (p.Y + 1) % max.Y}
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

func simulate(cucumbers map[util.Point]byte, max util.Point) int {
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
	defer util.Recover(log.Err)

	const filename = "input.txt"

	cucumbers, max := parse(filename)

	// Part 1
	log.Part1(simulate(cucumbers, max))
}
