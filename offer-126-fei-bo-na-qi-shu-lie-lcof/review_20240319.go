package main

func fib_20240319(n int) int {
	if n < 2 {
		return n
	}
	p, q, r := 0, 0, 1
	for i := 2; i <= n; i++ {
		p, q = q, r
		r = (p + q) % 1000000007
	}
	return r
}
