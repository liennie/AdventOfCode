package main

import (
	"math"
	"strings"

	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
	"github.com/liennie/AdventOfCode/pkg/space"
)

type Hailstone struct {
	pos space.Point3
	vel space.Point3
}

func parsePoint(raw string) space.Point3 {
	n := evil.SplitN(raw, ",", 3)
	evil.Assert(len(n) == 3, "invalid coords")
	return space.Point3{
		X: n[0],
		Y: n[1],
		Z: n[2],
	}
}

func parse(filename string) []Hailstone {
	res := []Hailstone{}
	for line := range load.File(filename) {
		pos, vel, ok := strings.Cut(line, "@")
		evil.Assert(ok, "invalid format")
		res = append(res, Hailstone{
			pos: parsePoint(pos),
			vel: parsePoint(vel),
		})
	}
	return res
}

type Line2 struct {
	k, c float64
}

func newLine2(pos, vel space.Point3) Line2 {
	if vel.X < 0 {
		vel = vel.Flip()
	}

	return Line2{
		k: float64(vel.Y) / float64(vel.X),
		c: float64(pos.Y) - float64(vel.Y)*(float64(pos.X)/float64(vel.X)),
	}
}

func almostEquals(a, b float64) bool {
	return math.Abs(a-b) < math.Pow(10, math.Log10(math.Abs(a))-9)
}

func intersect2(a, b Line2) (float64, float64) {
	if a.k == b.k {
		return math.Inf(1), math.Inf(1)
	}

	x := (b.c - a.c) / (a.k - b.k)
	y := x*a.k + a.c
	y2 := x*b.k + b.c
	evil.Assert(almostEquals(y, y2), a, b, x, y, y2, y-y2)
	return x, y
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	hailstones := parse(filename)

	const from = 200000000000000
	const to = 400000000000000
	// const from = 7
	// const to = 27

	// Part 1
	cnt := 0
	for i := 0; i < len(hailstones)-1; i++ {
		a := hailstones[i]
		for j := i + 1; j < len(hailstones); j++ {
			b := hailstones[j]

			x, y := intersect2(newLine2(a.pos, a.vel), newLine2(b.pos, b.vel))
			if x >= from && x <= to && y >= from && y <= to {
				if ((a.vel.X > 0) == (x > float64(a.pos.X))) && ((b.vel.X > 0) == (x > float64(b.pos.X))) {
					cnt++
				}
			}
		}
	}
	log.Part1(cnt)
}
