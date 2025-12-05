package main

import (
	"strconv"
	"strings"

	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
	"github.com/liennie/AdventOfCode/pkg/set"
)

func parse(filename string) []set.Range {
	res := []set.Range{}
	for line := range load.File(filename) {
		for _, r := range strings.Split(line, ",") {
			if r == "" {
				continue
			}
			min, max, ok := strings.Cut(r, "-")
			evil.Assert(ok, "wrong format %q", r)

			res = append(res, set.Range{
				Min: evil.Atoi(min),
				Max: evil.Atoi(max),
			})
		}
	}
	return res
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	ranges := parse(filename)

	// Part 1
	sum := 0
	for _, r := range ranges {
		for n := r.Min; n <= r.Max; n++ {
			s := strconv.Itoa(n)
			if len(s)%2 != 0 {
				continue
			}

			if s[:len(s)/2] == s[len(s)/2:] {
				sum += n
			}
		}
	}
	log.Part1(sum)

	// Part 2
	sum = 0
	for _, r := range ranges {
	n:
		for n := r.Min; n <= r.Max; n++ {
			s := strconv.Itoa(n)

		d:
			for d := 1; d <= len(s)/2; d++ {
				if len(s)%d != 0 {
					continue
				}

				rep := s[:d]
				for i := d; i < len(s); i += d {
					if s[i:i+d] != rep {
						continue d
					}
				}

				sum += n
				continue n
			}
		}
	}
	log.Part2(sum)
}
