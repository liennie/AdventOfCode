package main

import (
	"strings"

	"github.com/liennie/AdventOfCode/common/load"
	"github.com/liennie/AdventOfCode/common/log"
	"github.com/liennie/AdventOfCode/common/util"
)

func parse(filename string) []int {
	ch := load.File(filename)

	defer func() {
		for range <-ch {
		}
	}()

	return util.SliceAtoi(strings.Split(<-ch, ","))
}

func main() {
	defer util.Recover(log.Err)

	const filename = "input.txt"

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

	total := 0
	for _, c := range fishCount {
		total += c
	}
	log.Part2(total)
}
