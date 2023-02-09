package main

func scoreOfParentheses_20230209(s string) int {
	ans := 0
	var stack []int
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, 0)
		} else {
			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if last == 0 {
				last = 1
			} else {
				last *= 2
			}
			if len(stack) == 0 {
				ans += last
			} else {
				stack[len(stack)-1] = stack[len(stack)-1] + last
			}
		}
	}
	return ans
}
