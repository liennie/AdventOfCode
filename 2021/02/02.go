package main

import (
	"strings"

	"github.com/liennie/AdventOfCode/common/load"
	"github.com/liennie/AdventOfCode/common/log"
	"github.com/liennie/AdventOfCode/common/util"
)

func parse(filename string) (int, int) {
	pos, depth := 0, 0

	for line := range load.File(filename) {
		switch {
		case strings.HasPrefix(line, "forward"):
			pos += util.Atoi(line[len("forward")+1:])

		case strings.HasPrefix(line, "down"):
			depth += util.Atoi(line[len("down")+1:])

		case strings.HasPrefix(line, "up"):
			depth -= util.Atoi(line[len("up")+1:])
		}
	}

	return pos, depth
}

func main() {
	defer util.Recover(log.Err)

	const filename = "input.txt"

	pos, depth := parse(filename)

	// Part 1
	log.Part1(pos * depth)
}
