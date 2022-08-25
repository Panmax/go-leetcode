package main

func isValid_20220825(s string) bool {
	if len(s)%2 == 1 {
		return false
	}
	mapPair := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}
	stack := make([]byte, len(s))
	index := 0
	for i := 0; i < len(s); i++ {
		if mapPair[s[i]] > 0 {
			if index > 0 && stack[index-1] == mapPair[s[i]] {
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
