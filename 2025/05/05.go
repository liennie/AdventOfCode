package main

import (
	"strings"

	"github.com/liennie/AdventOfCode/pkg/channel"
	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
	"github.com/liennie/AdventOfCode/pkg/set"
)

func parse(filename string) (set.RangeSet, []int) {
	ch := load.Blocks(filename)
	defer channel.Drain(ch)

	ranges := set.RangeSet{}
	for line := range <-ch {
		min, max, ok := strings.Cut(line, "-")
		evil.Assert(ok, "invalid format %q", line)
		ranges.Add(set.Range{Min: evil.Atoi(min), Max: evil.Atoi(max)})
	}

	ids := []int{}
	for line := range <-ch {
		ids = append(ids, evil.Atoi(line))
	}

	return ranges, ids
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	ranges, ids := parse(filename)

	// Part 1
	cnt := 0
	for _, id := range ids {
		if ranges.Contains(id) {
			cnt++
		}
	}
	log.Part1(cnt)

	// Part 2
	log.Part2(ranges.Len())
}
