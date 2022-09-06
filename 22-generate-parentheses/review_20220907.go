package main

func generateParenthesis_20220907(n int) []string {
	res = make([]string, 0, 2*n)
	backtrace_20220907("", 0, 0, n)
	return res
}

func backtrace_20220907(cur string, left, rignt int, max int) {
	if len(cur) == max*2 {
		res = append(res, cur)
	} else {
		if left < max {
			cur += "("
			backtrace(cur, left+1, rignt, max)
			cur = cur[:len(cur)-1]
		}
		if rignt < left {
			cur += ")"
			backtrace(cur, left, rignt+1, max)
			cur = cur[:len(cur)-1]
		}
	}
}
