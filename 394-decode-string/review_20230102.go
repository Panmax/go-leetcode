package main

import (
	"strconv"
	"strings"
)

func decodeString_20230102(s string) string {
	var stack []string
	ptr := 0
	for ptr < len(s) {
		c := s[ptr]
		if c >= '0' && c <= '9' {
			d := getDigits(s, &ptr)
			stack = append(stack, d)
		} else if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '[' {
			stack = append(stack, string(c))
			ptr++
		} else {
			ptr++
			var sub []string
			for stack[len(stack)-1] != "[" {
				sub = append(sub, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1]
			for i := 0; i < len(sub)/2; i++ {
				sub[i], sub[len(sub)-i-1] = sub[len(sub)-i-1], sub[i]
			}
			repeatCount, _ := strconv.Atoi(stack[len(stack)-1])
			stack = stack[:len(stack)-1]
			repeatStr := strings.Repeat(getString(sub), repeatCount)
			stack = append(stack, repeatStr)
		}
	}

	return getString(stack)
}
