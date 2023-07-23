package main

func findLUSlength_202300723(strs []string) int {
	res := -1
	for i := 0; i < len(strs); i++ {
		check := true
		for j := 0; j < len(strs); j++ {
			if i != j && isSub_20230723(strs[i], strs[j]) {
				check = false
				break
			}
		}
		if check {
			res = max(res, len(strs[i]))
		}
	}
	return res
}

func isSub_20230723(s1 string, s2 string) bool {
	i, j := 0, 0
	for i < len(s1) && j < len(s2) {
		if s1[i] == s2[j] {
			i++
		}
		j++
	}
	return i == len(s1)
}
