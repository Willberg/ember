// 带懒标记的动态开点线段树
package segtree

import (
	"fmt"
	"github.com/emirpasic/gods/trees/redblacktree"
)

type Node struct {
	sum         int
	lazy        int
	left, right *Node
}

type Segment struct {
	segTree *Node
	mp      *redblacktree.Tree
	l, r    int
}

func CreateSegment(l, r int) Segment {
	return Segment{
		&Node{},
		redblacktree.NewWithIntComparator(),
		l,
		r,
	}
}

func (s *Segment) query(node *Node, l, r, from_, to_ int) int {
	if from_ <= l && r <= to_ {
		return node.sum
	}
	mid := l + (r-l)/2
	s.pushdown(node, mid-l+1, r-mid)
	res := 0
	if from_ <= mid {
		res = s.query(node.left, l, mid, from_, to_)
	}
	if to_ > mid {
		res += s.query(node.right, mid+1, r, from_, to_)
	}
	return res
}

func (s *Segment) update(node *Node, l, r, from_, to_, val int) {
	if from_ <= l && r <= to_ {
		node.sum += val * (r - l + 1)
		// 网上这行都是 node.lazy = val, 都错了
		node.lazy += val
		return
	}
	mid := l + (r-l)/2
	s.pushdown(node, mid-l+1, r-mid)
	if from_ <= mid {
		s.update(node.left, l, mid, from_, to_, val)
	}
	if to_ > mid {
		s.update(node.right, mid+1, r, from_, to_, val)
	}
	node.sum = node.left.sum + node.right.sum
}

func (s *Segment) pushdown(node *Node, leftLen, rightLen int) {
	if node.left == nil {
		node.left = &Node{}
	}
	if node.right == nil {
		node.right = &Node{}
	}
	if node.lazy == 0 {
		return
	}
	node.left.sum += node.lazy * leftLen
	node.right.sum += node.lazy * rightLen
	node.left.lazy = node.lazy
	node.right.lazy = node.lazy
	node.lazy = 0
}

func (s *Segment) Add(from_, to_, val int) {
	s.update(s.segTree, s.l, s.r, from_, to_-1, val)
	s.mp.Put(from_, nil)
	s.mp.Put(to_, nil)
}

func (s *Segment) Set(key, val int) {
	v := s.query(s.segTree, s.l, s.r, key, key)
	if v == val {
		return
	}
	s.update(s.segTree, s.l, s.r, key, key, val-v)
	s.mp.Put(key, nil)
	s.mp.Put(key+1, nil)
}

func (s *Segment) ToString() {
	keys := s.mp.Keys()
	fmt.Print("[")
	if len(keys) > 0 {
		pre := s.query(s.segTree, s.l, s.r, keys[0].(int), keys[0].(int))
		if pre != 0 {
			fmt.Printf("[%d, %d]", keys[0].(int), pre)
		}
		for _, key := range keys[1:] {
			val := s.query(s.segTree, s.l, s.r, key.(int), key.(int))
			if val != pre {
				fmt.Printf(",[%d, %d]", key.(int), val)
				pre = val
			}
		}
	}
	fmt.Println("]")
}
