package main

import (
	"ember/structure/queue/deque"
	"fmt"
	"index/suffixarray"
	"unsafe"
)

// 通过unsafe访问别的包的私有变量
func main() {
	saTest()
	dequeTest()
}

func saTest() {
	s := "hello, world!"
	sa := (*struct {
		_  []byte
		sa []int32
	})(unsafe.Pointer(suffixarray.New([]byte(s)))).sa
	for _, i := range sa {
		fmt.Println(i, s[i:])
	}
}

func dequeTest() {
	d := deque.CreateDeque()
	d.PushHead(1)
	d.PushHead(2)
	p := (*struct {
		_    *deque.Node
		_    *deque.Node
		size int
	})(unsafe.Pointer(&d))
	fmt.Println(p.size)
	d.PopTail()
	fmt.Println(p.size)
	fmt.Println(d.IsEmpty())
	p.size = 0
	fmt.Println(d.IsEmpty(), p.size)
}
