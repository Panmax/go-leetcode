package main

func letterCombinations_20230213(digits string) []string {
	res = make([]string, 0)
	if digits == "" {
		return res
	}
	backtrace_20230213(digits, 0, "")
	return res
}

func backtrace_20230213(digits string, index int, combination string) {
	if index == len(digits) {
		res = append(res, combination)
	} else {
		digit := digits[index]
		for _, c := range phoneMap[string(digit)] {
			backtrace_20230213(digits, index+1, combination+string(c))
		}
	}
}
