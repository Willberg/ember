package test

import (
	"ember/datastruct/tree/segtree"
	"math"
	"testing"
)

func TestSegTreeBase(t *testing.T) {
	s := segtree.CreateSegment(math.MinInt32, math.MaxInt32)
	s.ToString()

	s.Add(10, 30, 1)
	s.ToString()

	s.Add(20, 40, 1)
	s.ToString()

	s.Add(10, 40, -2)
	s.ToString()

	s = segtree.CreateSegment(math.MinInt32, math.MaxInt32)
	s.ToString()

	s.Add(10, 30, 1)
	s.ToString()

	s.Add(20, 40, 1)
	s.ToString()

	s.Add(10, 40, -1)
	s.ToString()

	s.Add(10, 40, -1)
	s.ToString()
}

func TestSegTreeBase1(t *testing.T) {
	s := segtree.CreateSegment(math.MinInt32, math.MaxInt32)
	s.ToString()

	s.Add(10, 30, 1)
	s.ToString()

	s.Add(20, 40, 1)
	s.ToString()

	s.Add(10, 40, -2)
	s.ToString()

	s.Set(20, 3)
	s.ToString()

	s.Set(1, 1)
	s.ToString()
}

func TestSegTreeBase2(t *testing.T) {
	s := segtree.CreateSegment(math.MinInt32, math.MaxInt32)
	s.ToString()

	s.Set(1, 1)
	s.Set(2, 2)
	s.ToString()

	s.Set(3, 2)
	s.ToString()

	s.Set(1, 2)
	s.ToString()

	s.Add(2, 4, 1)
	s.ToString()
}
