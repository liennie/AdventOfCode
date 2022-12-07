package main

import (
	"strings"

	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
)

func part1(filename string) (int, int) {
	pos, depth := 0, 0

	for line := range load.File(filename) {
		switch {
		case strings.HasPrefix(line, "forward"):
			pos += evil.Atoi(line[len("forward")+1:])

		case strings.HasPrefix(line, "down"):
			depth += evil.Atoi(line[len("down")+1:])

		case strings.HasPrefix(line, "up"):
			depth -= evil.Atoi(line[len("up")+1:])
		}
	}

	return pos, depth
}

func part2(filename string) (int, int) {
	aim, pos, depth := 0, 0, 0

	for line := range load.File(filename) {
		switch {
		case strings.HasPrefix(line, "forward"):
			forward := evil.Atoi(line[len("forward")+1:])
			pos += forward
			depth += aim * forward

		case strings.HasPrefix(line, "down"):
			aim += evil.Atoi(line[len("down")+1:])

		case strings.HasPrefix(line, "up"):
			aim -= evil.Atoi(line[len("up")+1:])
		}
	}

	return pos, depth
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	// Part 1
	pos, depth := part1(filename)
	log.Part1(pos * depth)

	// Part 2
	pos, depth = part2(filename)
	log.Part2(pos * depth)
}
