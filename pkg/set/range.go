package set

import "github.com/liennie/AdventOfCode/pkg/ints"

type Range struct {
	Min, Max int
}

func (r Range) Contains(other Range) bool {
	return r.Min <= other.Min && r.Max >= other.Max
}

func (r Range) Overlaps(other Range) bool {
	return r.Max >= other.Min && r.Min <= other.Max
}

func (r Range) Len() int {
	return r.Max - r.Min + 1
}

type RangeSet map[Range]struct{}

func (rs RangeSet) Add(r Range) {
	if r.Len() <= 0 {
		return
	}

	overlaps := []Range{}
	for o := range rs {
		if r.Overlaps(o) {
			overlaps = append(overlaps, o)
			delete(rs, o)
		}
	}

	rs[Range{
		Min: ints.MinFunc(func(o Range) int { return o.Min }, append(overlaps, r)...),
		Max: ints.MaxFunc(func(o Range) int { return o.Max }, append(overlaps, r)...),
	}] = struct{}{}
}

func (rs RangeSet) Remove(r Range) {
	if r.Len() <= 0 {
		return
	}

	overlaps := []Range{}
	for o := range rs {
		if r.Overlaps(o) {
			overlaps = append(overlaps, o)
			delete(rs, o)
		}
	}

	for _, o := range overlaps {
		if o.Min < r.Min {
			rs[Range{
				Min: o.Min,
				Max: r.Min - 1,
			}] = struct{}{}
		}
		if o.Max > r.Max {
			rs[Range{
				Min: r.Max + 1,
				Max: o.Max,
			}] = struct{}{}
		}
	}
}

func (rs RangeSet) Len() int {
	total := 0
	for o := range rs {
		total += o.Len()
	}
	return total
}
