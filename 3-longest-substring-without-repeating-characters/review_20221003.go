package main

func lengthOfLongestSubstring_20221003(s string) int {
	right := -1
	res := 0
	m := make(map[byte]int, 0)
	for i := range s {
		if i > 0 {
			// delete i-1
			delete(m, s[i-1])
		}
		for right+1 < len(s) && m[s[right+1]] == 0 {
			right++
			m[s[right]]++
		}
		res = max(res, right-i+1)
	}
	return res
}
