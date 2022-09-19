package main

func generateParenthesis_20220919(n int) []string {
	res = make([]string, 0)
	if n == 0 {
		return res
	}
	backtrace_20220919("", 0, 0, n)
	return res
}

func backtrace_20220919(cur string, left, right, n int) {
	if len(cur) == n*2 {
		res = append(res, cur)
	}
	if left < n {
		cur += "("
		backtrace_20220919(cur, left+1, right, n)
		cur = cur[:len(cur)-1]
	}
	if right < left {
		cur += ")"
		backtrace_20220919(cur, left, right+1, n)
		cur = cur[:len(cur)-1]
	}
}
