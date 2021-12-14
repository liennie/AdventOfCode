package main

import (
	"math"
	"strings"

	"github.com/liennie/AdventOfCode/common/load"
	"github.com/liennie/AdventOfCode/common/log"
	"github.com/liennie/AdventOfCode/common/util"
)

func parse(filename string) (string, map[string]string) {
	ch := load.File(filename)

	template := <-ch
	<-ch // empty line

	rules := map[string]string{}
	for line := range ch {
		r := strings.SplitN(line, " -> ", 2)
		if len(r[0]) != 2 || len(r[1]) != 1 {
			util.Panic("Invalid rule %s", line)
		}
		rules[r[0]] = r[1]
	}

	return template, rules
}

func step(template string, rules map[string]string) string {
	b := &strings.Builder{}

	for i := 1; i < len(template); i++ {
		b.WriteByte(template[i-1])

		if insert, ok := rules[template[i-1:i+1]]; ok {
			b.WriteString(insert)
		}
	}
	b.WriteByte(template[len(template)-1])

	return b.String()
}

func score(polymer string) int {
	count := map[byte]int{}
	for i := 0; i < len(polymer); i++ {
		count[polymer[i]]++
	}

	min := math.MaxInt
	max := 0
	for _, c := range count {
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
	defer util.Recover(log.Err)

	const filename = "input.txt"

	template, rules := parse(filename)

	// Part 1
	polymer := template
	for i := 0; i < 10; i++ {
		polymer = step(polymer, rules)
	}
	log.Part1(score(polymer))
}
