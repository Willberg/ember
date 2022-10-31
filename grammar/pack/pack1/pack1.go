package pack1

import (
	. "ember/grammar/pack/pack0"
	"fmt"
)

var P1 = initP1()

func init() {
	fmt.Println("init pack1")
	fmt.Println(P0)
}

func initP1() int {
	fmt.Println("init P1")
	return 1
}
