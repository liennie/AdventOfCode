package main

import (
	"container/list"
	"sort"
	"strings"

	"github.com/liennie/AdventOfCode/common/load"
	"github.com/liennie/AdventOfCode/common/log"
	"github.com/liennie/AdventOfCode/common/util"
)

type pair struct {
	from, to util.Point3
	vec      util.Point3
}

func (p pair) rot90(r util.Point3) pair {
	return pair{
		from: p.from.Rot90(r),
		to:   p.to.Rot90(r),
		vec:  p.vec.Rot90(r),
	}
}

type scanner struct {
	beacons []util.Point3
	pairs   []pair
	pos     util.Point3
}

func (s scanner) rot90(r util.Point3) scanner {
	beacons := make([]util.Point3, len(s.beacons))
	for i := range s.beacons {
		beacons[i] = s.beacons[i].Rot90(r)
	}

	pairs := make([]pair, len(s.pairs))
	for i := range s.pairs {
		pairs[i] = s.pairs[i].rot90(r)
	}

	return scanner{
		beacons: beacons,
		pairs:   pairs,
	}
}

func createPairs(beacons []util.Point3) []pair {
	res := []pair{}

	for i := 0; i < len(beacons)-1; i++ {
		for j := i + 1; j < len(beacons); j++ {
			res = append(res, pair{
				from: beacons[i],
				to:   beacons[j],
				vec:  beacons[j].Sub(beacons[i]),
			})
		}
	}

	return res
}

func parse(filename string) []scanner {
	res := []scanner{}

	ch := load.File(filename)

	for {
		s, ok := <-ch
		if !ok {
			break
		}

		if !strings.HasPrefix(s, "--- scanner ") {
			util.Panic("Invalid format")
		}

		scanner := scanner{}

		for line := range ch {
			if line == "" {
				break
			}

			coords := util.SplitN(line, ",", 3)
			scanner.beacons = append(scanner.beacons, util.Point3{
				X: coords[0],
				Y: coords[1],
				Z: coords[2],
			})
		}

		sort.Slice(scanner.beacons, func(i, j int) bool {
			return scanner.beacons[i].ManhattanLen() < scanner.beacons[j].ManhattanLen()
		})

		scanner.pairs = createPairs(scanner.beacons)

		res = append(res, scanner)
	}

	return res
}

var rots = []util.Point3{
	{X: 0, Y: 0, Z: 0},
	{X: 1, Y: 0, Z: 0},
	{X: 2, Y: 0, Z: 0},
	{X: 3, Y: 0, Z: 0},
	{X: 0, Y: 0, Z: 1},
	{X: 1, Y: 0, Z: 1},
	{X: 2, Y: 0, Z: 1},
	{X: 3, Y: 0, Z: 1},
	{X: 0, Y: 0, Z: 2},
	{X: 1, Y: 0, Z: 2},
	{X: 2, Y: 0, Z: 2},
	{X: 3, Y: 0, Z: 2},
	{X: 0, Y: 0, Z: 3},
	{X: 1, Y: 0, Z: 3},
	{X: 2, Y: 0, Z: 3},
	{X: 3, Y: 0, Z: 3},
	{X: 0, Y: 1, Z: 0},
	{X: 1, Y: 1, Z: 0},
	{X: 2, Y: 1, Z: 0},
	{X: 3, Y: 1, Z: 0},
	{X: 0, Y: 3, Z: 0},
	{X: 1, Y: 3, Z: 0},
	{X: 2, Y: 3, Z: 0},
	{X: 3, Y: 3, Z: 0},
}

func main() {
	defer util.Recover(log.Err)

	const filename = "input.txt"

	scanners := parse(filename)

	// Part 1
	disambiguated := []scanner{scanners[0]}

	ambiguous := list.New()
	for i := 1; i < len(scanners); i++ {
		ambiguous.PushBack(scanners[i])
	}

	for i := 0; i < len(disambiguated); i++ {
		dis := disambiguated[i]

		var next *list.Element
		for amb := ambiguous.Front(); amb != nil; amb = next {
			next = amb.Next()

			for _, rot := range rots {
				s := amb.Value.(scanner).rot90(rot)

				candidates := map[util.Point3]int{}
				for _, disp := range dis.pairs {
					for _, ambp := range s.pairs {
						if disp.vec.Equals(ambp.vec) {
							can := dis.pos.Add(disp.from).Sub(ambp.from)
							candidates[can]++
						} else if disp.vec.Equals(ambp.vec.Flip()) {
							can := dis.pos.Add(disp.from).Sub(ambp.to)
							candidates[can]++
						}
					}
				}

				var pos util.Point3
				var found int
				for p, c := range candidates {
					if c >= (11*12)/2 {
						pos = p
						found++
					}
				}

				if found == 1 {
					s.pos = pos
					disambiguated = append(disambiguated, s)
					ambiguous.Remove(amb)
				} else if found > 1 {
					log.Println("Oh no")
				}
			}
		}
	}

	if ambiguous.Len() > 0 {
		util.Panic("Oh no, %d ambiguous", ambiguous.Len())
	}

	beacons := map[util.Point3]bool{}
	for _, scanner := range disambiguated {
		for _, beacon := range scanner.beacons {
			beacons[scanner.pos.Add(beacon)] = true
		}
	}

	log.Part1(len(beacons))

	// Part 2
	max := 0
	for _, a := range disambiguated {
		for _, b := range disambiguated {
			if l := a.pos.Sub(b.pos).ManhattanLen(); l > max {
				max = l
			}
		}
	}
	log.Part2(max)
}
