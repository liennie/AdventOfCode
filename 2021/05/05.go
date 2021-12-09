package main

import (
	"strings"

	"github.com/liennie/AdventOfCode/common/load"
	"github.com/liennie/AdventOfCode/common/log"
	"github.com/liennie/AdventOfCode/common/util"
)

type line struct {
	start, end util.Point
}

func (l line) isHorizontal() bool {
	return l.start.Y == l.end.Y
}

func (l line) isVertical() bool {
	return l.start.X == l.end.X
}

func (l line) dir() util.Point {
	return l.end.Sub(l.start).Normalize()
}

func parse(filename string) []line {
	res := []line{}

	for l := range load.File(filename) {
		points := strings.SplitN(l, " -> ", 2)
		start := util.SplitN(points[0], ",", 2)
		end := util.SplitN(points[1], ",", 2)

		res = append(res, line{
			start: util.Point{
				X: start[0],
				Y: start[1],
			},
			end: util.Point{
				X: end[0],
				Y: end[1],
			},
		})
	}

	return res
}

func main() {
	defer util.Recover(log.Err)

	const filename = "input.txt"

	lines := parse(filename)

	// Part 1
	diag := map[util.Point]int{}
	for _, l := range lines {
		if l.isHorizontal() || l.isVertical() {
			dir := l.dir()

			for p := l.start; !p.Equals(l.end); p = p.Add(dir) {
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
	diag = map[util.Point]int{}
	for _, l := range lines {
		dir := l.dir()

		for p := l.start; !p.Equals(l.end); p = p.Add(dir) {
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
