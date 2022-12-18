package space

import (
	"github.com/liennie/AdventOfCode/pkg/ints"
)

type AABB3 struct {
	Min Point3
	Max Point3
}

func NewAABB3(p Point3) AABB3 {
	return AABB3{
		Min: p,
		Max: p,
	}
}

func (aabb AABB3) Add(p Point3) AABB3 {
	return AABB3{
		Min: Point3{
			X: ints.Min(aabb.Min.X, p.X),
			Y: ints.Min(aabb.Min.Y, p.Y),
			Z: ints.Min(aabb.Min.Z, p.Z),
		},
		Max: Point3{
			X: ints.Max(aabb.Max.X, p.X),
			Y: ints.Max(aabb.Max.Y, p.Y),
			Z: ints.Max(aabb.Max.Z, p.Z),
		},
	}
}

func (aabb AABB3) Contains(p Point3) bool {
	return p.X >= aabb.Min.X && p.X <= aabb.Max.X &&
		p.Y >= aabb.Min.Y && p.Y <= aabb.Max.Y &&
		p.Z >= aabb.Min.Z && p.Z <= aabb.Max.Z
}
