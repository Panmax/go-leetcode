package main

func letterCombinations_20220902(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	res = make([]string, 0)
	backtrace_20220902(digits, 0, "")
	return res
}

func backtrace_20220902(digits string, index int, combination string) {
	// 下边是 if else
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
