package main

func isValid_20230710(s string) bool {
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
		if mapPair[s[i]] > 0 {
			if len(stack) > 0 && stack[len(stack)-1] == mapPair[s[i]] {
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
