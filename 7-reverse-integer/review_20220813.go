package main

import "math"

func reverseInt_20220813(x int) int {
	res := 0
	for x != 0 {
		last := x % 10
		res = res*10 + last
		if res > math.MaxInt32 || res < math.MinInt32 {
			return 0
		}
		x /= 10
	}

	return res
}
