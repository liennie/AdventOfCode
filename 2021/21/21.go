package main

import (
	"math"
	"strings"

	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/ints"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
)

func parse(filename string) (int, int) {
	ch := load.File(filename)

	return ints.Atoi(strings.TrimPrefix(<-ch, "Player 1 starting position: ")),
		ints.Atoi(strings.TrimPrefix(<-ch, "Player 2 starting position: "))
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

type gameState struct {
	score  [2]int
	pos    [2]int
	player int
}

func move(from, by int) int {
	return (((from + by) - 1) % 10) + 1
}

func (st gameState) move(by int) gameState {
	st.pos[st.player] = move(st.pos[st.player], by)
	st.score[st.player] += st.pos[st.player]
	st.player = 1 - st.player

	return st
}

func playDeterministic(aStart, bStart int) stats {
	die := &deterministicDie{}
	gameState := gameState{
		pos: [2]int{aStart, bStart},
	}
	stats := stats{}

	for gameState.score[0] < 1000 && gameState.score[1] < 1000 {
		gameState = gameState.move(die.roll() + die.roll() + die.roll())
		stats.rolls += 3
	}

	stats.winnerScore = gameState.score[1-gameState.player]
	stats.loserScore = gameState.score[gameState.player]

	return stats
}

func playQuantum(aStart, bStart int) stats {
	stats := stats{}

	scores := [2]int{}

	states := map[gameState]int{
		{pos: [2]int{aStart, bStart}}: 1,
	}

	for len(states) > 0 {
		minScore := math.MaxInt
		minState := gameState{}

		for state := range states {
			if score := state.score[0] + state.score[1]; score < minScore {
				minScore = score
				minState = state
			}
		}

		count := states[minState]
		delete(states, minState)

		for first := 1; first <= 3; first++ {
			for second := 1; second <= 3; second++ {
				for third := 1; third <= 3; third++ {
					newState := minState.move(first + second + third)
					if newState.score[0] >= 21 {
						scores[0] += count
					} else if newState.score[1] >= 21 {
						scores[1] += count
					} else {
						states[newState] += count
					}
				}
			}
		}
	}

	var victor int
	if scores[0] > scores[1] {
		victor = 0
	} else {
		victor = 1
	}

	stats.winnerScore = scores[victor]
	stats.loserScore = scores[1-victor]

	return stats
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	aStart, bStart := parse(filename)

	// Part 1
	stats := playDeterministic(aStart, bStart)
	log.Part1(stats.loserScore * stats.rolls)

	// Part 2
	stats = playQuantum(aStart, bStart)
	log.Part2(stats.winnerScore)
}
