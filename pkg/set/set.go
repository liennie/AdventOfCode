package set

import (
	"cmp"
	"fmt"
	"iter"
	"slices"

	"golang.org/x/exp/maps"
)

type Set[T comparable] map[T]struct{}

type String = Set[string]

func New[T comparable](items ...T) Set[T] {
	set := Set[T]{}
	set.Add(items...)
	return set
}

func Collect[T comparable](items iter.Seq[T]) Set[T] {
	set := Set[T]{}
	for item := range items {
		set.Add(item)
	}
	return set
}

func (s Set[T]) Add(items ...T) {
	for _, item := range items {
		s[item] = struct{}{}
	}
}

func (s Set[T]) AddSeq(items iter.Seq[T]) {
	for item := range items {
		s[item] = struct{}{}
	}
}

func (s Set[T]) Remove(items ...T) {
	for _, item := range items {
		delete(s, item)
	}
}

func (s Set[T]) RemoveSeq(items iter.Seq[T]) {
	for item := range items {
		delete(s, item)
	}
}

func (s Set[T]) Contains(items ...T) bool {
	for _, item := range items {
		if _, ok := s[item]; !ok {
			return false
		}
	}
	return true
}

func (s Set[T]) ContainsSeq(items iter.Seq[T]) bool {
	for item := range items {
		if _, ok := s[item]; !ok {
			return false
		}
	}
	return true
}

func (s Set[T]) Intersects(items ...T) bool {
	for _, item := range items {
		if _, ok := s[item]; ok {
			return true
		}
	}
	return false
}

func (s Set[T]) IntersectsSeq(items iter.Seq[T]) bool {
	for item := range items {
		if _, ok := s[item]; ok {
			return true
		}
	}
	return false
}

func (s Set[T]) Clone() Set[T] {
	res := make(Set[T], len(s))
	for item := range s {
		res.Add(item)
	}
	return res
}

func (s Set[T]) Pop() (T, bool) {
	for item := range s {
		delete(s, item)
		return item, true
	}
	var zero T
	return zero, false
}

func (s Set[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for item := range s {
			if !yield(item) {
				return
			}
		}
	}
}

func (s Set[T]) String() string {
	return fmt.Sprint(maps.Keys(s))
}

func Intersection[T comparable](sets ...Set[T]) Set[T] {
	if len(sets) == 0 {
		return nil
	}

	slices.SortFunc(sets, func(a, b Set[T]) int { return cmp.Compare(len(a), len(b)) })

	res := make(Set[T], len(sets[0]))
items:
	for item := range sets[0] {
		for _, set := range sets[1:] {
			if !set.Contains(item) {
				continue items
			}
		}
		res.Add(item)
	}
	return res
}

func Union[T comparable](sets ...Set[T]) Set[T] {
	if len(sets) == 0 {
		return nil
	}

	cap := 0
	for _, set := range sets {
		cap += len(set)
	}

	res := make(Set[T], cap)
	for _, set := range sets {
		for item := range set {
			res.Add(item)
		}
	}
	return res
}
