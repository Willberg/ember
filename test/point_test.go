package test

import (
	"fmt"
	"testing"
)

type Human struct {
	name string
}

func changeName(h Human, name string) {
	h.name = name
}

func changeNameS(h *Human, name string) {
	(*h).name = name
}

func changeValue(o int, n int) {
	o = n
}

func changeValueS(o *int, n int) {
	*o = n
}

func swap(a [3]int, i, j int) {
	a[i], a[j] = a[j], a[i]
}

func Swap(a *[3]int, i, j int) {
	(*a)[i], (*a)[j] = (*a)[j], (*a)[i]
}

func swapS(a []int, i, j int) {
	a[i], a[j] = a[j], a[i]
}

func SwapS(a *[]int, i, j int) {
	(*a)[i], (*a)[j] = (*a)[j], (*a)[i]
}

func change(m map[int]int, i, j int) {
	m[i], m[j] = m[j], m[i]
}

func Change(m *map[int]int, i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func send(ch chan int, v int) {
	ch <- v
}

func sendS(ch *chan int, v int) {
	*ch <- v
}

func TestPoint(t *testing.T) {
	// go 传递的参数都是值传递
	// 只是slice, map, chan 都是指针类型（使用了别名而已，传递的是指针的副本，在原来的数据上修改）, 数组类型传递的是数组的副本（不改变原来的数据）
	// 基本类型， struct类型等其他类型都不会修改原来的数据
	a := [3]int{1, 2, 3}
	swap(a, 0, 2)
	for _, v := range a {
		fmt.Printf("%d,", v)
	}
	fmt.Println("")
	Swap(&a, 0, 1)
	for _, v := range a {
		fmt.Printf("%d,", v)
	}
	fmt.Println("")

	sa := []int{1, 2, 3}
	swapS(sa, 0, 2)
	for _, v := range sa {
		fmt.Printf("%d,", v)
	}
	fmt.Println("")
	SwapS(&sa, 0, 1)
	for _, v := range sa {
		fmt.Printf("%d,", v)
	}
	fmt.Println("")

	m := make(map[int]int)
	m[1] = 1
	m[2] = 2
	m[3] = 3
	change(m, 1, 3)
	for k, v := range m {
		fmt.Printf("(%d,%d),", k, v)
	}
	fmt.Println("")
	Change(&m, 1, 2)
	for k, v := range m {
		fmt.Printf("(%d,%d),", k, v)
	}
	fmt.Println("")

	ch := make(chan int, 4)
	send(ch, 1)
	fmt.Println(<-ch)
	sendS(&ch, 2)
	fmt.Println(<-ch)

	h := Human{"old"}
	changeName(h, "new")
	fmt.Println(h.name)
	changeNameS(&h, "new")
	fmt.Println(h.name)

	o := 1
	changeValue(o, 2)
	fmt.Println(o)
	changeValueS(&o, 2)
	fmt.Println(o)

	arr := [][]int{{1, 1}, {2, 2}}
	// 此处的a只是副本
	for i, a := range arr {
		a = append(a, i)
		fmt.Printf("%d,", len(a))
	}
	fmt.Println("")
	for i := range arr {
		fmt.Printf("%d,", len(arr[i]))
		// 此处的arr[i]指对应的二维数组中相应的元素
		arr[i] = append(arr[i], i)
	}
	fmt.Println("")
	for _, a := range arr {
		fmt.Printf("%d,", len(a))
	}
}
