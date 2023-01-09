package main

import (
	"strconv"
	"strings"
)

func decodeString_20230102(s string) string {
	var stack []string
	var ptr int
	for ptr < len(s) {
		if s[ptr] >= '0' && s[ptr] <= '9' {
			d := getDigits(s, &ptr)
			stack = append(stack, d)
		} else if (s[ptr] >= 'a' && s[ptr] <= 'z') || (s[ptr] >= 'A' && s[ptr] <= 'Z') || s[ptr] == '[' {
			stack = append(stack, string(s[ptr]))
			ptr++
		} else {
			// ptr++
			ptr++
			var sub []string
			for len(stack) > 0 && stack[len(stack)-1] != "[" {
				sub = append(sub, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1]
			for i := 0; i < len(sub)/2; i++ {
				sub[i], sub[len(sub)-i-1] = sub[len(sub)-i-1], sub[i]
			}
			s := getString(sub)
			repeat, _ := strconv.Atoi(stack[len(stack)-1])
			stack = stack[:len(stack)-1]
			stack = append(stack, strings.Repeat(s, repeat))
		}
	}
	return getString(stack)
}
