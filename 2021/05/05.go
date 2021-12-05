package main

import (
	"strings"

	"github.com/liennie/AdventOfCode/common/load"
	"github.com/liennie/AdventOfCode/common/log"
	"github.com/liennie/AdventOfCode/common/util"
)

type point struct {
	x, y int
}

func (p point) normalize() point {
	gcd := util.GCD(util.Abs(p.x), util.Abs(p.y))
	return point{
		x: p.x / gcd,
		y: p.y / gcd,
	}
}

func (p point) sub(other point) point {
	return point{
		x: p.x - other.x,
		y: p.y - other.y,
	}
}

func (p point) add(other point) point {
	return point{
		x: p.x + other.x,
		y: p.y + other.y,
	}
}

func (p point) equals(other point) bool {
	return p.x == other.x && p.y == other.y
}

type line struct {
	start, end point
}

func (l line) isHorizontal() bool {
	return l.start.y == l.end.y
}

func (l line) isVertical() bool {
	return l.start.x == l.end.x
}

func (l line) dir() point {
	return l.end.sub(l.start).normalize()
}

func parse(filename string) []line {
	res := []line{}

	for l := range load.File(filename) {
		points := strings.SplitN(l, " -> ", 2)
		start := strings.SplitN(points[0], ",", 2)
		end := strings.SplitN(points[1], ",", 2)

		res = append(res, line{
			start: point{
				x: util.Atoi(start[0]),
				y: util.Atoi(start[1]),
			},
			end: point{
				x: util.Atoi(end[0]),
				y: util.Atoi(end[1]),
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
	diag := map[point]int{}
	for _, l := range lines {
		if l.isHorizontal() || l.isVertical() {
			dir := l.dir()

			for p := l.start; !p.equals(l.end); p = p.add(dir) {
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
	diag = map[point]int{}
	for _, l := range lines {
		dir := l.dir()

		for p := l.start; !p.equals(l.end); p = p.add(dir) {
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
