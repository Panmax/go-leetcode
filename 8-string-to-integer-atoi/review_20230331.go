package main

import "math"

func myAtoi_20230331(s string) int {
	neg := false
	var res int
	// len(s) > 0
	for len(s) > 0 {
		if s[0] == ' ' {
			s = s[1:]
		} else {
			break
		}
	}

	for i := 0; i < len(s); i++ {
		if i == 0 && s[i] == '-' {
			neg = true
			continue
		}
		if i == 0 && s[i] == '+' {
			continue
		}
		if s[i] >= '0' && s[i] <= '9' {
			res = res*10 + int(s[i]-'0')
			if neg && -res < math.MinInt32 {
				return math.MinInt32
			}
			if !neg && res > math.MaxInt32 {
				return math.MaxInt32
			}
		} else {
			break
		}
	}
	if neg {
		res = -res
	}

	return res
}
