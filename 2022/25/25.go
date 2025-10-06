package main

import (
	"slices"
	"strconv"
	"strings"

	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/ints"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
)

var SNAFUVals = map[rune]int{
	'=': -2,
	'-': -1,
	'0': 0,
	'1': 1,
	'2': 2,
}

var SNAFUChars = map[int]rune{
	-2: '=',
	-1: '-',
	0:  '0',
	1:  '1',
	2:  '2',
}

func SNAFU2int(snafu string) int {

	res := 0
	for _, r := range snafu {
		if val, ok := SNAFUVals[r]; ok {
			res *= 5
			res += val
		} else {
			evil.Panic("invalid SNAFU char %c", r)
		}
	}
	return res
}

func Int2SNAFU(i int) string {
	evil.Assert(i >= 0, "negative i: %d", i)

	nn := evil.Split(strconv.FormatInt(int64(i), 5), "")
	slices.Reverse(nn)
	nn = append(nn, 0)
	for i, n := range nn {
		switch n {
		case 3:
			nn[i+1]++
			nn[i] = -2
		case 4:
			nn[i+1]++
			nn[i] = -1
		case 5:
			nn[i+1]++
			nn[i] = 0
		}
	}
	slices.Reverse(nn)

	res := &strings.Builder{}
	for _, n := range nn {
		if r, ok := SNAFUChars[n]; ok {
			res.WriteRune(r)
		} else {
			evil.Panic("invalid SNAFU val %d, in %d, nn %v", n, i, nn)
		}
	}

	return strings.TrimLeft(res.String(), "0")
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	lines := load.Slice(filename)

	// Part 1
	log.Part1(Int2SNAFU(ints.SumFunc(SNAFU2int, lines...)))
}
