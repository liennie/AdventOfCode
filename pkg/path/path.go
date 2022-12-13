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

func Shortest[N comparable](g Graph[N], start N, end End[N]) ([]N, int, error) {
	h := &pathHeap[N]{}
	shortest := map[N]int{
		start: 0,
	}
	heap.Push(h, &path[N]{node: start})

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

			if s, ok := shortest[edge.To]; ok {
				if l >= s {
					continue
				}
			}

			shortest[edge.To] = l
			heap.Push(h, &path[N]{
				len:  l,
				node: edge.To,
				prev: p,
			})
		}
	}

	return nil, 0, ErrNotFound
}
