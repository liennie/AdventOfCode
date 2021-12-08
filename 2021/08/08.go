package main

import (
	"sort"
	"strings"

	"github.com/liennie/AdventOfCode/common/load"
	"github.com/liennie/AdventOfCode/common/log"
	"github.com/liennie/AdventOfCode/common/util"
)

type entry struct {
	patterns []string
	digits   []string
}

func parse(filename string) []entry {
	res := []entry{}

	for line := range load.File(filename) {
		parts := strings.SplitN(line, " | ", 2)
		res = append(res, entry{
			patterns: strings.Split(parts[0], " "),
			digits:   strings.Split(parts[1], " "),
		})
	}

	return res
}

func sortSegments(s string) string {
	b := []byte(s)
	sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
	return string(b)
}

func xorSegments(a, b string) string {
	res := []rune{}
	for r := 'a'; r <= 'g'; r++ {
		if strings.ContainsRune(a, r) != strings.ContainsRune(b, r) {
			res = append(res, r)
		}
	}
	return string(res)
}

func disambiguate(e entry) int {
	digits := [10]string{}
	ambiguous := []string{}

	for _, pattern := range e.patterns {
		switch len(pattern) {
		case 2:
			digits[1] = pattern
		case 3:
			digits[7] = pattern
		case 4:
			digits[4] = pattern
		case 7:
			digits[8] = pattern
		default:
			ambiguous = append(ambiguous, pattern)
		}
	}

	for _, pattern := range ambiguous {
		if x1 := len(xorSegments(pattern, digits[1])); x1 == 3 {
			digits[3] = pattern
		} else if x1 == 6 {
			digits[6] = pattern
		} else if x4 := len(xorSegments(pattern, digits[4])); x4 == 2 {
			digits[9] = pattern
		} else if x4 == 5 {
			digits[2] = pattern
		} else if len(pattern) == 5 {
			digits[5] = pattern
		} else {
			digits[0] = pattern
		}
	}

	m := map[string]int{}
	for i, pattern := range digits {
		if len(pattern) == 0 {
			util.Panic("We missed %d", i)
		}
		m[sortSegments(pattern)] = i
	}

	res := 0
	for _, digit := range e.digits {
		d, ok := m[sortSegments(digit)]
		if !ok {
			util.Panic("Unknown digit %q, %+v", d, digits)
		}

		res *= 10
		res += d
	}
	return res
}

func main() {
	defer util.Recover(log.Err)

	const filename = "input.txt"

	entries := parse(filename)

	// Part 1
	count := 0
	for _, entry := range entries {
		for _, digit := range entry.digits {
			l := len(digit)
			if l == 2 || l == 3 || l == 4 || l == 7 {
				count++
			}
		}
	}
	log.Part1(count)

	// Part 2
	total := 0
	for _, entry := range entries {
		total += disambiguate(entry)
	}
	log.Part2(total)
}
