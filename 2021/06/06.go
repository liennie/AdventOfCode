package main

import (
	"github.com/liennie/AdventOfCode/pkg/channel"
	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/ints"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
)

func parse(filename string) []int {
	ch := load.File(filename)
	defer channel.Drain(ch)
	return evil.Split(<-ch, ",")
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	fish := parse(filename)

	// Part 1
	for d := 0; d < 80; d++ {
		l := len(fish)
		for i := 0; i < l; i++ {
			fish[i]--
			if fish[i] < 0 {
				fish[i] = 6
				fish = append(fish, 8)
			}
		}
	}
	log.Part1(len(fish))

	// Part 2
	fishCount := make([]int, 9)
	for _, f := range fish {
		fishCount[f]++
	}

	for d := 80; d < 256; d++ {
		new := fishCount[0]
		for i := 1; i < 9; i++ {
			fishCount[i-1] = fishCount[i]
		}
		fishCount[6] += new
		fishCount[8] = new
	}

	log.Part2(ints.Sum(fishCount...))
}
