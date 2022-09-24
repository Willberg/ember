package linear

import (
	. "ember/algorithm/mi/exgcd"
)

// i ^(-1) 表示 i mod p 的逆元
// 1 * 1 mod p = 1, 1的逆元就是1
// k = p / i 向下取整, j = p % i => p = k * i + j => k*i + j = 0 (mod p)
// 两边乘以 i^(-1) * j ^(-1) => k*j^(-1) + i ^ (-1) = 0 (mod p)
// 左边加上 -p * j ^(-1) => i^(-1) = (p-k)*j^(-1), 其中k < p,可以保证逆元是正数.
// j < i 所以 j 在之前已被算出
// 计算1-n每个数 mod p 的逆元， 如果 i 与 p 不互素时不存在相应的逆元, 所以p一般是一个较大的素数
func Inv(n, p int) []int {
	inv := make([]int, n+1)
	inv[1] = 1
	for i := 2; i <= n; i++ {
		inv[i] = (p - p/i) * inv[p%i] % p
	}
	return inv[1:]
}

// 线性同余方程 ax = 1 (mod p)
// 计算数组a的每个数的逆元
// a(n) 表示第n个数, s(n) 表示前n个数的前缀积 mod p， sv(n)表示sn的逆元, inv(a)表示a的逆元
// s(n) * sv(n) = 1 (mod p) => s(n-1) * a(n) * sv(n) = 1 (mod p), 所以 sv(n-1) = a(n) * sv(n) % p, 可以求出前缀积的逆元
// s(n) * sv(n) = 1 (mod p) => a(n) * s(n-1) * sv(n) = 1 (mod p), 所以 inv(a(n)) = s(n-1) * sv(n) % p, 可以求出a(n)的逆元
// 令s(0) = 1, sv(1) = inv(a(1))
func Inva(a []int, p int) []int {
	n := len(a)
	s, sv, inv := make([]int, n+1), make([]int, n+1), make([]int, n+1)
	s[0] = 1
	for i := 1; i <= n; i++ {
		s[i] = s[i-1] * a[i-1] % p
	}
	sn, pp, k := s[n], p, Gcd(s[n], p)
	if k != 1 {
		sn /= k
		pp /= k
	}
	sv[n] = ExGcdInv(sn, pp)
	for i := n; i >= 1; i-- {
		sv[i-1] = a[i-1] * sv[i] % p
	}
	for i := 1; i <= n; i++ {
		inv[i] = s[i-1] * sv[i] % p
	}
	return inv[1:]
}
