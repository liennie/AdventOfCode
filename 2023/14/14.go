package main

import (
	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
)

func parse(filename string) [][]byte {
	res := [][]byte{}

	for line := range load.File(filename) {
		res = append(res, []byte(line))
	}

	return res
}

func rollNorth(rocks [][]byte) {
	for x := range rocks[0] {
		space := 0
		canRoll := false
		for y := range rocks {
			switch rocks[y][x] {
			case '.':
				if !canRoll {
					space = y
					canRoll = true
				}
			case '#':
				space = y + 1
				canRoll = false
			case 'O':
				rocks[space][x], rocks[y][x] = rocks[y][x], rocks[space][x]
				space++
			}
		}
	}
}

func print(rocks [][]byte) {
	for _, line := range rocks {
		log.Printf("%c", line)
	}
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	rocks := parse(filename)

	// Part 1
	rollNorth(rocks)
	sum := 0
	for y := range rocks {
		for x := range rocks[y] {
			if rocks[y][x] == 'O' {
				sum += len(rocks) - y
			}
		}
	}
	log.Part1(sum)
}
