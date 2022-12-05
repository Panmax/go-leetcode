package main

func isValid_20221205(s string) bool {
	if len(s)%2 == 1 {
		return false
	}
	mapPair := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	var stack []byte
	for i := range s {
		if c, ok := mapPair[s[i]]; ok {
			if len(stack) > 0 && stack[len(stack)-1] == c {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}
