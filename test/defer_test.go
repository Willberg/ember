package test

import (
	"fmt"
	"testing"
)

func getOne() int {
	defer fmt.Println("Yes")
	return 1 + 1
}

func TestDefer(t *testing.T) {
	defer fmt.Println("fourth")
	defer fmt.Println("third")
	defer fmt.Println("second ")
	defer fmt.Println("first ")
	fmt.Println(getOne())
}
