package main

import (
	"regexp"

	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/ints"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
	"github.com/liennie/AdventOfCode/pkg/space"
)

type Robot struct {
	pos, vel space.Point
}

func parse(filename string) []Robot {
	re := regexp.MustCompile(`^p=(-?\d+),(-?\d+) v=(-?\d+),(-?\d+)$`)

	var res []Robot
	for line := range load.File(filename) {
		match := re.FindStringSubmatch(line)
		evil.Assert(match != nil, "line '", line, "' does not match regexp '", re.String(), "'")

		res = append(res, Robot{
			pos: space.Point{
				X: evil.Atoi(match[1]),
				Y: evil.Atoi(match[2]),
			},
			vel: space.Point{
				X: evil.Atoi(match[3]),
				Y: evil.Atoi(match[4]),
			},
		})
	}
	return res
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	robots := parse(filename)

	max := space.Point{X: 11, Y: 7}
	if filename == "input.txt" {
		max = space.Point{X: 101, Y: 103}
	}

	// Part 1
	t := 100
	for i := range robots {
		robot := &robots[i]
		robot.pos = robot.pos.Add(robot.vel.Scale(t)).Mod(max)
	}

	mid := max.Div(space.Point{X: 2, Y: 2})
	f := [4]int{}
	for _, robot := range robots {
		if robot.pos.X == mid.X || robot.pos.Y == mid.Y {
			continue
		}

		i := 0
		if robot.pos.X > mid.X {
			i |= 1
		}
		if robot.pos.Y > mid.Y {
			i |= 2
		}

		f[i]++
	}
	log.Part1(ints.Product(f[:]...))
}
