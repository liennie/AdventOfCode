package main

import (
	"cmp"
	"slices"
	"strings"

	"github.com/liennie/AdventOfCode/pkg/evil"
	"github.com/liennie/AdventOfCode/pkg/load"
	"github.com/liennie/AdventOfCode/pkg/log"
	"github.com/liennie/AdventOfCode/pkg/set"
	"github.com/liennie/AdventOfCode/pkg/space"
)

func parsePoint(raw string) space.Point3 {
	coords := evil.SplitN(raw, ",", 3)
	evil.Assert(len(coords) == 3)
	return space.Point3{
		X: coords[0],
		Y: coords[1],
		Z: coords[2],
	}
}

func parse(filename string) Tower {
	bricks := []*Brick{}
	for line := range load.File(filename) {
		min, max, ok := strings.Cut(line, "~")
		evil.Assert(ok)

		bricks = append(bricks, &Brick{aabb: space.NewAABB3(parsePoint(min), parsePoint(max))})
	}
	return NewTower(bricks)
}

type Brick struct {
	aabb space.AABB3
}

func (b *Brick) forEach(f func(space.Point3) bool) bool {
	if !b.aabb.Valid {
		return false
	}

	for x := b.aabb.Min.X; x <= b.aabb.Max.X; x++ {
		for y := b.aabb.Min.Y; y <= b.aabb.Max.Y; y++ {
			for z := b.aabb.Min.Z; z <= b.aabb.Max.Z; z++ {
				if !f(space.Point3{X: x, Y: y, Z: z}) {
					return false
				}
			}
		}
	}
	return true
}

type Tower struct {
	bricks []*Brick
	tower  map[space.Point3]*Brick
}

func NewTower(bricks []*Brick) Tower {
	t := map[space.Point3]*Brick{}
	for _, brick := range bricks {
		brick.forEach(func(p space.Point3) bool {
			t[p] = brick
			return true
		})
	}

	return Tower{
		bricks: bricks,
		tower:  t,
	}
}

func (t *Tower) move(brick *Brick, dir space.Point3) bool {
	if brick.aabb.Min.Add(dir).Z <= 0 {
		return false
	}

	if !brick.forEach(func(p space.Point3) bool {
		if other, ok := t.tower[p.Add(dir)]; ok && other != brick {
			return false
		}
		return true
	}) {
		return false
	}

	brick.forEach(func(p space.Point3) bool {
		delete(t.tower, p)
		return true
	})

	brick.aabb = space.NewAABB3(brick.aabb.Min.Add(dir), brick.aabb.Max.Add(dir))

	brick.forEach(func(p space.Point3) bool {
		t.tower[p] = brick
		return true
	})

	return true
}

func (t *Tower) fall() {
	falling := slices.Clone(t.bricks)
	slices.SortFunc(falling, func(a, b *Brick) int { return cmp.Compare(a.aabb.Min.Z, b.aabb.Min.Z) })

	gravity := space.Point3{Z: -1}
	for len(falling) > 0 {
		brick := falling[0]
		falling = falling[1:]

		if t.move(brick, gravity) {
			falling = append(falling, brick)
		}
	}
}

func main() {
	defer evil.Recover(log.Err)
	filename := load.Filename()

	tower := parse(filename)

	// Part 1
	tower.fall()

	supports := map[*Brick]set.Set[*Brick]{}
	supportedBy := map[*Brick]set.Set[*Brick]{}
	for _, brick := range tower.bricks {
		brick.forEach(func(p space.Point3) bool {
			if other, ok := tower.tower[p.Add(space.Point3{Z: 1})]; ok && other != brick {
				if supports[brick] == nil {
					supports[brick] = set.New[*Brick]()
				}
				supports[brick].Add(other)

				if supportedBy[other] == nil {
					supportedBy[other] = set.New[*Brick]()
				}
				supportedBy[other].Add(brick)
			}
			return true
		})
	}

	cnt := 0
bricks:
	for _, brick := range tower.bricks {
		for other := range supports[brick] {
			if len(supportedBy[other]) == 1 {
				continue bricks
			}
		}
		cnt++
	}
	log.Part1(cnt)
}
