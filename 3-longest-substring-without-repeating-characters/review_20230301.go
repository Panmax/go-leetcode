package main

func lengthOfLongestSubstring_20230301(s string) int {
	m := make(map[byte]int)
	right := -1
	res := 0
	for i := 0; i < len(s); i++ {
		if i > 0 {
			delete(m, s[i-1])
		}
		for right < len(s) && m[s[right]] == 0 {
			m[s[right]]++
			right++
		}
		// 因为 right 多往后走一步，计算长度时不需要+1
		res = max(res, right-i)
	}
	return res
}
