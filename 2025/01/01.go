package main

import (
	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/ints"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
)

func parse(filename string) []int {
	return load.Parse(filename, func(line string) int {
		switch line[0] {
		case 'L':
			return -evil.Atoi(line[1:])
		case 'R':
			return evil.Atoi(line[1:])
		default:
			evil.Panic("wrong format %q", line)
			return 0
		}
	})
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	rots := parse(filename)

	// Part 1
	cnt := 0
	cur := 50
	for _, n := range rots {
		cur = ints.Mod(cur+n, 100)
		if cur == 0 {
			cnt++
		}
	}
	log.Part1(cnt)

	// Part 2
	cnt = 0
	cur = 50
	for _, n := range rots {
		zero := cur == 0
		cur += n
		if cur >= 100 {
			cnt += cur / 100
			cur = ints.Mod(cur, 100)
		} else if cur <= 0 {
			cnt += (100 - cur) / 100
			cur = ints.Mod(cur, 100)
			if zero {
				cnt-- // don't double count zeros
			}
		}
	}
	log.Part2(cnt)
}
