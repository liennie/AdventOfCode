package path

import (
	"container/heap"
	"errors"
	"slices"
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
	ShortestRemainigDist(N) int
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

func getPaths[N comparable](p *path[N]) [][]N {
	path := []N{}
	for {
		path = append(path, p.node)
		switch len(p.prev) {
		case 0:
			slices.Reverse(path)
			return [][]N{path}

		case 1:
			p = p.prev[0]

		default:
			slices.Reverse(path)
			var paths [][]N
			for _, prev := range p.prev {
				for _, prevPath := range getPaths(prev) {
					paths = append(paths, append(prevPath, path...))
				}
			}
			return paths
		}
	}
}

func shortest[N comparable](g AStarGraph[N], start N, end End[N], all bool) ([][]N, int, error) {
	h := &pathHeap[N]{}
	sp := &path[N]{node: start}
	shortest := map[N]*path[N]{
		start: sp,
	}
	heap.Push(h, sp)

	for h.Len() > 0 {
		p := heap.Pop(h).(*path[N])

		if end.IsEnd(p.node) {
			return getPaths(p), p.len, nil
		}

		for _, edge := range g.Edges(p.node) {
			l := p.len + edge.Len
			k := g.ShortestRemainigDist(edge.To)

			if s, ok := shortest[edge.To]; ok {
				if l+k > s.len+s.h {
					continue
				} else if l+k == s.len+s.h {
					if all {
						s.prev = append(s.prev, p)
					} else {
						continue
					}
				} else {
					s.len = l
					s.h = k
					s.prev = []*path[N]{p}
					heap.Fix(h, s.idx)
				}
			} else {
				np := &path[N]{
					len:  l,
					h:    k,
					node: edge.To,
					prev: []*path[N]{p},
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

func (dijkstraGraph[N]) ShortestRemainigDist(N) int { return 0 }

func Shortest[N comparable](g Graph[N], start N, end End[N]) ([]N, int, error) {
	var ag AStarGraph[N]
	var ok bool
	if ag, ok = g.(AStarGraph[N]); !ok {
		ag = dijkstraGraph[N]{g}
	}
	paths, l, err := shortest(ag, start, end, false)
	if len(paths) > 0 {
		return paths[0], l, err
	}
	return nil, l, err
}

// AllShortest might not actually return all paths
// if there are multiple nodes that lead to the end node.
// It works if the end node is a dead end.
// It also doesn't work properly if the graph is an A*, so the Heuristic method
// is not used here.
func AllShortest[N comparable](g Graph[N], start N, end End[N]) ([][]N, int, error) {
	return shortest(dijkstraGraph[N]{g}, start, end, true)
}
