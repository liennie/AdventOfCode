package main

import (
	"strings"

	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/ints"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
	"github.com/liennie/AdventOfCode/pkg/set"
	"github.com/liennie/AdventOfCode/pkg/space"
)

type motion struct {
	dir  space.Point
	dist int
}

var dirMap = map[string]space.Point{
	"R": {X: 1, Y: 0},
	"L": {X: -1, Y: 0},
	"U": {X: 0, Y: 1},
	"D": {X: 0, Y: -1},
}

func parse(filename string) []motion {
	res := []motion{}
	for line := range load.File(filename) {
		dir, dist, _ := strings.Cut(line, " ")
		res = append(res, motion{
			dir:  dirMap[dir],
			dist: evil.Atoi(dist),
		})
	}
	return res
}

func move(knots []space.Point, dir space.Point) {
	knots[0] = knots[0].Add(dir)

	for i := 1; i < len(knots); i++ {
		v := knots[i-1].Sub(knots[i])
		if ints.Abs(v.X) >= 2 || ints.Abs(v.Y) >= 2 {
			knots[i] = knots[i].Add(space.Point{
				X: ints.Clamp(v.X, -1, 1),
				Y: ints.Clamp(v.Y, -1, 1),
			})
		}
	}
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	motions := parse(filename)

	// Part 1
	knots := make([]space.Point, 2)
	visited := set.New(space.Point{})
	for _, motion := range motions {
		for i := 0; i < motion.dist; i++ {
			move(knots, motion.dir)
			visited.Add(knots[1])
		}
	}
	log.Part1(len(visited))

	// Part 2
	knots = make([]space.Point, 10)
	visited = set.New(space.Point{})
	for _, motion := range motions {
		for i := 0; i < motion.dist; i++ {
			move(knots, motion.dir)
			visited.Add(knots[9])
		}
	}
	log.Part2(len(visited))
}
