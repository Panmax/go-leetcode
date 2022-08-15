package main

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func lengthOfLongestSubstring(s string) int {
	m := map[byte]int{}
	n := len(s)
	res := 0
	rightPtr := -1
	for i := 0; i < n; i++ {
		if i != 0 {
			delete(m, s[i-1])
		}

		for rightPtr+1 < n && m[s[rightPtr+1]] == 0 {
			m[s[rightPtr+1]]++
			rightPtr++
		}

		res = max(res, rightPtr-i+1)
	}
	return res
}
