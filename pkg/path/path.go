package path

import (
	"container/heap"
	"errors"
)

type Edge[N comparable] struct {
	Len int
	To  N
}

type Graph[N comparable] interface {
	Edges(N) []Edge[N]
}

type GraphFunc[N comparable] func(N) []Edge[N]

func (g GraphFunc[N]) Edges(n N) []Edge[N] {
	return g(n)
}

var _ Graph[int] = GraphFunc[int](nil)

type AStarGraph[N comparable] interface {
	Graph[N]
	Heuristic(N) int
}

type End[N comparable] interface {
	IsEnd(N) bool
}

type EndFunc[N comparable] func(N) bool

func (e EndFunc[N]) IsEnd(n N) bool {
	return e(n)
}

var _ End[int] = EndFunc[int](nil)

func EndConst[N comparable](e N) End[N] {
	return EndFunc[N](func(n N) bool { return n == e })
}

var ErrNotFound = errors.New("path not found")

func shortest[N comparable](g AStarGraph[N], start N, end End[N]) ([]N, int, error) {
	h := &pathHeap[N]{}
	sp := &path[N]{node: start}
	shortest := map[N]*path[N]{
		start: sp,
	}
	heap.Push(h, sp)

	for h.Len() > 0 {
		p := heap.Pop(h).(*path[N])

		if end.IsEnd(p.node) {
			path := []N{}

			for p := p; p != nil; p = p.prev {
				path = append(path, p.node)
			}
			for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
				path[i], path[j] = path[j], path[i]
			}

			return path, p.len, nil
		}

		for _, edge := range g.Edges(p.node) {
			l := p.len + edge.Len
			k := g.Heuristic(edge.To)

			if s, ok := shortest[edge.To]; ok {
				if l+k >= s.len+s.h {
					continue
				}

				s.len = l
				s.h = k
				s.prev = p
				heap.Fix(h, s.idx)
			} else {
				np := &path[N]{
					len:  l,
					h:    k,
					node: edge.To,
					prev: p,
				}
				shortest[edge.To] = np
				heap.Push(h, np)
			}
		}
	}

	return nil, 0, ErrNotFound
}

type dijkstraGraph[N comparable] struct {
	Graph[N]
}

func (dijkstraGraph[N]) Heuristic(N) int { return 0 }

func Shortest[N comparable](g Graph[N], start N, end End[N]) ([]N, int, error) {
	if ag, ok := g.(AStarGraph[N]); ok {
		return shortest(ag, start, end)
	}
	return shortest(dijkstraGraph[N]{g}, start, end)
}
