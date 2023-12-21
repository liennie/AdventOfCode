package main

import (
	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
	"github.com/liennie/AdventOfCode/pkg/set"
	"github.com/liennie/AdventOfCode/pkg/space"
)

func parse(filename string) [][]byte {
	res := [][]byte{}
	for line := range load.File(filename) {
		res = append(res, []byte(line))
	}
	return res
}

func findStart(garden [][]byte) space.Point {
	for y := range garden {
		for x := range garden[y] {
			if garden[y][x] == 'S' {
				garden[y][x] = '.'
				return space.Point{X: x, Y: y}
			}
		}
	}
	evil.Panic("start not found")
	return space.Point{}
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	garden := parse(filename)
	start := findStart(garden)

	// Part 1
	plots := set.New(start)
	for n := 0; n < 64; n++ {
		next := set.New[space.Point]()
		for pos := range plots {
			for _, dir := range []space.Point{{X: 1}, {Y: -1}, {X: -1}, {Y: 1}} {
				n := pos.Add(dir)
				if garden[n.Y][n.X] != '#' {
					next.Add(n)
				}
			}
		}
		plots = next
	}
	log.Part1(len(plots))
}
