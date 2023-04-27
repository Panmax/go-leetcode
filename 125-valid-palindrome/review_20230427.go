package main

import "strings"

func isPalindrome_20230427(s string) bool {
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
