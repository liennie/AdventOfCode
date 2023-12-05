package main

import (
	"math"
	"strings"

	"github.com/liennie/AdventOfCode/pkg/channel"
	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
	"github.com/liennie/AdventOfCode/pkg/set"
)

var mapIndex = map[string]int{
	"seed-to-soil":            0,
	"soil-to-fertilizer":      1,
	"fertilizer-to-water":     2,
	"water-to-light":          3,
	"light-to-temperature":    4,
	"temperature-to-humidity": 5,
	"humidity-to-location":    6,
}

func parse(filename string) ([]int, []map[set.Range]int) {
	ch := load.Blocks(filename)

	seedch := <-ch
	seeds := evil.Split(strings.TrimPrefix(<-seedch, "seeds: "), " ")
	channel.Drain(seedch)

	maps := make([]map[set.Range]int, len(mapIndex))

	for block := range ch {
		name := strings.TrimSuffix(<-block, " map:")

		i, ok := mapIndex[name]
		evil.Assert(ok, "invalid map ", name)

		m := map[set.Range]int{}
		maps[i] = m

		for line := range block {
			n := evil.Split(line, " ")
			evil.Assert(len(n) == 3, "invalid mapping ", n)

			m[set.Range{Min: n[1], Max: n[1] + n[2] - 1}] = n[0]
		}
	}
	for i, m := range maps {
		evil.Assert(m != nil, "missing map ", i)
	}

	return seeds, maps
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	seeds, maps := parse(filename)

	// Part 1
	min := math.MaxInt
	for _, seed := range seeds {
		num := seed
	maps:
		for _, m := range maps {
			for r, start := range m {
				if r.Contains(num) {
					num = start + num - r.Min
					continue maps
				}
			}
		}

		if num < min {
			min = num
		}
	}
	log.Part1(min)
}
