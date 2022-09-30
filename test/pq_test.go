package test

import (
	"container/heap"
	. "ember/datastruct/priorityqueue"
	"fmt"
	"github.com/emirpasic/gods/queues/priorityqueue"
	"github.com/emirpasic/gods/utils"
	"testing"
)

type stu struct {
	name string
	age  int
}

type Stu []stu

func (s Stu) Len() int {
	return len(s)
}

func (s Stu) Less(i, j int) bool {
	return s[i].age < s[j].age
}

func (s Stu) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s *Stu) Push(x interface{}) {
	*s = append(*s, x.(stu))
}

func (s *Stu) Pop() interface{} {
	old := *s
	n := len(old)
	x := old[n-1]
	*s = old[:n-1]
	return x
}

func TestPq(t *testing.T) {
	student := &Stu{{"Amy", 21}, {"Dav", 15}, {"Spo", 22}, {"Reb", 11}}
	heap.Init(student)
	one := stu{"hund", 9}
	heap.Push(student, one)
	for student.Len() > 0 {
		fmt.Printf("%v\n", heap.Pop(student))
	}
}

func TestPriorityQueue(t *testing.T) {
	pq := priorityqueue.NewWith(func(a, b interface{}) int {
		// "-" descending order
		return utils.IntComparator(a.(int), b.(int))
	})
	pq.Enqueue(2)
	pq.Enqueue(1)
	for !pq.Empty() {
		v, _ := pq.Dequeue()
		fmt.Println(v)
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	n := len(lists)
	if n == 0 {
		return nil
	}
	pq := CreatePriorityQueue(n, func(a, b interface{}) bool {
		return a.(*ListNode).Val < b.(*ListNode).Val
	})

	for _, l := range lists {
		if l != nil {
			pq.Push(l)
		}
	}
	dummy := &ListNode{}
	p := dummy
	for !pq.IsEmpty() {
		t := pq.Pop().(*ListNode)
		if t.Next != nil {
			pq.Push(t.Next)
		}
		p.Next = t
		p = p.Next
	}
	return dummy.Next
}

func mergeLists(lists [][]int) []int {
	n := len(lists)
	if n == 0 {
		return nil
	}
	pq := CreatePriorityQueue(n, func(a, b interface{}) bool {
		return a.([]int)[0] < b.([]int)[0]
	})
	for _, l := range lists {
		if l != nil {
			pq.Push(l)
		}
	}
	var ans []int
	for !pq.IsEmpty() {
		t := pq.Pop().([]int)
		ans = append(ans, t[0])
		t = t[1:]
		if len(t) > 0 {
			pq.Push(t)
		}
	}
	return ans
}

func createListNodes(a []int) *ListNode {
	dummy := ListNode{}
	p := &dummy
	for _, v := range a {
		t := ListNode{Val: v}
		p.Next = &t
		p = p.Next
	}
	return dummy.Next
}

// [[1,4,5],[1,3,4],[2,6]]
func TestListNodePq(t *testing.T) {
	var lists []*ListNode
	arr := [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}}
	a := mergeLists(arr)
	for _, v := range a {
		fmt.Printf("%d,", v)
	}
	fmt.Println("")

	for _, a := range arr {
		lists = append(lists, createListNodes(a))
	}
	for _, p := range lists {
		for p != nil {
			fmt.Printf("%d,", p.Val)
			p = p.Next
		}
		fmt.Println("")
	}

	p := mergeKLists(lists)
	for p != nil {
		fmt.Printf("%d,", p.Val)
		p = p.Next
	}
}
