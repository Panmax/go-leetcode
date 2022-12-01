package main

import "strings"

func removeKdigits_20221201(num string, k int) string {
	var stack []byte
	for i := range num {
		for k > 0 && len(stack) > 0 && stack[len(stack)-1] > num[i] {
			k--
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, num[i])
	}
	stack = stack[:len(stack)-k]
	s := strings.TrimLeft(string(stack), "0")
	if s == "" {
		s = "0"
	}
	return s
}
