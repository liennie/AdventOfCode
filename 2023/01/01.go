package main

import (
	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
)

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	lines := load.Slice(filename)

	// Part 1
	sum := 0

	for _, line := range lines {
		first, last := 0, 0
		ok := false

		for _, c := range line {
			if c >= '0' && c <= '9' {
				digit := int(c - '0')

				if !ok {
					first = digit
					ok = true
				}
				last = digit
			}
		}

		sum += first*10 + last
	}
	log.Part1(sum)
}
