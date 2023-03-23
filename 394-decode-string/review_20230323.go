package main

import (
	"strconv"
	"strings"
)

func decodeString_20230323(s string) string {
	var ptr int
	var stack []string
	for ptr < len(s) {
		if s[ptr] >= '0' && s[ptr] <= '9' {
			d := getDigits(s, &ptr)
			stack = append(stack, d)
		} else if (s[ptr] >= 'a' && s[ptr] <= 'z') || (s[ptr] >= 'A' && s[ptr] <= 'Z') || s[ptr] == '[' {
			stack = append(stack, string(s[ptr]))
			ptr++
		} else {
			ptr++
			var sub []string
			for len(stack) > 0 {
				if stack[len(stack)-1] == "[" {
					break
				}
				sub = append(sub, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1]
			for i := 0; i < len(sub)/2; i++ {
				sub[i], sub[len(sub)-1-i] = sub[len(sub)-1-i], sub[i]
			}
			d, _ := strconv.Atoi(stack[len(stack)-1])
			stack = stack[:len(stack)-1]
			repeatStr := strings.Repeat(getString(sub), d)
			stack = append(stack, repeatStr)
		}
	}
	return getString(stack)
}
