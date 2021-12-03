package main

import (
	"strconv"
	"strings"

	"github.com/liennie/AdventOfCode/common/load"
	"github.com/liennie/AdventOfCode/common/log"
	"github.com/liennie/AdventOfCode/common/set"
	"github.com/liennie/AdventOfCode/common/util"
)

func count(values set.String) []map[byte]int {
	counts := []map[byte]int{}

	for value := range values {
		for i := 0; i < len(value); i++ {
			b := value[i]

			if i >= len(counts) {
				counts = append(counts, make([]map[byte]int, i-len(counts)+1)...)
			}
			if counts[i] == nil {
				counts[i] = map[byte]int{}
			}

			counts[i][b]++
		}
	}

	return counts
}

func parseBin(s string) int {
	i, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		panic(err)
	}
	return int(i)
}

func mostCommon(counts map[byte]int) byte {
	max := 0
	var maxB byte
	for b, c := range counts {
		if c == max {
			if b > maxB {
				maxB = b
			}
		} else if c > max {
			max = c
			maxB = b
		}
	}
	return maxB
}

func leastCommon(counts map[byte]int) byte {
	min := int(^uint(0) >> 1)
	var minB byte
	for b, c := range counts {
		if c == min {
			if b < minB {
				minB = b
			}
		} else if c < min {
			min = c
			minB = b
		}
	}
	return minB
}

func rates(counts []map[byte]int) (int, int) {
	gamma := &strings.Builder{}
	epsilon := &strings.Builder{}

	for _, m := range counts {
		gamma.WriteByte(mostCommon(m))
		epsilon.WriteByte(leastCommon(m))
	}

	return parseBin(gamma.String()), parseBin(epsilon.String())
}

func rating(values set.String, criteria func(map[byte]int) byte) int {
	i := 0

	for len(values) > 1 {
		counts := count(values)
		b := criteria(counts[i])
		for value := range values {
			if value[i] != b {
				values.Remove(value)
			}
		}
		i++
	}

	for value := range values {
		return parseBin(value)
	}
	panic("Empty set")
}

func oxygenRating(values set.String) int {
	return rating(values, mostCommon)
}

func co2Rating(values set.String) int {
	return rating(values, leastCommon)
}

func main() {
	defer util.Recover(log.Err)

	const filename = "input.txt"

	// Part 1
	counts := count(load.Set(filename))
	gamma, epsilon := rates(counts)
	log.Part1(gamma * epsilon)

	// Part 2
	oxygen := oxygenRating(load.Set(filename))
	co2 := co2Rating(load.Set(filename))
	log.Part2(oxygen * co2)
}
