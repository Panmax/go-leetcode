package main

func isValid_20220908(s string) bool {
	// 判断偶数
	if len(s)%2 == 1 {
		return false
	}

	mapPair := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	index := 0
	stack := make([]byte, len(s))
	for i := range s {
		if c, ok := mapPair[s[i]]; ok {
			if index > 0 && stack[index-1] == c {
				index--
			} else {
				return false
			}
		} else {
			stack[index] = s[i]
			index++
		}
	}
	return index == 0
}

func main() {
	print(isValid_20220908("()"))
}
