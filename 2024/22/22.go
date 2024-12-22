package main

import (
	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/ints"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
)

func parse(filename string) []int {
	return load.Parse(filename, func(line string) int {
		return evil.Atoi(line)
	})
}

func next(secret int) int {
	// this can all be done using bitwise ops
	secret ^= secret * 64
	secret %= 16777216
	secret ^= secret / 32
	secret %= 16777216
	secret ^= secret * 2048
	secret %= 16777216
	return secret
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	secrets := parse(filename)

	// Part 1
	for i := range secrets {
		for range 2000 {
			secrets[i] = next(secrets[i])
		}
	}
	log.Part1(ints.Sum(secrets...))
}
