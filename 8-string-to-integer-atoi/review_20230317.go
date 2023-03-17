package main

import "math"

func myAtoi_20230317(s string) int {
	for len(s) > 0 {
		if s[0] == ' ' {
			s = s[1:]
		} else {
			break
		}
	}

	res := 0
	neg := false
	for i, c := range s {
		if i == 0 && c == '-' {
			neg = true
			continue
		}
		if i == 0 && c == '+' {
			continue
		}
		if c >= '0' && c <= '9' {
			res = res*10 + int(c-'0')
			if neg && -res <= math.MinInt32 {
				return math.MinInt32
			}
			if !neg && res >= math.MaxInt32 {
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
