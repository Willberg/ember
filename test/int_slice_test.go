package test

import (
	"fmt"
	"sort"
	"testing"
)

func TestIntSlice(t *testing.T) {
	pos := sort.IntSlice{1, 2, 4, 5, 6}
	k := pos.Search(3)
	fmt.Println(k)
}
