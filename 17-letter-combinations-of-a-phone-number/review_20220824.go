package main

func letterCombinations_20220824(digits string) []string {
	res = make([]string, 0)
	if len(digits) == 0 {
		return []string{}
	}
	backtrace(digits, 0, "")
	return res
}

func backtrace_20220824(digits string, index int, combination string) {
	if index == len(digits) {
		res = append(res, combination)
	} else {
		digit := string(digits[index])
		letters := phoneMap[digit]
		for _, letter := range letters {
			backtrace(digits, index+1, combination+string(letter))
		}
	}
}
