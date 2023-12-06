package main

import (
	"strings"

	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
)

type Race struct {
	time     int
	distance int
}

func parse(filename string) []Race {
	ch := load.File(filename)
	times := evil.Fields(strings.TrimPrefix(<-ch, "Time:"))
	distances := evil.Fields(strings.TrimPrefix(<-ch, "Distance:"))

	evil.Assert(len(times) == len(distances), "invalid input")

	res := make([]Race, len(times))
	for i := range times {
		res[i].time = times[i]
		res[i].distance = distances[i]
	}
	return res
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	races := parse(filename)

	// Part 1
	prod := 1
	for _, race := range races {
		cnt := 0
		for n := 0; n < race.time; n++ {
			dist := n * (race.time - n)
			if dist > race.distance {
				cnt++
			}
		}
		prod *= cnt
	}
	log.Part1(prod)
}
