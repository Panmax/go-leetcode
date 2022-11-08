package main

func scoreOfParentheses(s string) int {
	var ant int
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
				ant += last
			} else {
				stack[len(stack)-1] += last
			}
		}
	}
	return ant
}
