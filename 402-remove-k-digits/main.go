package main

import "strings"

func removeKdigits(num string, k int) string {
	var stack []byte
	for i := range num {
		// for
		for k > 0 && len(stack) > 0 && num[i] < stack[len(stack)-1] {
			k--
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, num[i])
	}
	stack = stack[:len(stack)-k]

	res := strings.TrimLeft(string(stack), "0")
	if res == "" {
		res = "0"
	}
	return res
}
