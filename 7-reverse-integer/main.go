package main

import (
	"fmt"
	"math"
)

func reverseInt(num int) int {
	result := 0
	for num != 0 {
		last := num % 10
		num = num / 10
		result = result*10 + last
		if result > math.MaxInt32 || result < math.MinInt32 {
			return 0
		}
	}
	return result
}

// https://leetcode.cn/problems/reverse-integer/
func main() {
	fmt.Println(reverseInt(0))
	fmt.Println(reverseInt(1234))
	fmt.Println(reverseInt(-321))
	fmt.Println(reverseInt(-120))
	fmt.Println(reverseInt(120))
}
