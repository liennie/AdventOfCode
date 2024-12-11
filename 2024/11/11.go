package main

import (
	"strconv"

	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/ints"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
)

func parse(filename string) []int {
	return evil.Fields(load.Line(filename))
}

var blinkRuleBuf = make([]int, 2)

// uses global buffer, cannot be called concurrently
func blinkRule(stone int) []int {
	if stone == 0 {
		blinkRuleBuf[0] = 1
		return blinkRuleBuf[:1]
	}

	if s := strconv.Itoa(stone); len(s)%2 == 0 {
		p := ints.Pow(10, len(s)/2)

		blinkRuleBuf[0] = stone / p
		blinkRuleBuf[1] = stone % p
		return blinkRuleBuf
	}

	blinkRuleBuf[0] = stone * 2024
	return blinkRuleBuf[:1]
}

func blink(stones map[int]int, res map[int]int) {
	for n, count := range stones {
		for _, m := range blinkRule(n) {
			res[m] += count
		}
	}
}

func sum(stones map[int]int) int {
	s := 0
	for _, count := range stones {
		s += count
	}
	return s
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	stones := parse(filename)

	stoneMap := map[int]int{}
	for _, stone := range stones {
		stoneMap[stone]++
	}

	resMap := map[int]int{}

	// Part 1
	blinks := 25
	for range blinks {
		blink(stoneMap, resMap)
		stoneMap, resMap = resMap, stoneMap
		clear(resMap)
	}
	log.Part1(sum(stoneMap))

	// Part 2
	blinks = 50 // 25 + 50 = 75 blinks total
	for range blinks {
		blink(stoneMap, resMap)
		stoneMap, resMap = resMap, stoneMap
		clear(resMap)
	}
	log.Part2(sum(stoneMap))
}
