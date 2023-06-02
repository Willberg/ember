package test

import (
	"ember/structure/queue/deque"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMyDeque(t *testing.T) {
	a := assert.New(t)
	dq := deque.CreateDeque()
	for i := 1; i < 10; i += 2 {
		dq.PushTail([]int{i, 0})
	}
	for i := 10; i > 0; i -= 2 {
		dq.PushHead([]int{i, 1})
	}
	for i := 0; i < 5; i++ {
		v := dq.PopHead().([]int)
		a.Equal([]int{2 * (i + 1), 1}, v)
		// 不能使用fmt.Printf和fmt.Print, 否则出现no tests were run bug
		//fmt.Printf("(%d, %d)", v[0], v[1])
	}
	for i := 9; !dq.IsEmpty(); i -= 2 {
		v := dq.PopTail().([]int)
		a.Equal(v, []int{i, 0})
		//fmt.Printf("(%d, %d)", v[0], v[1])
	}
}

func TestIntDeque(t *testing.T) {
	a := assert.New(t)
	dq := deque.CreateDeque()
	for i := 1; i < 10; i += 2 {
		dq.PushTail(i)
	}
	for i := 10; i > 0; i -= 2 {
		dq.PushHead(i)
	}
	for i := 0; i < 5; i++ {
		v := dq.PopHead().(int)
		a.Equal(v, 2*(i+1))
		fmt.Println(v)
	}
	for i := 9; !dq.IsEmpty(); i -= 2 {
		v := dq.PopTail().(int)
		a.Equal(v, i)
		//fmt.Print(v, ",")
	}
}

func TestMo(t *testing.T) {
	fmt.Println(1)
}
