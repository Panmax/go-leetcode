package main

import "strings"

func removeKdigits_20220927(num string, k int) string {
	stack := []byte{}
	for i := range num {
		// for !!!
		for k > 0 && len(stack) > 0 && num[i] < stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			k--
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
