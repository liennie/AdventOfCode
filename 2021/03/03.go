package main

import (
	"container/list"
	"strconv"
	"strings"

	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
)

func count(values *list.List) []map[byte]int {
	counts := []map[byte]int{}

	for node := values.Front(); node != nil; node = node.Next() {
		value := node.Value.(string)

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

func rating(values *list.List, criteria func(map[byte]int) byte) int {
	i := 0

	for values.Len() > 1 {
		counts := count(values)
		b := criteria(counts[i])
		var next *list.Element
		for node := values.Front(); node != nil; node = next {
			next = node.Next()

			value := node.Value.(string)
			if value[i] != b {
				values.Remove(node)
			}
		}
		i++
	}

	return parseBin(values.Front().Value.(string))
}

func oxygenRating(values *list.List) int {
	return rating(values, mostCommon)
}

func co2Rating(values *list.List) int {
	return rating(values, leastCommon)
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	// Part 1
	counts := count(load.List(filename))
	gamma, epsilon := rates(counts)
	log.Part1(gamma * epsilon)

	// Part 2
	oxygen := oxygenRating(load.List(filename))
	co2 := co2Rating(load.List(filename))
	log.Part2(oxygen * co2)
}
