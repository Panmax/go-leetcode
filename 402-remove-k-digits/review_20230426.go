package main

import "strings"

func removeKdigits_20230426(num string, k int) string {
	var stack []byte
	for i := range num {
		for k > 0 && len(stack) > 0 && num[i] < stack[len(stack)-1] {
			k--
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, num[i])
	}
	stack = stack[:len(stack)-k]
	s := string(stack)
	s = strings.TrimLeft(s, "0")
	if s == "" {
		return "0"
	}

	return s
}
