package test

import (
	"ember/algorithm/bytedance"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxSum(t *testing.T) {
	a := assert.New(t)
	a.Equal(bytedance.MaxSum([]int{1, 2, 4, 9}, 2533), 2499)
	a.Equal(bytedance.MaxSum([]int{4, 2, 9, 8}, 988822), 988499)
	a.Equal(bytedance.MaxSum([]int{9, 8}, 9), 8)
	a.Equal(bytedance.MaxSum([]int{9, 6, 3, 5}, 56449), 56399)
	a.Equal(bytedance.MaxSum([]int{9, 6, 3, 5, 1, 2, 4, 7, 8, 0}, 0x3f3f3f33), 1061109554)
}
