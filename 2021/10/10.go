package main

import (
	"sort"

	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
)

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	lines := load.Slice(filename)

	bmap := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
		'<': '>',
	}
	ssmap := map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}
	asmap := map[rune]int{
		')': 1,
		']': 2,
		'}': 3,
		'>': 4,
	}

	syntaxScore := 0
	autoScores := []int{}
lines:
	for _, line := range lines {
		stack := []rune{}
		for _, r := range line {
			switch r {
			case '(', '[', '{', '<':
				stack = append(stack, bmap[r])
			case ')', ']', '}', '>':
				if len(stack) == 0 || stack[len(stack)-1] != r {
					syntaxScore += ssmap[r]
					continue lines
				}
				stack = stack[:len(stack)-1]
			default:
				evil.Panic("Invalid rune %c", r)
			}
		}

		autoScore := 0
		for i := len(stack) - 1; i >= 0; i-- {
			autoScore *= 5
			autoScore += asmap[stack[i]]
		}
		autoScores = append(autoScores, autoScore)
	}
	log.Part1(syntaxScore)

	sort.Ints(autoScores)
	log.Part2(autoScores[len(autoScores)/2])
}
