package test

import (
	"fmt"
	"testing"
)

func TestRange(t *testing.T) {
	s := "abcs好"
	for i, v := range s {
		// v是rune(int32),可以var k rune = 'c'这样定义，因此可以直接v=='a'('a','c'是byte)判断， s[i]是byte(uint8), 汉字的 '好' 是rune
		if byte(v) != 'a' {
			if v == 'c' {
				fmt.Printf("%c\n", s[i])
			} else if s[i] == byte('s') {
				fmt.Printf("%c\n", s[i])
			} else if rune(s[i]) == 'b' {
				fmt.Printf("%c\n", rune(s[i]))
			} else if v == '好' {
				fmt.Printf("%c\n", v)
			}
		}
	}
}
