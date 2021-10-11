package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type PriorityQueue struct {
	list []*ListNode
	size int
}

func (r *PriorityQueue) Init() {
	r.list = make([]*ListNode, 0)
	r.size = 0
}

func (r *PriorityQueue) Empty() bool {
	return r.size == 0
}

func (r *PriorityQueue) Peek() *ListNode {
	if r.size == 0 {
		return nil
	}
	return r.list[0]
}

func (r *PriorityQueue) Offer(l *ListNode) {
	if r.size < len(r.list) {
		r.list[r.size] = l
	} else {
		r.list = append(r.list, l)
	}
	r.size++
	r.swim(r.size - 1)
}

func (r *PriorityQueue) Poll() *ListNode {
	if r.size == 0 {
		return nil
	}
	top := r.list[0]
	r.list[0] = r.list[r.size-1]
	r.size--
	r.sink(0)
	return top
}

func (r *PriorityQueue) swim(pos int) {
	for pos > 0 {
		k := pos / 2
		if pos%2 == 0 {
			k = pos/2 - 1
		}
		if r.list[k].Val <= r.list[pos].Val {
			break
		}
		r.list[k], r.list[pos] = r.list[pos], r.list[k]
		pos = k
	}
}

func (r *PriorityQueue) sink(pos int) {
	for 2*pos+1 < r.size {
		i := 2*pos + 1
		if i+1 < r.size && r.list[i].Val > r.list[i+1].Val {
			i++
		}
		if r.list[pos].Val <= r.list[i].Val {
			break
		}
		r.list[pos], r.list[i] = r.list[i], r.list[pos]
		pos = i
	}
}

func initData() *[]*ListNode {
	// lists = [[1,4,5],[1,3,4],[2,6]]
	//arr := [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}}
	arr := [][]int{{-4}, {-10, -6, -6}, {0, 3}, {2}, {-10, -9, -8, 3, 4, 4}, {-10, -10, -8, -6, -4, -3, 1}, {2}, {-9, -4, -2, 4, 4}, {-4, 0}}
	list := make([]*ListNode, len(arr))
	for i, v1 := range arr {
		dummy := new(ListNode)
		p := dummy
		for _, v2 := range v1 {
			l := new(ListNode)
			l.Val = v2
			p.Next = l
			p = p.Next
		}
		list[i] = dummy.Next
	}
	return &list
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	queue := new(PriorityQueue)
	queue.Init()
	for _, l := range lists {
		if l != nil {
			queue.Offer(l)
		}
	}

	dummy := new(ListNode)
	p := dummy
	for !queue.Empty() {
		top := queue.Poll()
		fmt.Printf("%d,", top.Val)
		p.Next = top
		p = p.Next
		if top.Next != nil {
			queue.Offer(top.Next)
		}
	}
	return dummy.Next
}

func main() {
	list := initData()
	sortedList := mergeKLists(*list)
	fmt.Println("")
	for sortedList != nil {
		fmt.Printf("%d,", sortedList.Val)
		sortedList = sortedList.Next
	}
}
