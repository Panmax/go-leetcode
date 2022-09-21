package main

func longestPalindrome_20220921(s string) string {
	var maxLength, maxLeft int
	for i := range s {
		length := 1
		left := i - 1
		right := i + 1
		for left >= 0 && s[left] == s[i] {
			left--
			length++
		}
		for right < len(s) && s[right] == s[i] {
			right++
			length++
		}
		for left >= 0 && right < len(s) && s[right] == s[left] {
			left--
			right++
			length += 2
		}
		if length > maxLength {
			maxLength = length
			maxLeft = left
		}
	}
	return s[maxLeft+1 : maxLeft+1+maxLength]
}
