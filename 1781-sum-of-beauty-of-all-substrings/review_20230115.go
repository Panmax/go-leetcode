package main

func beautySum_20230115(s string) int {
	var ans int
	for i := 0; i < len(s); i++ {
		counts := [26]int{}
		var maxFreq int
		for _, val := range s[i:] {
			counts[val-'a']++
			maxFreq = max(maxFreq, counts[val-'a'])
			minFreq := len(s)
			for _, c := range counts {
				if c > 0 {
					minFreq = min(minFreq, c)
				}
			}
			ans += maxFreq - minFreq
		}
	}
	return ans
}
