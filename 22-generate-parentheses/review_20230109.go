package main

func generateParenthesis_20230109(n int) []string {
	res = make([]string, 0, n*2)
	backtrace_20230109("", 0, 0, n)
	return res

}

func backtrace_20230109(cur string, left, right, max int) {
	if len(cur) == max*2 {
		res = append(res, cur)
	} else {
		if left < max {
			cur += "("
			backtrace_20230109(cur, left+1, right, max)
			cur = cur[:len(cur)-1]
		}
		// < left
		if right < left {
			cur += ")"
			backtrace_20230109(cur, left, right+1, max)
			cur = cur[:len(cur)-1]
		}
	}
}
