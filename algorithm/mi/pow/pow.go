package pow

// 费马小定理: 当p为质数时， 有a^(p-1) = 1 (mod p). 可以由(b + 1)^p 展开证明: (b+1)^p = b+1 (mod p), 然后令a = b+1得到
// a * a^(p-2) = 1 (mod p), 所以 a^(p-2) 就是 a % p 的乘法逆元
func FermatInv(a, b int) int {
	return PowMod(a, b-2, b)
}

// 快速幂取模(a^n % mod), 可以折半快速计算
func PowMod(a, n, mod int) int {
	ret := 1
	for n > 0 {
		if n&1 == 1 {
			ret = ret * a % mod
		}
		a = a * a % mod
		n >>= 1
	}
	return ret
}
