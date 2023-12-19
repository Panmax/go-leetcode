package main

func beautySum(s string) int {
	var ans int
	for i := range s {
		counts := [26]int{}
		var maxFreq int
		for _, val := range s[i:] {
			counts[val-'a']++
			maxFreq = max(maxFreq, counts[val-'a'])
			minFreq := len(s)
			for _, c := range counts {
				// 加判断
				if c > 0 {
					minFreq = min(minFreq, c)
				}
			}
			ans += maxFreq - minFreq
		}
	}
	return ans
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}
