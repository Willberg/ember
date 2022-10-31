package main

// import同一个包多次，只会执行一次
import (
	. "ember/grammar/pack/pack1"
	. "ember/grammar/pack/pack2"
	"fmt"
)

func main() {
	fmt.Println(P1, P2)
}
