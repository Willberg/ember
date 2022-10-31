package pack2

import (
	. "ember/grammar/pack/pack0"
	"fmt"
)

var P2 = initP2()

func init() {
	fmt.Println("init pack2")
	fmt.Println(P0)
}

func initP2() int {
	fmt.Println("init P2")
	return 2
}
