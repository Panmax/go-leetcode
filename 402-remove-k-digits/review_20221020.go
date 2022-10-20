package main

import "strings"

func removeKdigits_20221020(num string, k int) string {
	var stack []byte
	for i := range num {
		for k > 0 && len(stack) > 0 && num[i] < stack[len(stack)-1] {
			k--
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, num[i])
	}
	stack = stack[:len(stack)-k]
	res := string(stack)
	res = strings.TrimLeft(res, "0")
	if res == "" {
		res = "0"
	}
	return res
}
