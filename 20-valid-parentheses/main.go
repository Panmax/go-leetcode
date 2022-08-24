package main

func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}

	mapPair := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}

	index := 0
	stack := make([]byte, len(s))
	for i := range s {
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
