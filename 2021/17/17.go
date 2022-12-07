package main

import (
	"strings"

	"github.com/liennie/AdventOfCode/pkg/channel"
	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/ints"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
	"github.com/liennie/AdventOfCode/pkg/space"
)

func parse(filename string) (space.Point, space.Point) {
	ch := load.File(filename)
	defer channel.Drain(ch)

	line := strings.TrimPrefix(<-ch, "target area: ")
	coords := strings.SplitN(line, ", ", 2)

	var min, max space.Point
	for _, coord := range coords {
		if strings.HasPrefix(coord, "x=") {
			rang := evil.SplitN(coord[2:], "..", 2)
			min.X = ints.Min(rang[0], rang[1])
			max.X = ints.Max(rang[0], rang[1])
		} else {
			rang := evil.SplitN(coord[2:], "..", 2)
			min.Y = ints.Min(rang[0], rang[1])
			max.Y = ints.Max(rang[0], rang[1])
		}
	}

	return min, max
}

func valid(vel, min, max space.Point) bool {
	pos := space.Point{X: 0, Y: 0}
	for pos.X <= max.X && pos.Y >= min.Y {
		if pos.X >= min.X && pos.Y <= max.Y {
			return true
		}

		pos = pos.Add(vel)
		if vel.X > 0 {
			vel.X--
		} else if vel.X < 0 {
			vel.X++
		}
		vel.Y--
	}

	return false
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

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
	count := 0
	maxX := max.X
	minX := 0
	var maxY, minY int
	if max.Y < 0 {
		maxY = -min.Y - 1
		minY = min.Y
	} else if min.Y > 0 {
		maxY = max.Y
		minY = 0
	} else {
		log.Part2("Inf")
		return
	}

	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			if valid(space.Point{X: x, Y: y}, min, max) {
				count++
			}
		}
	}
	log.Part2(count)
}
