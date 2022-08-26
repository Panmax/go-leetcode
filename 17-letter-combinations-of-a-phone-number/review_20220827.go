package main

func letterCombinations_20220827(digits string) []string {
	res = make([]string, 0)
	if len(digits) == 0 {
		return res
	}

	backtrace_202208027(digits, 0, "")

	return res
}

func backtrace_202208027(digits string, index int, combination string) {
	if index == len(digits) {
		res = append(res, combination)
	} else {
		digit := digits[index]
		letters := phoneMap[string(digit)]
		for _, letter := range letters {
			backtrace(digits, index+1, combination+string(letter))
		}
	}
}
