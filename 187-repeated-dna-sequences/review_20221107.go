package main

func findRepeatedDnaSequences_20221107(s string) []string {
	m := make(map[string]int)
	var ans []string
	// <=
	for i := 0; i <= len(s)-L; i++ {
		sub := s[i : i+L]
		m[sub]++
		if m[sub] == 2 {
			ans = append(ans, sub)
		}
	}
	return ans
}
