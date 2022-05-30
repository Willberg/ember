package test

import (
	"fmt"
	"sort"
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
	// sa为指向相应位置的指针
	for i, sa := range arr {
		// append之后数组扩容，指向了别处，因此原来数组不变
		sa = append(sa, i)
		fmt.Printf("%d,", len(sa))
	}
	fmt.Println("")
	for i := range arr {
		fmt.Printf("%d,", len(arr[i]))
		// 此处的arr[i]指对应的二维数组中相应的元素， append之后数组扩容，但是又重新赋值给了arr[i]，因此原来数组发生改变
		arr[i] = append(arr[i], i)
	}
	fmt.Println("")
	for i := range arr {
		fmt.Printf("%d,", len(arr[i]))
	}
	fmt.Println("")

	arr1 := [][]int{{1, 2}, {3, 4}}
	for _, sa := range arr1 {
		// sa指向的数组交换位置， sa值（指针）不变，因此原来数据交换位置
		sort.Slice(sa, func(i, j int) bool {
			return sa[i] > sa[j]
		})
	}
	for _, sa := range arr1 {
		fmt.Print("(")
		for _, v := range sa {
			fmt.Printf("%d,", v)
		}
		fmt.Print("),")
	}
	fmt.Println("")

	arr2 := [2][]int{{1, 2}, {3, 4}}
	for _, sa := range arr2 {
		// sa指向的数组交换位置， sa值（指针）不变，因此原来数据交换位置
		sort.Slice(sa, func(i, j int) bool {
			return sa[i] > sa[j]
		})
	}
	for _, sa := range arr2 {
		fmt.Print("(")
		for _, v := range sa {
			fmt.Printf("%d,", v)
		}
		fmt.Print("),")
	}
	fmt.Println("")

	arr3 := [][]int{{1, 2, 3, 4, 5, 6}, {2, 4, 6, 8, 10, 12}}
	// 更内层的arr变量名覆盖了外层的arr变量名
	for _, arr := range arr3 {
		for _, v := range arr {
			fmt.Printf("%d,", v)
		}
		fmt.Println("")
	}
}
