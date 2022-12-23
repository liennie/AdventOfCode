package space

import (
	"github.com/liennie/AdventOfCode/pkg/ints"
)

type AABB struct {
	Valid bool
	Min   Point
	Max   Point
}

func NewAABB(points ...Point) AABB {
	aabb := AABB{}
	for _, p := range points {
		aabb = aabb.Add(p)
	}
	return aabb
}

func (aabb AABB) Add(p Point) AABB {
	if !aabb.Valid {
		return AABB{Valid: true, Min: p, Max: p}
	}
	return AABB{
		Valid: true,
		Min: Point{
			X: ints.Min(aabb.Min.X, p.X),
			Y: ints.Min(aabb.Min.Y, p.Y),
		},
		Max: Point{
			X: ints.Max(aabb.Max.X, p.X),
			Y: ints.Max(aabb.Max.Y, p.Y),
		},
	}
}

func (aabb AABB) Contains(p Point) bool {
	return aabb.Valid &&
		p.X >= aabb.Min.X && p.X <= aabb.Max.X &&
		p.Y >= aabb.Min.Y && p.Y <= aabb.Max.Y
}

func (aabb AABB) Expand(n int) AABB {
	if !aabb.Valid {
		return aabb
	}
	return AABB{
		Valid: true,
		Min:   aabb.Min.Add(Point{-n, -n}),
		Max:   aabb.Max.Add(Point{n, n}),
	}
}
