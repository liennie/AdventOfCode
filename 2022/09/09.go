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

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	motions := parse(filename)

	// Part 1
	var head, tail space.Point
	visited := set.New(tail)
	for _, motion := range motions {
		for i := 0; i < motion.dist; i++ {
			head = head.Add(motion.dir)

			v := head.Sub(tail)
			if ints.Abs(v.X) >= 2 || ints.Abs(v.Y) >= 2 {
				tail = tail.Add(space.Point{
					X: ints.Clamp(v.X, -1, 1),
					Y: ints.Clamp(v.Y, -1, 1),
				})

				visited.Add(tail)
			}
		}
	}
	log.Part1(len(visited))
}
