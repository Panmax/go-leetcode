package main

func findRepeatedDnaSequences_20230108(s string) []string {
	var res []string
	m := make(map[string]int)
	for i := 0; i+L <= len(s); i++ {
		sub := s[i : i+L]
		m[sub]++
		if m[sub] == 2 {
			res = append(res, sub)
		}
	}
	return res
}
