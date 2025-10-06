package space

import (
	"iter"

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

func (aabb AABB) Overlaps(other AABB) bool {
	return (aabb.Max.X >= other.Min.X && aabb.Min.X <= other.Max.X) &&
		(aabb.Max.Y >= other.Min.Y && aabb.Min.Y <= other.Max.Y)
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

func (aabb AABB) Size() Point {
	if !aabb.Valid {
		return Point{0, 0}
	}

	return aabb.Max.Sub(aabb.Min).Add(Point{1, 1})
}

func (aabb AABB) Clamp(p Point) Point {
	if !aabb.Valid {
		return p
	}
	return Point{
		X: ints.Clamp(p.X, aabb.Min.X, aabb.Max.X),
		Y: ints.Clamp(p.Y, aabb.Min.Y, aabb.Max.Y),
	}
}

func (aabb AABB) Wrap(p Point) Point {
	if !aabb.Valid {
		return p
	}
	return Point{
		X: ints.Wrap(p.X, aabb.Min.X, aabb.Max.X),
		Y: ints.Wrap(p.Y, aabb.Min.Y, aabb.Max.Y),
	}
}

func (aabb AABB) All() iter.Seq[Point] {
	return func(yield func(Point) bool) {
		if !aabb.Valid {
			return
		}

		for y := aabb.Min.Y; y <= aabb.Max.Y; y++ {
			for x := aabb.Min.X; x <= aabb.Max.X; x++ {
				if !yield(Point{X: x, Y: y}) {
					return
				}
			}
		}
	}
}
