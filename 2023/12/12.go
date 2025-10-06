package main

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
)

type Springs struct {
	raw    []byte
	groups []int
}

func (s Springs) unfold(n int) Springs {
	r := [][]byte{}
	g := []int{}
	for i := 0; i < n; i++ {
		r = append(r, s.raw)
		g = append(g, s.groups...)
	}

	return Springs{
		raw:    bytes.Join(r, []byte{'?'}),
		groups: g,
	}
}

func parse(filename string) []Springs {
	res := []Springs{}

	for line := range load.File(filename) {
		field, groups, ok := strings.Cut(line, " ")
		evil.Assert(ok, "invalid format")
		for _, c := range field {
			evil.Assert(c == '.' || c == '#' || c == '?', "invalid char %c", c)
		}
		res = append(res, Springs{
			raw:    []byte(field),
			groups: evil.Split(groups, ","),
		})
	}

	return res
}

var cache = map[string]int{}

func combinations(raw []byte, groups []int, l int) (res int) {
	cacheKey := fmt.Sprintf("%v %v %v", raw, groups, l)
	if r, ok := cache[cacheKey]; ok {
		return r
	}
	defer func() {
		cache[cacheKey] = res
	}()

	if len(groups) == 0 {
		for _, c := range raw {
			if c == '#' {
				return 0
			}
		}
		return 1
	}

	to := l - groups[0]
	for _, g := range groups[1:] {
		to -= g
		to--
	}
	evil.Assert(to >= 0, "groups are too large")

	res = 0

	for i := 0; i <= to; i++ {
		start := bytes.Repeat([]byte{'.'}, i)
		start = append(start, bytes.Repeat([]byte{'#'}, groups[0])...)
		if l > len(start) {
			start = append(start, '.')
		}

		if !overlaps(start, raw[:len(start)]) {
			continue
		}

		res += combinations(raw[len(start):], groups[1:], l-len(start))
	}

	return res
}

func overlaps(a, b []byte) bool {
	evil.Assert(len(a) == len(b), "different len: %d != %d", len(a), len(b))

	for i := range a {
		if a[i] != '?' && b[i] != '?' && a[i] != b[i] {
			return false
		}
	}
	return true
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	springs := parse(filename)

	// Part 1
	sum := 0
	for _, spr := range springs {
		sum += combinations(spr.raw, spr.groups, len(spr.raw))
	}
	log.Part1(sum)

	// Part 2
	unfolded := make([]Springs, len(springs))
	for i := range springs {
		unfolded[i] = springs[i].unfold(5)
	}
	sum = 0
	for _, spr := range unfolded {
		sum += combinations(spr.raw, spr.groups, len(spr.raw))
	}
	log.Part2(sum)
}
