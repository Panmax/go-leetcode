package main

func calculate_20231218(s string) int {
	n := len(s)
	var res int
	var stack []int
	stack = append(stack, 1)
	sign := 1
	for i := 0; i < n; {
		if s[i] == ' ' {
			i++
		} else if s[i] == '+' {
			sign = stack[len(stack)-1]
			i++
		} else if s[i] == '-' {
			sign = -stack[len(stack)-1]
			i++
		} else if s[i] == '(' {
			stack = append(stack, sign)
			i++
		} else if s[i] == ')' {
			stack = stack[:len(stack)-1]
			i++
		} else {
			var num int
			for i < n && s[i] >= '0' && s[i] <= '9' {
				num = num*10 + int(s[i]-'0')
				i++
			}
			res += sign * num
		}
	}

	return res
}
