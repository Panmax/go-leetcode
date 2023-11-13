package main

func calculate(s string) int {
	var stack []int
	sign := 1
	var ans int
	stack = append(stack, 1)
	for i := 0; i < len(s); {
		switch s[i] {
		case ' ':
			i++
		case '+':
			sign = stack[len(stack)-1]
			i++
		case '-':
			sign = -stack[len(stack)-1]
			i++
		case '(':
			stack = append(stack, sign)
			i++
		case ')':
			stack = stack[:len(stack)-1]
			i++
		default:
			var num int
			for i < len(s) && s[i] >= '0' && s[i] <= '9' {
				num = 10*num + int(s[i]-'0')
				i++
			}
			ans += sign * num
		}
	}
	return ans
}
