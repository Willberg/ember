package test

import (
	"fmt"
	"strings"
	"testing"
)

const (
	RED   = "\033[31m"
	BLACK = "\033[30m"
	RESET = "\033[0m"
)

func TestStringEqual(t *testing.T) {
	s := "Hello"
	str := "Hello1"

	ret := strings.EqualFold(s, str)
	fmt.Println(ret) //  false
	fmt.Printf(RED+"%s"+RESET+"\n", "Red")
	fmt.Printf(BLACK+"%s"+RESET+"\n", "BLACK")
	fmt.Println(add("a", "b"))
	fmt.Println(add(3, 4))
}

func add[T int | string](a, b T) T {
	return a + b
}
