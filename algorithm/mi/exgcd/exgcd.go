package exgcd

// 乘法逆元，拓展欧几里得算法求ax+by=gcd(a,b)的解 gcd(a, b) == gcd(b, a%b)
// x = y', y = x' - ky', k = a/b的向下取整
func ExGcd(a, b int, x, y *int) int {
	// x, y 最终被赋值为可行解
	if b == 0 {
		// b=0， gcd(a,b) = a, ax = a, x=1, y = 0 为可行解
		*x = 1
		*y = 0
		return a
	}
	d := ExGcd(b, a%b, x, y)
	t := *x
	*x = *y
	*y = t - a/b*(*y)
	return d
}

// a, b 应该互质
func ExGcdInv(a, b int) int {
	var x, y int
	ExGcd(a, b, &x, &y)
	// 保证x为正数， a(x+kb) + b(y-ka) = gcd(a, b)
	v := x
	for k := 1; v < 0; k++ {
		v = (x + k*b) % b
	}
	return v
}

func Gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return Gcd(b, a%b)
}

func IsPrime(a int) bool {
	for i := 2; i*i <= a; i++ {
		if a%i == 0 {
			return false
		}
	}
	return true
}
