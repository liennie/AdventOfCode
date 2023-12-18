package main

import (
	"fmt"

	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
	"github.com/liennie/AdventOfCode/pkg/space"
)

type Plan struct {
	dir space.Point
	l   int
}

func parse1(filename string) []Plan {
	res := []Plan{}
	for line := range load.File(filename) {
		var d string
		var n int

		_, err := fmt.Sscanf(line, "%s %d", &d, &n)
		evil.Err(err)

		p := Plan{}

		switch d {
		case "U":
			p.dir = space.Point{Y: -1}
		case "D":
			p.dir = space.Point{Y: 1}
		case "L":
			p.dir = space.Point{X: -1}
		case "R":
			p.dir = space.Point{X: 1}
		default:
			evil.Panic("invalid dir %q", d)
		}

		p.l = n

		res = append(res, p)
	}
	return res
}

func parse2(filename string) []Plan {
	res := []Plan{}
	for line := range load.File(filename) {
		var d string
		var n int
		var hn int
		var hd int

		_, err := fmt.Sscanf(line, "%s %d (#%05x%d)", &d, &n, &hn, &hd)
		evil.Err(err)

		p := Plan{}

		switch hd {
		case 3:
			p.dir = space.Point{Y: -1}
		case 1:
			p.dir = space.Point{Y: 1}
		case 2:
			p.dir = space.Point{X: -1}
		case 0:
			p.dir = space.Point{X: 1}
		default:
			evil.Panic("invalid dir %d", hd)
		}

		p.l = hn

		res = append(res, p)
	}
	return res
}

func dig(plans []Plan) int {
	total := 0
	tl := 0
	x := 0
	pos := space.Point{}
	for _, plan := range plans {
		l := plan.dir.Scale(plan.l)
		pos = pos.Add(l)
		x += l.X
		total += x * l.Y
		tl += plan.l
	}
	evil.Assert(pos == space.Point{})

	total += tl/2 + 1

	return total
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	log.Part1(dig(parse1(filename)))
	log.Part2(dig(parse2(filename)))
}
