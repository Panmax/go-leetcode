package main

import "strings"

func isPalindrome_20231106(s string) bool {
	s = strings.ToLower(s)
	i, j := 0, len(s)-1
	for i < j {
		if !isChar_20231106(s[i]) {
			i++
			continue
		}
		if !isChar_20231106(s[j]) {
			j--
			continue
		}
		if s[i] == s[j] {
			i++
			j--
			continue
		}
		return false
	}
	return true
}

func isChar_20231106(c byte) bool {
	if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') {
		return true
	}
	return false
}
