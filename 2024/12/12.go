package main

import (
	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
	"github.com/liennie/AdventOfCode/pkg/set"
	"github.com/liennie/AdventOfCode/pkg/space"
)

func parse(filename string) map[space.Point]rune {
	res := map[space.Point]rune{}
	load.Grid(filename, func(p space.Point, r rune) {
		res[p] = r
	})
	return res
}

func area(region set.Set[space.Point]) int {
	return len(region)
}

func perimeter(region set.Set[space.Point]) int {
	total := 0

	for p := range region {
		for d := range space.Orthogonal() {
			if !region.Contains(p.Add(d)) {
				total++
			}
		}
	}

	return total
}

func sides(region set.Set[space.Point]) int {
	fences := map[space.Point3]set.RangeSet{}

	for p := range region {
		for d := range space.Orthogonal() {
			if region.Contains(p.Add(d)) {
				continue
			}

			horizontal := d.X != 0
			k := space.Point3{}
			if horizontal {
				k = space.Point3{
					X: d.X,
					Z: p.X,
				}
			} else {
				k = space.Point3{
					Y: d.Y,
					Z: p.Y,
				}
			}

			if fences[k] == nil {
				fences[k] = set.RangeSet{}
			}

			if horizontal {
				fences[k].Add(set.Range{p.Y, p.Y + 1})
			} else {
				fences[k].Add(set.Range{p.X, p.X + 1})
			}
		}
	}

	total := 0
	for _, rs := range fences {
		total += len(rs)
	}
	return total
}

func getRegion(seed space.Point, plots map[space.Point]rune, region set.Set[space.Point]) {
	region.Add(seed)

	plant := plots[seed]
	delete(plots, seed)

	for d := range space.Orthogonal() {
		if next := seed.Add(d); plots[next] == plant {
			getRegion(next, plots, region)
		}
	}
}

func getRegions(plots map[space.Point]rune) []set.Set[space.Point] {
	regions := []set.Set[space.Point]{}
	for len(plots) > 0 {
		var seed space.Point
		for seed = range plots {
			break
		}

		region := set.New[space.Point]()
		getRegion(seed, plots, region)
		regions = append(regions, region)
	}
	return regions
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	plots := parse(filename)

	// Part 1
	regions := getRegions(plots)
	total := 0
	for _, region := range regions {
		total += area(region) * perimeter(region)
	}
	log.Part1(total)

	// Part 2
	total = 0
	for _, region := range regions {
		total += area(region) * sides(region)
	}
	log.Part2(total)
}
