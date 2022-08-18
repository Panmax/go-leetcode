package main

func longestPalindrome_20220817(s string) string {
	if len(s) < 2 {
		return s
	}
	maxLength := 0
	maxLeft := 0
	for i := 0; i < len(s); i++ {
		// 注意这里从1开始
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
		for left >= 0 && right < len(s) && s[left] == s[right] {
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
