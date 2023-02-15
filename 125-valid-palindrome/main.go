package main

import "strings"

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	start, end := 0, len(s)-1
	for start < end {
		if !isChar(s[start]) {
			start++
			continue
		}
		if !isChar(s[end]) {
			end--
			continue
		}
		if s[start] == s[end] {
			start++
			end--
			continue
		}
		return false
	}
	return true
}

func isChar(c byte) bool {
	if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') {
		return true
	}
	return false
}

func main() {
	isPalindrome("A man, a plan, a canal: Panama")
}
