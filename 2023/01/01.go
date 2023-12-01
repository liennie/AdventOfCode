package main

import (
	"strings"

	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
)

var validDigits = map[string]int{
	"0":     0,
	"1":     1,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	lines := load.Slice(filename)

	// Part 1
	sum := 0

	for _, line := range lines {
		first, last := 0, 0
		ok := false

		for _, c := range line {
			if c >= '0' && c <= '9' {
				digit := int(c - '0')

				if !ok {
					first = digit
					ok = true
				}
				last = digit
			}
		}

		sum += first*10 + last
	}
	log.Part1(sum)

	// Part 2
	sum = 0
	for _, line := range lines {
		first, last := 0, 0
		firstIndex, lastIndex := len(line), -1

		for digit, value := range validDigits {
			if index := strings.Index(line, digit); index >= 0 && index < firstIndex {
				firstIndex = index
				first = value
			}
			if index := strings.LastIndex(line, digit); index >= 0 && index > lastIndex {
				lastIndex = index
				last = value
			}
		}

		sum += first*10 + last
	}
	log.Part2(sum)
}
