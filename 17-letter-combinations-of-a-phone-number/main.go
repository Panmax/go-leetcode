package main

var phoneMap map[string]string = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

var res []string

func letterCombinations(digits string) []string {
	res = make([]string, 0)
	if len(digits) == 0 {
		return res
	}
	backtrace(digits, 0, "")
	return res
}

func backtrace(digits string, index int, combination string) {
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
