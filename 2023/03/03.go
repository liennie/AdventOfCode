package main

import (
	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
	"github.com/liennie/AdventOfCode/pkg/space"
)

type PartNumber struct {
	num int
	pos space.AABB
}

type Symbol struct {
	symbol rune
	pos    space.Point
}

func parse(filename string) ([]PartNumber, []Symbol) {
	numbers := []PartNumber{}
	symbols := []Symbol{}

	y := 0
	for line := range load.File(filename) {
		isNum := false
		num := 0
		numPos := space.AABB{}

		for x, c := range line {
			if c >= '0' && c <= '9' {
				isNum = true
				num = num*10 + int(c-'0')
				numPos = numPos.Add(space.Point{X: x, Y: y})
			} else {
				if isNum {
					numbers = append(numbers, PartNumber{
						num: num,
						pos: numPos,
					})

					isNum = false
					num = 0
					numPos = space.AABB{}
				}

				if c != '.' {
					symbols = append(symbols, Symbol{
						symbol: c,
						pos:    space.Point{X: x, Y: y},
					})
				}
			}
		}

		if isNum {
			numbers = append(numbers, PartNumber{
				num: num,
				pos: numPos,
			})
		}

		y++
	}

	return numbers, symbols
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	numbers, symbols := parse(filename)

	// Part 1
	sum := 0
numbers:
	for _, number := range numbers {
		for _, symbol := range symbols {
			if number.pos.Expand(1).Contains(symbol.pos) {
				sum += number.num
				continue numbers
			}
		}
	}
	log.Part1(sum)

	// Part 2
	sum = 0
	for _, symbol := range symbols {
		if symbol.symbol != '*' {
			continue
		}

		prod := 1
		cnt := 0
		for _, number := range numbers {
			if number.pos.Expand(1).Contains(symbol.pos) {
				prod *= number.num
				cnt++
			}
		}

		if cnt == 2 {
			sum += prod
		}
	}
	log.Part2(sum)
}
