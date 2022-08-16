package main

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	maxLength := 0
	maxStart := 0
	for i := 0; i < len(s); i++ {
		left := i - 1
		right := i + 1
		for left >= 0 && s[left] == s[i] {
			left--
		}
		for right < len(s) && s[right] == s[i] {
			right++
		}
		for left >= 0 && right < len(s) && s[left] == s[right] {
			left--
			right++
		}
		length := right - left - 1
		if length > maxLength {
			maxLength = length
			maxStart = left
		}
	}
	return s[maxStart+1 : maxStart+1+maxLength]
}
