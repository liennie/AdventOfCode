package space

import (
	"github.com/liennie/AdventOfCode/pkg/ints"
)

type AABB struct {
	Min Point
	Max Point
}

func NewAABB(p Point) AABB {
	return AABB{
		Min: p,
		Max: p,
	}
}

func (aabb AABB) Add(p Point3) AABB {
	return AABB{
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
	return p.X >= aabb.Min.X && p.X <= aabb.Max.X &&
		p.Y >= aabb.Min.Y && p.Y <= aabb.Max.Y
}

func (aabb AABB) Expand(n int) AABB {
	return AABB{
		Min: aabb.Min.Add(Point{-n, -n}),
		Max: aabb.Max.Add(Point{n, n}),
	}
}
