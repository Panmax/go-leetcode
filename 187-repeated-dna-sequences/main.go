package main

const L = 10

func findRepeatedDnaSequences(s string) []string {
	m := make(map[string]int)
	ans := make([]string, 0)
	// <=
	for i := 0; i+L <= len(s); i++ {
		sub := s[i : i+L]
		m[sub]++
		if m[sub] == 2 {
			ans = append(ans, sub)
		}
	}
	return ans
}
