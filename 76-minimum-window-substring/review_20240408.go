package main

import (
	"fmt"
	"math"
)

func minWindow_20240408(s string, t string) string {
	countMap, tMap := map[byte]int{}, map[byte]int{}
	for i := range t {
		tMap[t[i]]++
	}
	var check func() bool
	check = func() bool {
		for c := range tMap {
			if countMap[c] < tMap[c] {
				return false
			}
		}
		return true
	}

	sL, sR := -1, -1
	length := math.MaxInt32
	for l, r := 0, 0; r < len(s); r++ {
		if tMap[s[r]] > 0 {
			countMap[s[r]]++
		}
		for check() && l <= r {
			if r-l+1 < length {
				length = r - l + 1
				sL, sR = l, r+1
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

func main() {
	fmt.Print(minWindow_20240408("a", "a"))
}
