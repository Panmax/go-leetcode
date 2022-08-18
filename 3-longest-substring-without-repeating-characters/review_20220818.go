package main

func lengthOfLongestSubstring_20220818(s string) int {
	maxLength := 0
	right := -1
	charMap := make(map[byte]int, len(s))

	for i := 0; i < len(s); i++ {
		if i > 0 {
			delete(charMap, s[i-1])
		}

		for right+1 < len(s) && charMap[s[right+1]] == 0 {
			right++
			charMap[s[right]]++
		}

		if right-i+1 > maxLength {
			maxLength = right - i + 1
		}
	}
	return maxLength
}
