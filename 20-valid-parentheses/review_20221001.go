package main

func isValid_20221001(s string) bool {
	if len(s)%2 == 1 {
		return false
	}
	mapPair := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	var stack []byte
	for _, b := range s {
		if c, ok := mapPair[byte(b)]; ok {
			if len(stack) > 0 && stack[len(stack)-1] == c {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else {
			stack = append(stack, byte(b))
		}
	}
	return len(stack) == 0
}

func main() {
	isValid_20221001("()")
}
