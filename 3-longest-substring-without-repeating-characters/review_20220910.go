package main

func lengthOfLongestSubstring_20220910(s string) int {
	m := make(map[byte]int)
	n := len(s)
	res := 0
	ptr := -1

	for i := range s {
		if i > 0 {
			delete(m, s[i-1])
		}

		for ptr+1 < n && m[s[ptr+1]] == 0 {
			m[s[ptr+1]]++
			ptr++
		}
		res = max(res, ptr-i+1)
	}
	return res
}
