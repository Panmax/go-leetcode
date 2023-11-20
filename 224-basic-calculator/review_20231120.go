package main

func calculate_20231120(s string) int {
	var stack []int
	stack = append(stack, 1)
	sign := 1
	var ans int
	n := len(s)
	for i := 0; i < n; {
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
			for i < n && s[i] >= '0' && s[i] <= '9' {
				num = num*10 + int(s[i]-'0')
				i++
			}
			ans += sign * num
		}
	}
	return ans
}
