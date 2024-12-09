package main

import (
	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/ints"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
)

func parse(filename string) []int {
	return evil.Split(load.Line(filename), "")
}

func createMemory(blocks []int) []int {
	memory := make([]int, 0, ints.Sum(blocks...))
	for i, n := range blocks {
		isFile := i%2 == 0
		id := i / 2

		for range n {
			if isFile {
				memory = append(memory, id)
			} else {
				memory = append(memory, -1)
			}
		}
	}
	return memory
}

func checksum(memory []int) int {
	total := 0
	for i, n := range memory {
		if n == -1 {
			break
		}

		total += i * n
	}
	return total
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	blocks := parse(filename)

	// Part 1
	memory := createMemory(blocks)
	for i, j := 0, len(memory)-1; i < j; {
		switch {
		case memory[i] != -1:
			i++

		case memory[j] == -1:
			j--

		default:
			// memory[i] == -1
			// memory[j] != -1

			memory[i], memory[j] = memory[j], memory[i]
		}
	}
	log.Part1(checksum(memory))
}
