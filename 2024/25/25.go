package main

import (
	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
)

type lock [5]int
type key [5]int

func parse(filename string) ([]key, []lock) {
	var keys []key
	var locks []lock
	for schema := range load.Blocks(filename) {
		lastLine := ""
		pat := [5]int{-1, -1, -1, -1, -1}
		for line := range schema {
			for i, c := range line {
				if c == '#' {
					pat[i]++
				}
			}
			lastLine = line
		}
		switch lastLine {
		case "#####":
			keys = append(keys, pat)
		case ".....":
			locks = append(locks, pat)
		default:
			evil.Panic("invalid last line %q", lastLine)
		}
	}
	return keys, locks
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	keys, locks := parse(filename)

	// Part 1
	total := 0
	for _, key := range keys {
	locks:
		for _, lock := range locks {
			for i := range key {
				if key[i]+lock[i] > 5 {
					continue locks
				}
			}
			total++
		}
	}
	log.Part1(total)
}
