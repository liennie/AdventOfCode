package main

import (
	"container/list"
	"math"
	"strings"

	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/ints"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
	"github.com/liennie/AdventOfCode/pkg/space"
)

type cube struct {
	min, max space.Point3
}

func rangeIntersects(aMin, aMax, bMin, bMax int) (int, int, bool) {
	if aMin > bMin {
		aMin, aMax, bMin, bMax = bMin, bMax, aMin, aMax
	}

	if aMax >= bMin {
		return bMin, ints.Min(aMax, bMax), true
	}

	return 0, -1, false
}

func (c cube) intersects(other cube) (cube, bool) {
	res := cube{}
	var okx, oky, okz bool

	res.min.X, res.max.X, okx = rangeIntersects(c.min.X, c.max.X, other.min.X, other.max.X)
	res.min.Y, res.max.Y, oky = rangeIntersects(c.min.Y, c.max.Y, other.min.Y, other.max.Y)
	res.min.Z, res.max.Z, okz = rangeIntersects(c.min.Z, c.max.Z, other.min.Z, other.max.Z)

	return res, okx && oky && okz
}

func (c cube) remove(other cube) ([]cube, bool) {
	cutout, ok := c.intersects(other)
	if !ok {
		return []cube{c}, false
	}

	res := []cube{}

	for _, split := range []cube{
		{
			min: space.Point3{X: math.MinInt, Y: math.MinInt, Z: math.MinInt},
			max: space.Point3{X: cutout.min.X - 1, Y: math.MaxInt, Z: math.MaxInt},
		},
		{
			min: space.Point3{X: cutout.max.X + 1, Y: math.MinInt, Z: math.MinInt},
			max: space.Point3{X: math.MaxInt, Y: math.MaxInt, Z: math.MaxInt},
		},
		{
			min: space.Point3{X: cutout.min.X, Y: math.MinInt, Z: math.MinInt},
			max: space.Point3{X: cutout.max.X, Y: cutout.min.Y - 1, Z: math.MaxInt},
		},
		{
			min: space.Point3{X: cutout.min.X, Y: cutout.max.Y + 1, Z: math.MinInt},
			max: space.Point3{X: cutout.max.X, Y: math.MaxInt, Z: math.MaxInt},
		},
		{
			min: space.Point3{X: cutout.min.X, Y: cutout.min.Y, Z: math.MinInt},
			max: space.Point3{X: cutout.max.X, Y: cutout.max.Y, Z: cutout.min.Z - 1},
		},
		{
			min: space.Point3{X: cutout.min.X, Y: cutout.min.Y, Z: cutout.max.Z + 1},
			max: space.Point3{X: cutout.max.X, Y: cutout.max.Y, Z: math.MaxInt},
		},
	} {
		if in, ok := c.intersects(split); ok {
			res = append(res, in)
		}
	}

	return res, true
}

func (c cube) size() int {
	return (c.max.X - c.min.X + 1) * (c.max.Y - c.min.Y + 1) * (c.max.Z - c.min.Z + 1)
}

type step struct {
	cube cube
	on   bool
}

func parse(filename string) []step {
	res := []step{}
	for line := range load.File(filename) {
		s := step{}

		if strings.HasPrefix(line, "on") {
			s.on = true
			line = strings.TrimPrefix(line, "on ")
		} else {
			line = strings.TrimPrefix(line, "off ")
		}

		for _, c := range strings.SplitN(line, ",", 3) {
			p := strings.SplitN(c, "=", 2)
			m := ints.SplitN(p[1], "..", 2)

			switch p[0] {
			case "x":
				s.cube.min.X = ints.Min(m[0], m[1])
				s.cube.max.X = ints.Max(m[0], m[1])
			case "y":
				s.cube.min.Y = ints.Min(m[0], m[1])
				s.cube.max.Y = ints.Max(m[0], m[1])
			case "z":
				s.cube.min.Z = ints.Min(m[0], m[1])
				s.cube.max.Z = ints.Max(m[0], m[1])
			}
		}

		res = append(res, s)
	}
	return res
}

func count(cubes *list.List, limit cube) int {
	total := 0

	for c := cubes.Front(); c != nil; c = c.Next() {
		if lc, ok := c.Value.(cube).intersects(limit); ok {
			total += lc.size()
		}
	}

	return total
}

func reboot(steps []step) *list.List {
	res := list.New()

	for _, step := range steps {
		var next *list.Element
		for c := res.Front(); c != nil; c = next {
			next = c.Next()

			if newCubes, ok := c.Value.(cube).remove(step.cube); ok {
				res.Remove(c)
				for _, nc := range newCubes {
					res.PushBack(nc)
				}
			}
		}

		if step.on {
			res.PushBack(step.cube)
		}
	}

	return res
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	steps := parse(filename)
	cubes := reboot(steps)

	// Part 1
	log.Part1(count(cubes, cube{
		min: space.Point3{X: -50, Y: -50, Z: -50},
		max: space.Point3{X: 50, Y: 50, Z: 50},
	}))

	// Part 2
	log.Part2(count(cubes, cube{
		min: space.Point3{X: math.MinInt, Y: math.MinInt, Z: math.MinInt},
		max: space.Point3{X: math.MaxInt, Y: math.MaxInt, Z: math.MaxInt},
	}))
}
