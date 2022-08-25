package main

var res []string

func generateParenthesis(n int) []string {
	if n <= 0 {
		return []string{}
	}
	res = make([]string, 0, 2*n)
	backtrace("", 0, 0, n)
	return res
}

func backtrace(cur string, left, right int, max int) {
	if len(cur) == 2*max {
		res = append(res, cur)
	}
	if left < max {
		cur += "("
		backtrace(cur, left+1, right, max)
		cur = cur[:len(cur)-1]
	}
	if left > right {
		cur += ")"
		backtrace(cur, left, right+1, max)
		cur = cur[:len(cur)-1]
	}
}
