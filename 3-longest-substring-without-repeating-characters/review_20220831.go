package main

func lengthOfLongestSubstring_20220831(s string) int {
	maxLength := 0
	// 初始值为-1
	rightPtr := -1
	charMap := make(map[byte]int)
	for i := range s {
		if i > 0 {
			// 删掉前一个元素
			delete(charMap, s[i-1])
		}

		for rightPtr+1 < len(s) && charMap[s[rightPtr+1]] == 0 {
			charMap[s[rightPtr+1]]++
			rightPtr++
		}
		// 长度是 rightPtr - i + 1
		maxLength = max(maxLength, rightPtr-i+1)
	}
	return maxLength
}
