package main

import (
	"strings"

	"github.com/liennie/AdventOfCode/common/load"
	"github.com/liennie/AdventOfCode/common/log"
	"github.com/liennie/AdventOfCode/common/util"
)

type board struct {
	numbers [][]int
	marked  [][]bool
}

func newBoard() board {
	b := board{}

	b.numbers = make([][]int, 5)
	for i := range b.numbers {
		b.numbers[i] = make([]int, 5)
	}

	b.marked = make([][]bool, 5)
	for i := range b.marked {
		b.marked[i] = make([]bool, 5)
	}

	return b
}

func (b *board) mark(n int) {
	for i, line := range b.numbers {
		for j, num := range line {
			if n == num {
				b.marked[i][j] = true
			}
		}
	}
}

func (b *board) won() bool {
	for i := 0; i < 5; i++ {
		col := true
		row := true

		for j := 0; j < 5; j++ {
			if !b.marked[i][j] {
				col = false
			}
			if !b.marked[j][i] {
				row = false
			}
			if !col && !row {
				break
			}
		}

		if row || col {
			return true
		}
	}

	return false
}

func (b *board) score() int {
	total := 0

	for i, line := range b.numbers {
		for j, num := range line {
			if !b.marked[i][j] {
				total += num
			}
		}
	}

	return total
}

func (b *board) reset() {
	for i, line := range b.marked {
		for j := range line {
			b.marked[i][j] = false
		}
	}
}

func parseLine(line string) []int {
	res := make([]int, 5)
	for i := 0; i < 5; i++ {
		res[i] = util.Atoi(strings.TrimSpace(line[i*3 : i*3+2]))
	}
	return res
}

func parse(filename string) ([]int, []board) {
	ch := load.File(filename)

	numbers := util.SliceAtoi(strings.Split(<-ch, ","))
	<-ch // Empty line

	boards := []board{}

boards:
	for {
		b := newBoard()

		for i := 0; i < 5; i++ {
			line, ok := <-ch
			if !ok {
				break boards
			}

			b.numbers[i] = parseLine(line)
		}
		<-ch // Empty line

		boards = append(boards, b)
	}

	return numbers, boards
}

func main() {
	defer util.Recover(log.Err)

	const filename = "input.txt"

	numbers, boards := parse(filename)

	// Part 1
numbers:
	for _, num := range numbers {
		for _, board := range boards {
			board.mark(num)
			if board.won() {
				log.Part1(board.score() * num)
				break numbers
			}
		}
	}

	// Part 2
	for _, board := range boards {
		board.reset()
	}

	won := map[int]bool{}

numbers2:
	for _, num := range numbers {
		for i, board := range boards {
			board.mark(num)
			if board.won() {
				won[i] = true
			}

			if len(won) == len(boards) {
				log.Part1(board.score() * num)
				break numbers2
			}
		}
	}
}
