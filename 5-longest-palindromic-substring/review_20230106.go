package main

func longestPalindrome_20230106(s string) string {
	var maxLength, minLeft int

	for i := 0; i < len(s); i++ {
		length := 1
		left := i - 1
		for left >= 0 && s[left] == s[i] {
			length++
			left--
		}
		right := i + 1
		for right < len(s) && s[right] == s[i] {
			length++
			right++
		}
		for left >= 0 && right < len(s) && s[left] == s[right] {
			length += 2
			left--
			right++
		}
		// 两个值都可以在这里边修改
		if length > maxLength {
			maxLength = length
			minLeft = left
		}
	}
	return s[minLeft+1 : minLeft+1+maxLength]
}
