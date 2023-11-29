package main

func lengthOfLongestSubstring_20221129(s string) int {
	n := len(s)
	m := make(map[byte]bool)
	right := -1
	var res int
	for i := 0; i < n; i++ {
		if i > 0 {
			delete(m, s[i-1])
		}
		for right+1 < n && !m[s[right+1]] {
			m[s[right+1]] = true
			right++
		}
		res = max(res, right-i+1)
	}
	return res
}
