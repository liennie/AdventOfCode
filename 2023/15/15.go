package main

import (
	"strings"

	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
)

func parse(filename string) []string {
	res := []string{}
	for line := range load.File(filename) {
		res = append(res, strings.Split(line, ",")...)
	}
	return res
}

func hash(s string) int {
	cur := byte(0)
	for _, c := range s {
		evil.Assert(c < 128, "non ASCII character")

		cur += byte(c)
		cur *= 17
	}
	return int(cur)
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	seq := parse(filename)

	// Part 1
	sum := 0
	for _, s := range seq {
		sum += hash(s)
	}
	log.Part1(sum)
}
