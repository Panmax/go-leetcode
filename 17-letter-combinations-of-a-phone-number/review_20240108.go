package main

func letterCombinations_20240108(digits string) []string {
	res = make([]string, 0)
	if len(digits) == 0 {
		return res
	}
	backtrace_20240108(digits, 0, "")
	return res
}

func backtrace_20240108(digits string, index int, comb string) {
	if index == len(digits) {
		res = append(res, comb)
	} else {
		digit := digits[index]
		letters := phoneMap[string(digit)]
		for _, letter := range letters {
			backtrace_20240108(digits, index+1, comb+string(letter))
		}
	}
}
