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

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	patterns := parse(filename)

	// Part 1
	sum := 0
patterns:
	for _, pattern := range patterns {
	horizontal:
		for mid := range pattern {
			if mid == 0 { // XXX reflection on the edge? 🤔
				continue
			}

			for offset := 0; offset < mid && offset < len(pattern)-mid; offset++ {
				for j := range pattern[mid] {
					if pattern[mid-1-offset][j] != pattern[mid+offset][j] {
						continue horizontal
					}
				}
			}

			sum += 100 * mid
			continue patterns
		}

	vertical:
		for mid := range pattern[0] {
			if mid == 0 { // XXX reflection on the edge? 🤔
				continue
			}

			for offset := 0; offset < mid && offset < len(pattern[0])-mid; offset++ {
				for j := range pattern {
					if pattern[j][mid-1-offset] != pattern[j][mid+offset] {
						continue vertical
					}
				}
			}

			sum += mid
			continue patterns
		}
	}
	log.Part1(sum)
}
