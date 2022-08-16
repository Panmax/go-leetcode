package main

func lengthOfLongestSubstring_20220816(s string) int {
	charMap := make(map[byte]int, len(s))
	n := len(s)

	rightPtr := -1
	res := 0
	for i := range s {
		if i != 0 {
			delete(charMap, s[i-1])
		}

		for rightPtr+1 < n && charMap[s[rightPtr+1]] == 0 {
			charMap[s[rightPtr+1]]++
			rightPtr++
		}

		res = max(res, rightPtr-i+1)
	}
	return res
}
