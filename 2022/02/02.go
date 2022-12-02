package main

import (
	"strings"

	"github.com/liennie/AdventOfCode/common/load"
	"github.com/liennie/AdventOfCode/common/log"
	"github.com/liennie/AdventOfCode/common/util"
)

type shape int

const (
	rock     shape = 1
	paper    shape = 2
	scissors shape = 3
)

func (s shape) score(op shape) int {
	if s == 0 || op == 0 {
		util.Panic("Zero shape")
	}

	score := int(s)

	if s == op {
		score += 3
	} else if (s == rock && op == scissors) ||
		(s == paper && op == rock) ||
		(s == scissors && op == paper) {
		score += 6
	}

	return score
}

func parseShape(in string) shape {
	switch in {
	case "A":
		return rock
	case "B":
		return paper
	case "C":
		return scissors
	}
	util.Panic("Invalid shape %s", in)
	return 0
}

type strategy int

const (
	strategyX strategy = iota
	strategyY
	strategyZ
)

func parseStrategy(in string) strategy {
	switch in {
	case "X":
		return strategyX
	case "Y":
		return strategyY
	case "Z":
		return strategyZ
	}
	util.Panic("Invalid strategy %s", in)
	return 0
}

type round struct {
	opponent shape
	player   strategy
}

func parse(filename string) []round {
	res := []round{}

	for line := range load.File(filename) {
		op, pl, _ := strings.Cut(line, " ")
		res = append(res, round{
			opponent: parseShape(op),
			player:   parseStrategy(pl),
		})
	}

	return res
}

func score(rounds []round, tr func(round) shape) int {
	score := 0

	for _, round := range rounds {
		score += tr(round).score(round.opponent)
	}

	return score
}

var guessedMap = map[strategy]shape{
	strategyX: rock,
	strategyY: paper,
	strategyZ: scissors,
}

func guessedStrategy(r round) shape {
	return guessedMap[r.player]
}

var correctMap = map[round]shape{
	{player: strategyX, opponent: rock}:     scissors,
	{player: strategyX, opponent: paper}:    rock,
	{player: strategyX, opponent: scissors}: paper,
	{player: strategyY, opponent: rock}:     rock,
	{player: strategyY, opponent: paper}:    paper,
	{player: strategyY, opponent: scissors}: scissors,
	{player: strategyZ, opponent: rock}:     paper,
	{player: strategyZ, opponent: paper}:    scissors,
	{player: strategyZ, opponent: scissors}: rock,
}

func correctStrategy(r round) shape {
	return correctMap[r]
}

func main() {
	defer util.Recover(log.Err)

	const filename = "input.txt"

	rounds := parse(filename)

	// Part 1
	log.Part1(score(rounds, guessedStrategy))

	// Part 2
	log.Part2(score(rounds, correctStrategy))
}
