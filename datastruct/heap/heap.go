package heap

type IntHeap []int

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(v interface{}) { *h = append(*h, v.(int)) }
func (h *IntHeap) Pop() interface{} {
	n := len(*h)
	v := (*h)[n-1]
	*h = (*h)[:n-1]
	return v
}
