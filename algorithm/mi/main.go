package main

import (
	. "ember/algorithm/mi/exgcd"
	. "ember/algorithm/mi/linear"
	. "ember/algorithm/mi/pow"
	"fmt"
)

func testGcd() {
	a, b := 111, 76
	fmt.Println(Gcd(a, b))
}

func testExgcd() {
	a, b := 111, int(1e9+7)
	if Gcd(a, b) == 1 {
		fmt.Println(ExGcdInv(a, b))
	}
}

func testPowMod() {
	a, n, mod := 2, 15, 5
	fmt.Println(PowMod(a, n, mod))
}

func testIsPrime() {
	fmt.Println(IsPrime(int(1e2 + 7)))
}

func testFermatInv() {
	a, b := 111, int(1e9+7)
	if Gcd(a, b) == 1 {
		fmt.Println(FermatInv(a, b))
	}
}

func testLinearInv() {
	inv := Inv(111, int(1e9+7))
	for i, v := range inv {
		if i > 0 && i%20 == 0 {
			fmt.Println("")
		}
		fmt.Printf("%d,", v)
	}
	fmt.Println("")
}

func testArrInv(a []int, b int) {
	inv := Inva(a, b)
	for i, v := range inv {
		if i > 0 && i%20 == 0 {
			fmt.Println("")
		}
		fmt.Printf("%d,", v)
	}
	fmt.Println("")
}

func testLinearInva() {
	var a []int
	for i := 1; i <= 111; i++ {
		a = append(a, i)
	}
	testArrInv(a, int(1e9+7))
}

func testLinearInvn() {
	var a []int
	for i := 1; i <= 100; i++ {
		if IsPrime(i) {
			a = append(a, i)
		}
	}
	for i, v := range a {
		if i > 0 && i%20 == 0 {
			fmt.Println("")
		}
		fmt.Printf("%d,", v)
	}
	fmt.Println("")
	testArrInv(a, int(1e9+7))
}

func testLinearInvno() {
	var a []int
	for i := 1; i <= 100; i++ {
		a = append(a, i)
	}
	for i, v := range a {
		if i > 0 && i%20 == 0 {
			fmt.Println("")
		}
		fmt.Printf("%d,", v)
	}
	fmt.Println("")
	testArrInv(a, 107)
}

func main() {
	testGcd()
	testIsPrime()
	testExgcd()
	testPowMod()
	testFermatInv()
	testLinearInv()
	testLinearInva()
	testLinearInvn()
	testLinearInvno()
}
