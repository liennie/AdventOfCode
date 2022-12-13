package path

type path[N comparable] struct {
	len  int
	node N
	prev *path[N]
}

type pathHeap[N comparable] struct {
	data []*path[N]
}

func (h *pathHeap[N]) Len() int {
	return len(h.data)
}

func (h *pathHeap[N]) Less(i int, j int) bool {
	return h.data[i].len < h.data[j].len
}

func (h *pathHeap[N]) Swap(i int, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *pathHeap[N]) Push(x any) {
	h.data = append(h.data, x.(*path[N]))
}

func (h *pathHeap[N]) Pop() any {
	el := h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	return el
}
