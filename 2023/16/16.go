package main

import (
	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
	"github.com/liennie/AdventOfCode/pkg/set"
	"github.com/liennie/AdventOfCode/pkg/space"
)

type Beam struct {
	pos space.Point
	dir space.Point
}

func parse(filename string) [][]byte {
	res := [][]byte{}
	for line := range load.File(filename) {
		res = append(res, []byte(line))
	}
	return res
}

func energize(grid [][]byte, start Beam) int {
	aabb := space.NewAABB(space.Point{}, space.Point{X: len(grid[0]) - 1, Y: len(grid) - 1})

	beams := set.New(start)
	energized := map[space.Point]set.Set[space.Point]{}
	for len(beams) > 0 {
		beam, _ := beams.Pop()
		if energized[beam.pos] == nil {
			energized[beam.pos] = set.New[space.Point]()
		}
		energized[beam.pos].Add(beam.dir)

		nextBeams := []Beam{}
		switch grid[beam.pos.Y][beam.pos.X] {
		case '.':
			nextBeams = append(nextBeams, Beam{
				pos: beam.pos.Add(beam.dir),
				dir: beam.dir,
			})
		case '\\':
			reflected := space.Point{X: beam.dir.Y, Y: beam.dir.X}
			nextBeams = append(nextBeams, Beam{
				pos: beam.pos.Add(reflected),
				dir: reflected,
			})
		case '/':
			reflected := space.Point{X: -beam.dir.Y, Y: -beam.dir.X}
			nextBeams = append(nextBeams, Beam{
				pos: beam.pos.Add(reflected),
				dir: reflected,
			})
		case '|':
			if beam.dir.X == 0 {
				nextBeams = append(nextBeams, Beam{
					pos: beam.pos.Add(beam.dir),
					dir: beam.dir,
				})
			} else {
				nextBeams = append(nextBeams,
					Beam{
						pos: beam.pos.Add(space.Point{Y: 1}),
						dir: space.Point{Y: 1},
					},
					Beam{
						pos: beam.pos.Add(space.Point{Y: -1}),
						dir: space.Point{Y: -1},
					},
				)
			}
		case '-':
			if beam.dir.Y == 0 {
				nextBeams = append(nextBeams, Beam{
					pos: beam.pos.Add(beam.dir),
					dir: beam.dir,
				})
			} else {
				nextBeams = append(nextBeams,
					Beam{
						pos: beam.pos.Add(space.Point{X: 1}),
						dir: space.Point{X: 1},
					},
					Beam{
						pos: beam.pos.Add(space.Point{X: -1}),
						dir: space.Point{X: -1},
					},
				)
			}
		}

		for _, nextBeam := range nextBeams {
			if aabb.Contains(nextBeam.pos) && !energized[nextBeam.pos].Contains(nextBeam.dir) {
				beams.Add(Beam{
					pos: nextBeam.pos,
					dir: nextBeam.dir,
				})
			}
		}
	}
	return len(energized)
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	grid := parse(filename)

	// Part 1
	log.Part1(energize(grid, Beam{
		pos: space.Point{X: 0, Y: 0},
		dir: space.Point{X: 1},
	}))

	// Part 2
	max := 0
	for y := range grid {
		if e := energize(grid, Beam{
			pos: space.Point{X: 0, Y: y},
			dir: space.Point{X: 1},
		}); e > max {
			max = e
		}
		if e := energize(grid, Beam{
			pos: space.Point{X: len(grid[y]) - 1, Y: y},
			dir: space.Point{X: -1},
		}); e > max {
			max = e
		}
	}
	for x := range grid[0] {
		if e := energize(grid, Beam{
			pos: space.Point{X: x, Y: 0},
			dir: space.Point{Y: 1},
		}); e > max {
			max = e
		}
		if e := energize(grid, Beam{
			pos: space.Point{X: x, Y: len(grid) - 1},
			dir: space.Point{Y: -1},
		}); e > max {
			max = e
		}
	}
	log.Part2(max)
}
