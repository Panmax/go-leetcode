package main

func letterCombinations_20220918(digits string) []string {
	res = make([]string, 0)
	if len(digits) == 0 {
		return res
	}
	backtrace_20220918(digits, 0, "")
	return res
}

func backtrace_20220918(digits string, index int, combation string) {
	// if else
	if index == len(digits) {
		res = append(res, combation)
	} else {
		digit := digits[index]
		letters := phoneMap[string(digit)]
		for _, l := range letters {
			backtrace_20220918(digits, index+1, combation+string(l))
		}
	}
}
