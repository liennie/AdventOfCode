package main

import (
	"regexp"

	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/ints"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
)

type resources struct {
	ore      int
	clay     int
	obsidian int
	geode    int
}

func (r resources) add(other resources) resources {
	return resources{
		ore:      r.ore + other.ore,
		clay:     r.clay + other.clay,
		obsidian: r.obsidian + other.obsidian,
		geode:    r.geode + other.geode,
	}
}

func (r resources) sub(other resources) resources {
	return resources{
		ore:      r.ore - other.ore,
		clay:     r.clay - other.clay,
		obsidian: r.obsidian - other.obsidian,
		geode:    r.geode - other.geode,
	}
}

func (r resources) times(sc int) resources {
	return resources{
		ore:      r.ore * sc,
		clay:     r.clay * sc,
		obsidian: r.obsidian * sc,
		geode:    r.geode * sc,
	}
}

func (r resources) canAfford(cost resources) bool {
	return r.ore >= cost.ore &&
		r.clay >= cost.clay &&
		r.obsidian >= cost.obsidian &&
		r.geode >= cost.geode
}

type blueprint struct {
	id int

	ore      resources
	clay     resources
	obsidian resources
	geode    resources

	max resources
}

var blueprintRe = regexp.MustCompile(`Blueprint (\d+): Each ore robot costs (\d+) ore. Each clay robot costs (\d+) ore. Each obsidian robot costs (\d+) ore and (\d+) clay. Each geode robot costs (\d+) ore and (\d+) obsidian.`)

func parse(filename string) []blueprint {
	res := []blueprint{}
	for line := range load.File(filename) {
		match := blueprintRe.FindStringSubmatch(line)
		if match == nil {
			evil.Panic("Line %q does not match", line)
		}

		bl := blueprint{
			id:       evil.Atoi(match[1]),
			ore:      resources{ore: evil.Atoi(match[2])},
			clay:     resources{ore: evil.Atoi(match[3])},
			obsidian: resources{ore: evil.Atoi(match[4]), clay: evil.Atoi(match[5])},
			geode:    resources{ore: evil.Atoi(match[6]), obsidian: evil.Atoi(match[7])},
		}
		bl.max = resources{
			ore:      ints.Max(bl.ore.ore, bl.clay.ore, bl.obsidian.ore, bl.geode.ore),
			clay:     ints.Max(bl.ore.clay, bl.clay.clay, bl.obsidian.clay, bl.geode.clay),
			obsidian: ints.Max(bl.ore.obsidian, bl.clay.obsidian, bl.obsidian.obsidian, bl.geode.obsidian),
			geode:    ints.Max(bl.ore.geode, bl.clay.geode, bl.obsidian.geode, bl.geode.geode),
		}
		res = append(res, bl)
	}
	return res
}

func maximize(bl blueprint, time int, res resources, prod resources) int {
	buy := func(cost resources, np resources) int {
		if np.ore > bl.max.ore ||
			np.clay > bl.max.clay ||
			np.obsidian > bl.max.obsidian {
			return res.geode + time*prod.geode
		}

		resWait := func(res int, cost int, prod int) int {
			if res >= cost {
				return 0
			}
			if prod == 0 {
				return time + 1
			}
			return (cost-res-1)/prod + 1
		}

		wait := ints.Max(
			resWait(res.ore, cost.ore, prod.ore),
			resWait(res.clay, cost.clay, prod.clay),
			resWait(res.obsidian, cost.obsidian, prod.obsidian),
			resWait(res.geode, cost.geode, prod.geode),
		) + 1

		if wait > time {
			return res.geode + time*prod.geode
		}

		return maximize(bl, time-wait, res.add(prod.times(wait)).sub(cost), np)
	}

	return ints.Max(
		buy(bl.ore, prod.add(resources{ore: 1})),
		buy(bl.clay, prod.add(resources{clay: 1})),
		buy(bl.obsidian, prod.add(resources{obsidian: 1})),
		buy(bl.geode, prod.add(resources{geode: 1})),
	)
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	blueprints := parse(filename)

	// Part 1
	total := 0
	for _, bl := range blueprints {
		log.Println(bl)
		geodes := maximize(bl, 24, resources{}, resources{ore: 1})
		total += geodes * bl.id
	}
	log.Part1(total)
}
