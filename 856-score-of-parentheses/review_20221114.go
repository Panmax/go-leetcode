package main

func scoreOfParentheses_20221114(s string) int {
	ans := 0
	var stack []int
	for _, c := range s {
		if c == '(' {
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
				stack[len(stack)-1] += last
			}
		}
	}
	return ans
}
