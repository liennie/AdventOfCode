package main

import (
	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
)

func parse(filename string) [][][]byte {
	res := [][][]byte{}

	for block := range load.Blocks(filename) {
		pattern := [][]byte{}

		for line := range block {
			pattern = append(pattern, []byte(line))
		}

		res = append(res, pattern)
	}

	return res
}

func reflectionValue(pattern [][]byte, useSmudge bool) int {
horizontal:
	for mid := range pattern {
		if mid == 0 { // XXX reflection on the edge? 🤔
			continue
		}

		smudged := false

		for offset := 0; offset < mid && offset < len(pattern)-mid; offset++ {
			for j := range pattern[mid] {
				if pattern[mid-1-offset][j] != pattern[mid+offset][j] {
					if useSmudge && !smudged {
						smudged = true
						continue
					}
					continue horizontal
				}
			}
		}

		if !useSmudge || smudged {
			return 100 * mid
		}
	}

vertical:
	for mid := range pattern[0] {
		if mid == 0 { // XXX reflection on the edge? 🤔
			continue
		}

		smudged := false

		for offset := 0; offset < mid && offset < len(pattern[0])-mid; offset++ {
			for j := range pattern {
				if pattern[j][mid-1-offset] != pattern[j][mid+offset] {
					if useSmudge && !smudged {
						smudged = true
						continue
					}
					continue vertical
				}
			}
		}

		if !useSmudge || smudged {
			return mid
		}
	}

	return 0
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	patterns := parse(filename)

	// Part 1
	sum := 0
	for _, pattern := range patterns {
		sum += reflectionValue(pattern, false)
	}
	log.Part1(sum)

	// Part 2
	sum = 0
	for _, pattern := range patterns {
		sum += reflectionValue(pattern, true)
	}
	log.Part2(sum)
}
