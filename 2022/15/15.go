package main

import (
	"regexp"

	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/ints"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
	"github.com/liennie/AdventOfCode/pkg/set"
	"github.com/liennie/AdventOfCode/pkg/space"
)

type sensor struct {
	pos    space.Point
	beacon space.Point
}

var inputRe = regexp.MustCompile(`^Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)$`)

func parse(filename string) []sensor {
	res := []sensor{}
	for line := range load.File(filename) {
		match := inputRe.FindStringSubmatch(line)
		if match == nil {
			evil.Panic("Line %q does not match", line)
		}

		res = append(res, sensor{
			pos: space.Point{
				X: evil.Atoi(match[1]),
				Y: evil.Atoi(match[2]),
			},
			beacon: space.Point{
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

	sensors := parse(filename)

	// Part 1
	const line = 2000000
	rs := set.RangeSet{}
	for _, sensor := range sensors {
		d := sensor.beacon.Sub(sensor.pos).ManhattanLen()
		h := ints.Abs(line - sensor.pos.Y)
		rs.Add(set.Range{
			Min: sensor.pos.X - d + h,
			Max: sensor.pos.X + d - h,
		})
	}
	for _, sensor := range sensors {
		if sensor.beacon.Y == line {
			rs.Remove(set.Range{Min: sensor.beacon.X, Max: sensor.beacon.X})
		}
	}
	log.Part1(rs.Len())
}
