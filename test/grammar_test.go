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

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x)
	defer fmt.Printf("defer f(%d)\n", x)
	f(x - 1)
}

func TestPanic(t *testing.T) {
	// 执行过的延迟语句一定会执行，即使出现宕机
	f(3)
}

func fr(x int) {
	defer func() {
		p := recover()
		if p != nil {
			fmt.Printf("%v", fmt.Errorf("error: %v\n", p))
		}
		fmt.Printf("defer fr(%d)\n", x)
		panic(p)
	}()
	fmt.Printf("fr(%d)\n", x+0/x)
	fr(x - 1)
}

func TestRecover(t *testing.T) {
	fr(3)
}

func h(x, y int) {
	defer fmt.Printf("defer h(%d)\n", x)
	fmt.Printf("h(%d)\n", x/y)
	fmt.Printf("panic h(%d)\n", x)
}

func h1(x, y int) {
	defer func() {
		fmt.Printf("defer h1(%d)\n", x)
		if p := recover(); p != nil {
			fmt.Printf("%v\n", fmt.Errorf("%v", p))
		}
	}()
	fmt.Printf("h1(%d)\n", x/y)
	fmt.Printf("h1(%d) panic\n", x/y)
}

func g(x int) {
	defer fmt.Printf("defer g(%d)\n", x)
	h1(x, 0)
	fmt.Printf("panic h1(%d)\n", x)
	h(x, 0)
	fmt.Printf("g(%d)\n", x)
}

func TestPanic1(t *testing.T) {
	// 宕机之后代码不会执行， 宕机会传递, 除非recover
	g(3)
}
