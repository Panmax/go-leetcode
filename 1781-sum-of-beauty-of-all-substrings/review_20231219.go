package main

func beautySum_20231219(s string) int {
	n := len(s)
	var res int
	for i := 0; i < n; i++ {
		var maxFreq int
		counts := [26]int{}
		for _, c := range s[i:] {
			counts[c-'a']++
			maxFreq = max(maxFreq, counts[c-'a'])
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
