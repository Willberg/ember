package test

import (
	"fmt"
	"testing"
)

func TestFunc(t *testing.T) {
	v := testCallback(1, callback)
	fmt.Println(v)
}

func callback(t int) int {
	fmt.Println(t)
	return t + 1
}

func testCallback(v int, cb func(int) int) int {
	return cb(v)
}
