package main

func lengthOfLongestSubstring_20220822(s string) int {
	right := -1
	charMap := make(map[byte]int)
	maxLength := 0

	for i := 0; i < len(s); i++ {
		if i > 0 {
			delete(charMap, s[i-1])
		}

		for right+1 < len(s) && charMap[s[right+1]] == 0 {
			charMap[s[right+1]]++
			right++
		}

		maxLength = max(right-i+1, maxLength)
	}
	return maxLength
}
