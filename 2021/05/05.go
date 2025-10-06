package main

import (
	"strings"

	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
	"github.com/liennie/AdventOfCode/pkg/space"
)

type line struct {
	start, end space.Point
}

func (l line) isHorizontal() bool {
	return l.start.Y == l.end.Y
}

func (l line) isVertical() bool {
	return l.start.X == l.end.X
}

func (l line) dir() space.Point {
	return l.end.Sub(l.start).Norm()
}

func parse(filename string) []line {
	res := []line{}

	for l := range load.File(filename) {
		points := strings.SplitN(l, " -> ", 2)
		start := evil.SplitN(points[0], ",", 2)
		end := evil.SplitN(points[1], ",", 2)

		res = append(res, line{
			start: space.Point{
				X: start[0],
				Y: start[1],
			},
			end: space.Point{
				X: end[0],
				Y: end[1],
			},
		})
	}

	return res
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	lines := parse(filename)

	// Part 1
	diag := map[space.Point]int{}
	for _, l := range lines {
		if l.isHorizontal() || l.isVertical() {
			dir := l.dir()

			for p := l.start; p != l.end; p = p.Add(dir) {
				diag[p]++
			}
			diag[l.end]++
		}
	}

	count := 0
	for _, c := range diag {
		if c >= 2 {
			count++
		}
	}

	log.Part1(count)

	// Part 2
	diag = map[space.Point]int{}
	for _, l := range lines {
		dir := l.dir()

		for p := l.start; p != l.end; p = p.Add(dir) {
			diag[p]++
		}
		diag[l.end]++
	}

	count = 0
	for _, c := range diag {
		if c >= 2 {
			count++
		}
	}

	log.Part2(count)
}
