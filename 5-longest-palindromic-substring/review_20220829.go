package main

func longestPalindrome_20220829(s string) string {
	start, maxLength := 0, 0
	for i := range s {
		length := 1
		left := i - 1
		right := i + 1
		for left >= 0 && s[left] == s[i] {
			length++
			left--
		}
		for right < len(s) && s[right] == s[i] {
			length++
			right++
		}
		for left >= 0 && right < len(s) && s[left] == s[right] {
			length += 2
			left--
			right++
		}
		if length > maxLength {
			maxLength = length
			start = left + 1
		}
	}
	return s[start : start+maxLength]
}
