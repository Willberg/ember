package test

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestSortSearch(t *testing.T) {
	a := assert.New(t)
	pos := sort.IntSlice{1, 2, 4, 5, 6}
	k := pos.Search(3)
	a.Equal(k, 2)
}

func BenchmarkSortSearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pos := sort.IntSlice{1, 2, 4, 5, 6}
		pos.Search(3)
	}
}
