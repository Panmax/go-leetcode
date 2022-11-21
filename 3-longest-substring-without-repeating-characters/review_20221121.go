package main

func lengthOfLongestSubstring_20221121(s string) int {
	right := -1
	charMap := make(map[byte]int, len(s))
	res := 0
	for i := range s {
		if i != 0 {
			delete(charMap, s[i-1])
		}
		for right+1 < len(s) && charMap[s[right+1]] == 0 {
			charMap[s[right+1]]++
			right++
		}
		res = max(res, right-i+1)
	}

	return res
}
