package main

import (
	"strings"

	"github.com/liennie/AdventOfCode/common/load"
	"github.com/liennie/AdventOfCode/common/log"
	"github.com/liennie/AdventOfCode/common/set"
	"github.com/liennie/AdventOfCode/common/util"
)

type rucksack struct {
	compartments []set.String
}

func parse(filename string) []rucksack {
	res := []rucksack{}
	for line := range load.File(filename) {
		items := strings.Split(line, "")
		res = append(res, rucksack{
			compartments: []set.String{
				set.New(items[:len(line)/2]...),
				set.New(items[len(line)/2:]...),
			},
		})
	}
	return res
}

func priority(item string) int {
	if len(item) != 1 {
		util.Panic("Invalid item len %d", len(item))
	}

	c := item[0]
	if c >= 'a' && c <= 'z' {
		return int(c-'a') + 1
	} else if c >= 'A' && c <= 'Z' {
		return int(c-'A') + 27
	}

	util.Panic("Invalid item %c", c)
	return 0
}

func main() {
	defer util.Recover(log.Err)

	const filename = "input.txt"

	rucksacks := parse(filename)

	// Part 1
	sum := 0
	for _, rucksack := range rucksacks {
		for item := range set.Intersection(rucksack.compartments...) {
			sum += priority(item)
		}
	}
	log.Part1(sum)

	// Part 2
	sum = 0
	for i := 2; i < len(rucksacks); i += 3 {
		for item := range set.Intersection(
			set.Union(rucksacks[i-2].compartments...),
			set.Union(rucksacks[i-1].compartments...),
			set.Union(rucksacks[i].compartments...),
		) {
			sum += priority(item)
		}
	}
	log.Part2(sum)
}
