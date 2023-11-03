package main

func scoreOfParentheses_20231103(s string) int {
	var stack []int
	var ans int
	for i := range s {
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
