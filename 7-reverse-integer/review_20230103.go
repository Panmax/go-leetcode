package main

import "math"

func reverseInt_20230103(x int) int {
	var res int
	for x != 0 {
		mod := x % 10
		x /= 10
		res = res*10 + mod
		if res > math.MaxInt32 || res < math.MinInt32 {
			return 0
		}
	}
	return res
}
