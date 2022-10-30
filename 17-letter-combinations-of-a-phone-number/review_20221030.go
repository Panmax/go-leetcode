package main

func letterCombinations_202230(digits string) []string {
	res = make([]string, 0)
	if digits == "" {
		return res
	}
	backtrace_20220930(digits, 0, "")
	return res
}

func backtrace_20220930(digits string, index int, combination string) {
	if index == len(digits) {
		res = append(res, combination)
	} else {
		digit := digits[index]
		for _, s := range phoneMap[string(digit)] {
			backtrace_20220930(digits, index+1, combination+string(s))
		}
	}
}
