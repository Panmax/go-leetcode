package main

func findLUSlength_20230403(strs []string) int {
	res := -1
	for i := 0; i < len(strs); i++ {
		check := true
		for j := 0; j < len(strs); j++ {
			if i != j && isSub(strs[i], strs[j]) {
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
