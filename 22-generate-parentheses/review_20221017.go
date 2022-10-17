package main

func generateParenthesis_20221017(n int) []string {
	res = make([]string, 0, n*2)
	backtrace_20221017("", 0, 0, n)
	return res
}

func backtrace_20221017(cur string, left, right, max int) {
	if len(cur) == 2*max {
		res = append(res, cur)
	} else {
		if left < max {
			cur += "("
			backtrace_20221017(cur, left+1, right, max)
			cur = cur[:len(cur)-1]
		}
		if right < left {
			cur += ")"
			backtrace_20221017(cur, left, right+1, max)
			cur = cur[:len(cur)-1]
		}
	}
}
