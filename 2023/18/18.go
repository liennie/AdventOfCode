package main

import (
	"fmt"
	"image/color"

	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
	"github.com/liennie/AdventOfCode/pkg/set"
	"github.com/liennie/AdventOfCode/pkg/space"
)

type Plan struct {
	dir   space.Point
	l     int
	color color.RGBA
}

func parse(filename string) []Plan {
	res := []Plan{}
	for line := range load.File(filename) {
		var d string
		var n int
		c := color.RGBA{
			A: 255,
		}
		_, err := fmt.Sscanf(line, "%s %d (#%02x%02x%02x)", &d, &n, &c.R, &c.G, &c.B)
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
		p.color = c

		res = append(res, p)
	}
	return res
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	plans := parse(filename)

	// Part 1
	dug := set.Set[space.Point]{}
	dir := map[space.Point]space.Point{}
	aabb := space.AABB{}
	pos := space.Point{}
	for _, plan := range plans {
		aabb = aabb.Add(pos)
		for i := 0; i < plan.l; i++ {
			dug.Add(pos)
			if plan.dir.Y != 0 {
				dir[pos] = plan.dir
			}
			pos = pos.Add(plan.dir)
		}
		if plan.dir.Y != 0 {
			dir[pos] = plan.dir
		}
	}
	evil.Assert(pos == space.Point{})
	for y := aabb.Min.Y; y <= aabb.Max.Y; y++ {
		prev := space.Point{}
		dig := false
		for x := aabb.Min.X; x <= aabb.Max.X; x++ {
			if d := dir[space.Point{X: x, Y: y}]; d != (space.Point{}) && d != prev {
				dig = !dig
				prev = d
			}
			if dig {
				dug.Add(space.Point{X: x, Y: y})
			}
		}
	}
	log.Part1(len(dug))
}
