package bytedance

import (
	"slices"
	"strconv"
)

// 字节面试题
// 数组A中给定可以使用的1~9的数，返回由A数组中的元素组成的小于n（n > 0）的最大数。 例如：A = {1, 2, 4, 9}，n = 2533，返回2499。
func MaxSum(a []int, n int) (res int) {
	s := strconv.Itoa(n)
	m := len(s)
	slices.Sort(a)
	var dfs func(i int, isLimit int, val int)
	dfs = func(i, isLimit, val int) {
		if i == m {
			if val < n {
				res = max(res, val)
			}
			return
		}
		if isLimit == 0 {
			dfs(i+1, 0, val*10+a[len(a)-1])
			return
		}
		up := int(s[i] - '0')
		for _, x := range a {
			if x > up {
				break
			}
			if x == up {
				dfs(i+1, 1, val*10+x)
			} else {
				dfs(i+1, 0, val*10+x)
			}
		}
	}
	dfs(0, 1, 0)
	return
}
