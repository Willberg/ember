package test

import (
	"container/heap"
	. "ember/datastruct/heap"
	"fmt"
	"testing"
)

func TestHeap(t *testing.T) {
	h := &IntHeap{2, 1, 5, 100, 3, 6, 4, 5}
	heap.Init(h)
	heap.Push(h, 3)
	fmt.Printf("minimum: %d\n", (*h)[0])
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
}

func TestIntHeap(t *testing.T) {
	h := &IntHeap{}
	heap.Push(h, 2)
	heap.Push(h, 1)
	heap.Push(h, 3)
	fmt.Println(heap.Pop(h))
	fmt.Println(heap.Pop(h))
	fmt.Println(heap.Pop(h))
	fmt.Println(heap.Pop(h))
}
