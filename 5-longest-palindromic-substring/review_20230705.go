package main

func longestPalindrome_20230705(s string) string {
	var maxLength int
	var minLeft int

	for i := 0; i < len(s); i++ {
		length := 1

		left := i - 1
		for left >= 0 && s[left] == s[i] {
			length += 1
			left -= 1
		}

		right := i + 1
		for right < len(s) && s[right] == s[i] {
			length += 1
			right += 1
		}

		for left >= 0 && right < len(s) && s[left] == s[right] {
			length += 2
			left -= 1
			right += 1
		}

		if length > maxLength {
			maxLength = length
			minLeft = left
		}
	}
	return s[minLeft+1 : minLeft+1+maxLength]
}
