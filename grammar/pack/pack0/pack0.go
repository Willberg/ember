package pack0

import "fmt"

var P0 = initP0()

func init() {
	fmt.Println("init pack0")
}

// 可以在同一个文件中定义两次，两次都会执行（应该是合并到一个init函数中，然后执行）
func init() {
	fmt.Println("init pack0 again")
}

func initP0() int {
	fmt.Println("init P0")
	return 0
}
