package main

import (
	"math"
)

func myAtoi(s string) int {
	// 去掉空格前缀
	for len(s) > 0 {
		if s[0] == ' ' {
			s = s[1:]
		} else {
			break
		}
	}

	negative := false
	res := 0
	for i, c := range s {
		if i == 0 && c == '-' {
			negative = true
			continue
		} else if i == 0 && c == '+' {
			continue
		}
		if c >= '0' && c <= '9' {
			res = res*10 + int(c-'0')
			if negative {
				if -res <= math.MinInt32 {
					return math.MinInt32
				}
			} else {
				if res >= math.MaxInt32 {
					return math.MaxInt32
				}
			}
		} else {
			break
		}
	}
	if negative {
		res = -res
	}
	return res
}

func main() {
	println(myAtoi("-2147483647"))
}
