package path

type path[N comparable] struct {
	len  int
	h    int
	node N
	prev *path[N]
	idx  int
}

type pathHeap[N comparable] struct {
	data []*path[N]
}

func (h *pathHeap[N]) Len() int {
	return len(h.data)
}

func (h *pathHeap[N]) Less(i int, j int) bool {
	return h.data[i].len+h.data[i].h < h.data[j].len+h.data[j].h
}

func (h *pathHeap[N]) Swap(i int, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
	h.data[i].idx = i
	h.data[j].idx = j
}

func (h *pathHeap[N]) Push(x any) {
	p := x.(*path[N])
	p.idx = len(h.data)
	h.data = append(h.data, p)
}

func (h *pathHeap[N]) Pop() any {
	el := h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	el.idx = -1
	return el
}
