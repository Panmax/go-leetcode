package main

func beautySum_20230224(s string) int {
	var res int
	for i := 0; i < len(s); i++ {
		var counts [26]int
		var maxFreq int
		for _, val := range s[i:] {
			counts[val-'a']++
			maxFreq = max(maxFreq, counts[val-'a'])
			minFreq := len(s)
			for _, count := range counts {
				if count > 0 {
					minFreq = min(minFreq, count)
				}
			}
			res += maxFreq - minFreq
		}
	}
	return res
}
