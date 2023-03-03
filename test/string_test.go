package test

import (
	"fmt"
	"strings"
	"testing"
)

func TestStringEqual(t *testing.T) {
	s := "Hello"
	str := "Hello1"

	ret := strings.EqualFold(s, str)
	fmt.Println(ret) //  false
	fmt.Printf("\033[0;0;31m%s\033[0m\n", "Red")
	fmt.Printf("\033[0;0;30m%s\033[0m\n", "BLACK")
}
