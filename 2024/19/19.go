package main

import (
	"cmp"
	"slices"
	"strings"

	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
)

func parse(filename string) (patterns []string, designs []string) {
	ch := load.File(filename)
	patterns = strings.Split(<-ch, ",")
	for i := range patterns {
		patterns[i] = strings.TrimSpace(patterns[i])
	}

	<-ch // empty line

	for line := range ch {
		designs = append(designs, line)
	}

	return
}

var cache = map[string]int{}

func designPossible(design string, patterns []string) (possible int) {
	if possible, ok := cache[design]; ok {
		return possible
	}
	defer func() {
		cache[design] = possible
	}()

	if len(design) == 0 {
		return 1
	}

	for _, pattern := range patterns {
		if strings.HasPrefix(design, pattern) {
			possible += designPossible(design[len(pattern):], patterns)
		}
	}

	return
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	patterns, designs := parse(filename)

	slices.SortFunc(patterns, func(a, b string) int {
		return -cmp.Compare(len(a), len(b))
	})

	// Part 1
	count := 0
	for _, design := range designs {
		if designPossible(design, patterns) > 0 {
			count++
		}
	}
	log.Part1(count)

	// Part 2
	count = 0
	for _, design := range designs {
		count += designPossible(design, patterns)
	}
	log.Part2(count)
}
