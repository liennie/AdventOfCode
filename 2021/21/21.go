package main

import (
	"strings"

	"github.com/liennie/AdventOfCode/common/load"
	"github.com/liennie/AdventOfCode/common/log"
	"github.com/liennie/AdventOfCode/common/util"
)

func parse(filename string) (int, int) {
	ch := load.File(filename)

	return util.Atoi(strings.TrimPrefix(<-ch, "Player 1 starting position: ")),
		util.Atoi(strings.TrimPrefix(<-ch, "Player 2 starting position: "))
}

type die interface {
	roll() int
}

type deterministicDie struct {
	v int
}

func (d *deterministicDie) roll() int {
	d.v++
	if d.v > 100 {
		d.v = 1
	}

	return d.v
}

type stats struct {
	winnerScore int
	loserScore  int
	rolls       int
}

func play(aStart, bStart int, die die) stats {
	score := [2]int{}
	pos := [2]int{aStart, bStart}
	player := 0
	stats := stats{}

	for score[0] < 1000 && score[1] < 1000 {
		pos[player] += die.roll() + die.roll() + die.roll()
		pos[player] = ((pos[player] - 1) % 10) + 1
		score[player] += pos[player]
		player = 1 - player

		stats.rolls += 3
	}

	stats.winnerScore = score[1-player]
	stats.loserScore = score[player]

	return stats
}

func main() {
	defer util.Recover(log.Err)

	const filename = "input.txt"

	aStart, bStart := parse(filename)

	// Part 1
	stats := play(aStart, bStart, &deterministicDie{})
	log.Part1(stats.loserScore * stats.rolls)
}
