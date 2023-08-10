package main

func generateParenthesis_20230810(n int) []string {
	if n <= 0 {
		return nil
	}
	res = make([]string, 0)
	backtrace_20230810("", 0, 0, n)
	return res
}

func backtrace_20230810(cur string, left int, right int, max int) {
	if len(cur) == 2*max {
		res = append(res, cur)
		return
	}
	if left < max {
		cur += "("
		backtrace_20230810(cur, left+1, right, max)
		cur = cur[:len(cur)-1]
	}
	if left > right {
		cur += ")"
		backtrace_20230810(cur, left, right+1, max)
		cur = cur[:len(cur)-1]
	}
}
