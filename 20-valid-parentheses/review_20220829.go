package main

func isValid_20220829(s string) bool {
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
		if c, ok := mapPair[s[i]]; ok {
			if index-1 >= 0 && stack[index-1] == c {
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
