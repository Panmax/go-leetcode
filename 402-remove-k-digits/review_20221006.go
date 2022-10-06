package main

import "strings"

func removeKdigits_20221006(num string, k int) string {
	var stack []byte
	for i := range num {
		for k > 0 && len(stack) > 0 && num[i] < stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, num[i])
	}
	stack = stack[:len(stack)-k]
	str := strings.TrimLeft(string(stack), "0")
	if str == "" {
		str = "0"
	}
	return str
}
