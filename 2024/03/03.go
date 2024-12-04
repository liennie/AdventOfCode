package main

import (
	"regexp"

	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
)

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	memory := load.Slice(filename)

	// Part 1
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	sum := 0
	for _, line := range memory {
		for _, match := range re.FindAllStringSubmatch(line, -1) {
			sum += evil.Atoi(match[1]) * evil.Atoi(match[2])
		}
	}
	log.Part1(sum)

	// Part 2
	re = regexp.MustCompile(`(mul)\((\d{1,3}),(\d{1,3})\)|(do)\(\)|(don't)\(\)`)
	sum = 0
	enabled := true
	for _, line := range memory {
		for _, match := range re.FindAllStringSubmatch(line, -1) {
			switch {
			case match[1] == "mul":
				if enabled {
					sum += evil.Atoi(match[2]) * evil.Atoi(match[3])
				}

			case match[4] == "do":
				enabled = true

			case match[5] == "don't":
				enabled = false
			}

		}
	}
	log.Part2(sum)
}
