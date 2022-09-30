package PriorityQueue

type PriorityQueue struct {
	q              []interface{}
	size, capacity int
	Less           func(a, b interface{}) bool
}

func CreatePriorityQueue(n int, less func(a, b interface{}) bool) PriorityQueue {
	return PriorityQueue{make([]interface{}, n), 0, n, less}
}

func (pq *PriorityQueue) Push(v interface{}) bool {
	if pq.IsFull() {
		return false
	}
	pq.q[pq.size] = v
	pq.size++
	pq.up(pq.size - 1)
	return true
}

func (pq *PriorityQueue) Pop() interface{} {
	if pq.IsEmpty() {
		return nil
	}
	top := pq.q[0]
	pq.q[0] = pq.q[pq.size-1]
	pq.size--
	pq.down(0)
	return top
}

func (pq *PriorityQueue) Top() interface{} {
	return pq.q[0]
}

func (pq *PriorityQueue) IsFull() bool {
	return pq.size == pq.capacity
}

func (pq *PriorityQueue) IsEmpty() bool {
	return pq.size == 0
}

func (pq *PriorityQueue) up(pos int) {
	for pos > 0 {
		i := (pos - 1) / 2
		if pq.less(i, pos) {
			return
		}
		pq.q[i], pq.q[pos] = pq.q[pos], pq.q[i]
		pos = i
	}
}

func (pq *PriorityQueue) down(pos int) {
	for pos < pq.size {
		small, l := pos, 2*pos+1
		if l < pq.size && !pq.less(small, l) {
			small = l
		}
		r := l + 1
		if r < pq.size && !pq.less(small, r) {
			small = r
		}
		if small == pos {
			return
		}
		pq.q[pos], pq.q[small] = pq.q[small], pq.q[pos]
		pos = small
	}
}

func (pq *PriorityQueue) less(i, j int) bool {
	a, b := pq.q[i], pq.q[j]
	return pq.Less(a, b)
}
