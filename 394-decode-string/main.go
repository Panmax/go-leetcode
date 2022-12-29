package main

import (
	"strconv"
	"strings"
)

func decodeString(s string) string {
	var stack []string
	ptr := 0
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
			for stack[len(stack)-1] != "[" {
				sub = append(sub, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1]
			for i := 0; i < len(sub)/2; i++ {
				sub[i], sub[len(sub)-i-1] = sub[len(sub)-i-1], sub[i]
			}
			count, _ := strconv.Atoi(stack[len(stack)-1])
			stack = stack[:len(stack)-1]
			repeat := strings.Repeat(getString(sub), count)
			stack = append(stack, repeat)
		}
	}
	return getString(stack)
}

func getDigits(s string, ptr *int) string {
	d := ""
	for ; *ptr < len(s) && (s[*ptr] >= '0' && s[*ptr] <= '9'); *ptr++ {
		d += string(s[*ptr])
	}
	return d
}

func getString(s []string) string {
	ret := ""
	for i := range s {
		ret += s[i]
	}
	return ret
}
