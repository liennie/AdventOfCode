package main

import (
	"github.com/liennie/AdventOfCode/pkg/channel"
	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
)

func get(filename string) string {
	ch := load.File(filename)
	defer channel.Drain(ch)
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
	evil.Panic("Marker of len %d not found", mlen)
	return 0
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	signal := get(filename)

	// Part 1
	log.Part1(findMarker(4, signal))

	// Part 2
	log.Part2(findMarker(14, signal))
}
