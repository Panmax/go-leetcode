package main

import "strings"

func isPalindrome_20230306(s string) bool {
	s = strings.ToLower(s)
	start, end := 0, len(s)-1
	for start < end {
		if !isChar_20230306(s[start]) {
			start++
			continue
		}
		if !isChar_20230306(s[end]) {
			end--
			continue
		}
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}

func isChar_20230306(c byte) bool {
	// 0-9
	if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') {
		return true
	}
	return false
}
