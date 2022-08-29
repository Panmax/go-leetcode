package main

func generateParenthesis_20220829(n int) []string {
	if n <= 0 {
		return []string{}
	}
	res = make([]string, 0, 2*n)
	backtrace("", 0, 0, n)
	return res
}

func backtrace_20220829(cur string, left, right int, max int) {
	if 2*max == len(cur) {
		res = append(res, cur)
	}
	if left < max {
		cur += "("
		backtrace(cur, left+1, right, max)
		cur = cur[:len(cur)-1]
	}
	if right < left {
		cur += ")"
		backtrace(cur, left, right+1, max)
		cur = cur[:len(cur)-1]
	}
}
