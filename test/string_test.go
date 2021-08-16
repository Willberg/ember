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
}
