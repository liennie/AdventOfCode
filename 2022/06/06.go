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

func findMarker(mlen int, signal string) int {
	for i := mlen; i < len(signal); i++ {
		if isUniq(signal[i-mlen : i]) {
			return i
		}
	}
	util.Panic("Marker of len %d not found", mlen)
	return 0
}

func main() {
	defer util.Recover(log.Err)

	const filename = "input.txt"

	signal := get(filename)

	// Part 1
	log.Part1(findMarker(4, signal))

	// Part 2
	log.Part2(findMarker(14, signal))
}
