package main

import (
	"bytes"
	"strings"

	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
)

type Springs struct {
	raw    []byte
	groups []int
}

func parse(filename string) []Springs {
	res := []Springs{}

	for line := range load.File(filename) {
		field, groups, ok := strings.Cut(line, " ")
		evil.Assert(ok, "invalid format")
		for _, c := range field {
			evil.Assert(c == '.' || c == '#' || c == '?', "invalid char ", c)
		}
		res = append(res, Springs{
			raw:    []byte(field),
			groups: evil.Split(groups, ","),
		})
	}

	return res
}

func generate(groups []int, l int) [][]byte {
	if l < 0 {
		return [][]byte{{}}
	}
	if len(groups) == 0 {
		return [][]byte{
			bytes.Repeat([]byte{'.'}, l),
		}
	}

	res := [][]byte{}

	to := l - groups[0]
	for _, g := range groups[1:] {
		to -= g
		to--
	}
	evil.Assert(to >= 0, "groups are too large")
	for i := 0; i <= to; i++ {
		nl := l - (i + groups[0] + 1)
		for _, g := range generate(groups[1:], nl) {
			gr := bytes.Repeat([]byte{'.'}, i)
			gr = append(gr, bytes.Repeat([]byte{'#'}, groups[0])...)
			if nl >= 0 {
				gr = append(gr, '.')
			}
			gr = append(gr, g...)
			res = append(res, gr)
		}
	}

	return res
}

func overlaps(a, b []byte) bool {
	evil.Assert(len(a) == len(b))

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
		for _, gr := range generate(spr.groups, len(spr.raw)) {
			if overlaps(gr, spr.raw) {
				sum++
			}
		}
	}
	log.Part1(sum)
}
