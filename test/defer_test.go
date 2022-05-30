package test

import (
	"fmt"
	"testing"
)

func TestDefer(t *testing.T) {
	defer fmt.Println("third")
	defer fmt.Print("second ")
	defer fmt.Print("first ")
}
