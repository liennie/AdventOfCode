package main

import (
	"github.com/liennie/AdventOfCode/common/load"
	"github.com/liennie/AdventOfCode/common/log"
	"github.com/liennie/AdventOfCode/common/util"
)

func get(filename string) string {
	ch := load.File(filename)
	defer util.Drain(ch)
	return <-ch
}

func isUniq(seq string) bool {
	for i := 0; i < len(seq)-1; i++ {
		for j := i + 1; j < len(seq); j++ {
			if seq[i] == seq[j] {
				return false
			}
		}
	}
	return true
}

func main() {
	defer util.Recover(log.Err)

	const filename = "input.txt"

	signal := get(filename)

	// Part 1
	const sop = 4
	for i := sop; i < len(signal); i++ {
		if isUniq(signal[i-sop : i]) {
			log.Part1(i)
			break
		}
	}
}
