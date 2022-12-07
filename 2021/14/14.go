package main

import (
	"math"
	"strings"

	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
)

func parse(filename string) (map[string]int, map[string]string) {
	ch := load.File(filename)

	template := <-ch
	<-ch // empty line

	pairs := map[string]int{}
	for i := 1; i < len(template); i++ {
		pairs[template[i-1:i+1]]++
	}
	pairs[template[len(template)-1:]+"."]++

	rules := map[string]string{}
	for line := range ch {
		r := strings.SplitN(line, " -> ", 2)
		if len(r[0]) != 2 || len(r[1]) != 1 {
			evil.Panic("Invalid rule %s", line)
		}
		rules[r[0]] = r[1]
	}

	return pairs, rules
}

func step(pairs map[string]int, rules map[string]string) map[string]int {
	newPairs := map[string]int{}

	for pair, count := range pairs {
		if insert, ok := rules[pair]; ok {
			newPairs[string(pair[0])+insert] += count
			newPairs[insert+string(pair[1])] += count
		} else {
			newPairs[pair] = count
		}
	}

	return newPairs
}

func score(pairs map[string]int) int {
	counts := map[byte]int{}
	for pair, count := range pairs {
		counts[pair[0]] += count
	}

	min := math.MaxInt
	max := 0
	for _, c := range counts {
		if c > max {
			max = c
		}
		if c < min {
			min = c
		}
	}

	return max - min
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	pairs, rules := parse(filename)

	// Part 1
	for i := 0; i < 10; i++ {
		pairs = step(pairs, rules)
	}
	log.Part1(score(pairs))

	// Part 2
	for i := 10; i < 40; i++ {
		pairs = step(pairs, rules)
	}
	log.Part2(score(pairs))
}
