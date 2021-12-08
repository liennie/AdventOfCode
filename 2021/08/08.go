package main

import (
	"strings"

	"github.com/liennie/AdventOfCode/common/load"
	"github.com/liennie/AdventOfCode/common/log"
	"github.com/liennie/AdventOfCode/common/util"
)

type entry struct {
	patterns []string
	digits   []string
}

func parse(filename string) []entry {
	res := []entry{}

	for line := range load.File(filename) {
		parts := strings.SplitN(line, " | ", 2)
		res = append(res, entry{
			patterns: strings.Split(parts[0], " "),
			digits:   strings.Split(parts[1], " "),
		})
	}

	return res
}

func main() {
	defer util.Recover(log.Err)

	const filename = "input.txt"

	entries := parse(filename)

	// Part 1
	count := 0
	for _, entry := range entries {
		for _, digit := range entry.digits {
			l := len(digit)
			if l == 2 || l == 3 || l == 4 || l == 7 {
				count++
			}
		}
	}
	log.Part1(count)
}
