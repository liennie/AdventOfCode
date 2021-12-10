package main

import (
	"github.com/liennie/AdventOfCode/common/load"
	"github.com/liennie/AdventOfCode/common/log"
	"github.com/liennie/AdventOfCode/common/util"
)

func main() {
	defer util.Recover(log.Err)

	const filename = "input.txt"

	lines := load.Slice(filename)

	bmap := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
		'<': '>',
	}
	smap := map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}

	// Part 1
	score := 0
lines:
	for _, line := range lines {
		stack := []rune{}
		for _, r := range line {
			switch r {
			case '(', '[', '{', '<':
				stack = append(stack, bmap[r])
			case ')', ']', '}', '>':
				if len(stack) == 0 || stack[len(stack)-1] != r {
					score += smap[r]
					continue lines
				}
				stack = stack[:len(stack)-1]
			default:
				util.Panic("Invalid rune %c", r)
			}
		}
	}
	log.Part1(score)
}
