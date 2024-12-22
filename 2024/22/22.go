package main

import (
	"fmt"
	"slices"

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
	secret ^= secret << 6  // secret * 64
	secret &= 16777216 - 1 // % 16777216
	secret ^= secret >> 5  // secret / 32
	secret &= 16777216 - 1 // % 16777216
	secret ^= secret << 11 // secret * 2048
	secret &= 16777216 - 1 // % 16777216
	return secret
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	secrets := parse(filename)

	const gen = 2000

	// Part 1
	secretsCpy := slices.Clone(secrets)
	for i := range secretsCpy {
		for range gen {
			secretsCpy[i] = next(secretsCpy[i])
		}
	}
	log.Part1(ints.Sum(secretsCpy...))

	// Part 2
	const seqLen = 4

	total := map[string]int{}
	first := map[string]int{}
	changeBuf := make([]int, 0, gen)
	for i := range secrets {
		clear(first)
		changes := changeBuf[:0]
		for range gen {
			prev := secrets[i] % 10
			secrets[i] = next(secrets[i])
			diff := secrets[i]%10 - prev

			changes = append(changes, diff)
			if len(changes) > seqLen {
				changes = changes[len(changes)-seqLen:]
			}
			if len(changes) == seqLen {
				key := fmt.Sprint(changes)
				if _, ok := first[key]; !ok {
					first[key] = secrets[i] % 10
				}
			}
		}
		for key, val := range first {
			total[key] += val
		}
	}
	max := 0
	maxSeq := ""
	for seq, val := range total {
		if val > max {
			max = val
			maxSeq = seq
		}
	}
	log.Print(maxSeq)
	log.Part2(max)
}
