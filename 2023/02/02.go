package main

import (
	"strings"

	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
)

type Cubes map[string]int
type Game []Cubes

func parse(filename string) map[int]Game {
	res := map[int]Game{}

	for line := range load.File(filename) {
		line = strings.TrimPrefix(line, "Game ")

		sid, line, ok := strings.Cut(line, ": ")
		evil.Assert(ok, "missing colon")
		id := evil.Atoi(sid)

		for _, turn := range strings.Split(line, ";") {
			c := Cubes{}

			for _, cubes := range strings.Split(turn, ",") {
				amt, color, ok := strings.Cut(strings.TrimSpace(cubes), " ")
				evil.Assert(ok, "invalid cubes")

				c[color] += evil.Atoi(amt)
			}

			res[id] = append(res[id], c)
		}
	}

	return res
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	games := parse(filename)

	// Part 1
	limit := Cubes{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	sum := 0
games:
	for id, game := range games {
		for _, turn := range game {
			for color, amt := range turn {
				if amt > limit[color] {
					continue games
				}
			}
		}

		sum += id
	}

	log.Part1(sum)
}
