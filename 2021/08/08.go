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

func doMagic(e entry) int {
	digits := [10]string{}
	fives := []string{}
	sixes := []string{}

	for _, pattern := range e.patterns {
		switch len(pattern) {
		case 2:
			digits[1] = sortSegments(pattern)
		case 3:
			digits[7] = sortSegments(pattern)
		case 4:
			digits[4] = sortSegments(pattern)
		case 7:
			digits[8] = sortSegments(pattern)
		case 5:
			fives = append(fives, sortSegments(pattern))
		case 6:
			sixes = append(sixes, sortSegments(pattern))
		}
	}

	for _, f := range fives {
		if x := xorSegments(f, digits[7]); len(x) == 2 {
			digits[3] = f
		} else if x := xorSegments(f, digits[4]); len(x) == 5 {
			digits[2] = f
		} else {
			digits[5] = f
		}
	}

	for _, s := range sixes {
		if x := xorSegments(s, digits[4]); len(x) == 2 {
			digits[9] = s
		} else if x := xorSegments(s, digits[1]); len(x) == 6 {
			digits[6] = s
		} else {
			digits[0] = s
		}
	}

	m := map[string]int{}
	for i, pattern := range digits {
		if len(pattern) == 0 {
			util.Panic("We missed %d", i)
		}
		m[pattern] = i
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
		total += doMagic(entry)
	}
	log.Part2(total)
}
