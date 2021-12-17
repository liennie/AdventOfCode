package main

import (
	"strings"

	"github.com/liennie/AdventOfCode/common/load"
	"github.com/liennie/AdventOfCode/common/log"
	"github.com/liennie/AdventOfCode/common/util"
)

func parse(filename string) (util.Point, util.Point) {
	ch := load.File(filename)
	defer util.Drain(ch)

	line := strings.TrimPrefix(<-ch, "target area: ")
	coords := strings.SplitN(line, ", ", 2)

	var min, max util.Point
	for _, coord := range coords {
		if strings.HasPrefix(coord, "x=") {
			rang := util.SplitN(coord[2:], "..", 2)
			min.X = util.Min(rang[0], rang[1])
			max.X = util.Max(rang[0], rang[1])
		} else {
			rang := util.SplitN(coord[2:], "..", 2)
			min.Y = util.Min(rang[0], rang[1])
			max.Y = util.Max(rang[0], rang[1])
		}
	}

	return min, max
}

func main() {
	defer util.Recover(log.Err)

	const filename = "test.txt"

	min, max := parse(filename)

	// Part 1
	if max.Y < 0 {
		log.Part1((min.Y * (min.Y + 1)) / 2)
	} else if min.Y > 0 {
		log.Part1((max.Y * (max.Y + 1)) / 2)
	} else {
		log.Part1("Inf")
	}

	// Part 2
}
