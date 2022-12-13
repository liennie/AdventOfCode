package main

import (
	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
	"golang.org/x/exp/slices"
)

type value interface {
	compare(other value) int
}

type integer int

func (i integer) compare(other value) int {
	switch other := other.(type) {
	case integer:
		if i < other {
			return -1
		} else if i > other {
			return 1
		}
		return 0

	case list:
		return list{i}.compare(other)
	}
	return 0
}

type list []value

func (l list) compare(other value) int {
	switch other := other.(type) {
	case list:
		for i := 0; i < len(l) && i < len(other); i++ {
			if c := l[i].compare(other[i]); c != 0 {
				return c
			}
		}
		return integer(len(l)).compare(integer(len(other)))

	case integer:
		return l.compare(list{other})
	}
	return 0
}

func parseValue(v string) (el value, rest string) {
	defer func() {
		if rest != "" && rest[0] == ',' {
			rest = rest[1:]
		}
	}()

	switch v[0] {
	case '[':
		res := list{}
		v = v[1:]
		for v[0] != ']' {
			el, v = parseValue(v)
			res = append(res, el)
		}
		return res, v[1:]

	default:
		for i, c := range v {
			if c < '0' || c > '9' {
				return integer(evil.Atoi(v[:i])), v[i:]
			}
		}
		return integer(evil.Atoi(v)), ""
	}
}

func parse(filename string) [][2]value {
	res := [][2]value{}
	i := 0
	for line := range load.File(filename) {
		if line == "" {
			i = 0
			continue
		}

		if i == 0 {
			res = append(res, [2]value{})
		}
		res[len(res)-1][i], _ = parseValue(line)
		i++
	}
	return res
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	pairs := parse(filename)

	// Part 1
	sum := 0
	for i, pair := range pairs {
		if pair[0].compare(pair[1]) < 0 {
			sum += i + 1
		}
	}
	log.Part1(sum)

	// Part 2
	all := []value{
		list{list{integer(2)}},
		list{list{integer(6)}},
	}
	for _, pair := range pairs {
		all = append(all, pair[:]...)
	}
	slices.SortFunc(all, func(a, b value) bool { return a.compare(b) < 0 })
	div1, ok1 := slices.BinarySearchFunc[value](all, integer(2), func(a, b value) int { return a.compare(b) })
	div2, ok2 := slices.BinarySearchFunc[value](all, integer(6), func(a, b value) int { return a.compare(b) })
	if !ok1 || !ok2 {
		evil.Panic("Dividers not found")
	}
	log.Part2((div1 + 1) * (div2 + 1))
}
