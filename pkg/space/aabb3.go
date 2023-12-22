package space

import (
	"github.com/liennie/AdventOfCode/pkg/ints"
)

type AABB3 struct {
	Valid bool
	Min   Point3
	Max   Point3
}

func NewAABB3(points ...Point3) AABB3 {
	aabb := AABB3{}
	for _, p := range points {
		aabb = aabb.Add(p)
	}
	return aabb
}

func (aabb AABB3) Add(p Point3) AABB3 {
	if !aabb.Valid {
		return AABB3{Valid: true, Min: p, Max: p}
	}
	return AABB3{
		Valid: true,
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
	return aabb.Valid &&
		p.X >= aabb.Min.X && p.X <= aabb.Max.X &&
		p.Y >= aabb.Min.Y && p.Y <= aabb.Max.Y &&
		p.Z >= aabb.Min.Z && p.Z <= aabb.Max.Z
}

func (aabb AABB3) Overlaps(other AABB3) bool {
	return (aabb.Max.X >= other.Min.X && aabb.Min.X <= other.Max.X) &&
		(aabb.Max.Y >= other.Min.Y && aabb.Min.Y <= other.Max.Y) &&
		(aabb.Max.Z >= other.Min.Z && aabb.Min.Z <= other.Max.Z)
}

func (aabb AABB3) Expand(n int) AABB3 {
	if !aabb.Valid {
		return aabb
	}
	return AABB3{
		Valid: true,
		Min:   aabb.Min.Add(Point3{-n, -n, -n}),
		Max:   aabb.Max.Add(Point3{n, n, n}),
	}
}
