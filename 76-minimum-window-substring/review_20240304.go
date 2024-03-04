package main

import "math"

func minWindow_20240304(s string, t string) string {
	countMap, tMap := map[byte]int{}, map[byte]int{}
	for i := range t {
		tMap[t[i]]++
	}
	check := func() bool {
		for k := range tMap {
			if countMap[k] < tMap[k] {
				return false
			}
		}
		return true
	}
	length := math.MaxInt32
	sL, sR := -1, -1
	for l, r := 0, 0; r < len(s); r++ {
		if tMap[s[r]] > 0 {
			countMap[s[r]]++
		}
		for check() && l < r {
			if r-l+1 < length {
				length = r - l + 1
				sL, sR = l, r+1 // r+1 或者 l+length
			}
			if tMap[s[l]] > 0 {
				countMap[s[l]]--
			}
			l++
		}
	}

	if sL == -1 {
		return ""
	}
	return s[sL:sR]
}
